package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/donkeysharp/time-to-get-a-job-backend/internal/domain/models"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/providers"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/repository"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/utils"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/web"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

// var InvalidParametersError = fmt.Errorf("Invalid parameters")
var ErrPasswordsDoNotMatch = errors.New("passwords do not match")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrExistingAccount = errors.New("account already exists")
var ErrIncorrectFields = errors.New("incorrect fields")

type AccountService struct {
	AccountRepository *repository.AccountRepository
	EmailProvider     *providers.EmailProvider
	Settings          *web.Settings
}

type LoginInfo struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterInfo struct {
	Name            string `json:"name" validate:"required"`
	LastName        string `json:"lastName" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
}

type ResetPasswordInfo struct {
	ResetPasswordToken string `json:"resetToken"`
	Password           string `json:"password"`
	ConfirmaPassword   string `json:"confirmPassword"`
}

func NewAccountService(accountRepo *repository.AccountRepository, emailProvider *providers.EmailProvider, settings *web.Settings) *AccountService {
	return &AccountService{
		AccountRepository: accountRepo,
		EmailProvider:     emailProvider,
		Settings:          settings,
	}
}

func (me *AccountService) SignUp(info *RegisterInfo) (bool, error) {
	err := utils.Validate.Struct(info)
	if err != nil {
		log.Warnf("SignUp validation error: %v", err.Error())
		return false, ErrIncorrectFields
	}

	if info.Password != info.ConfirmPassword {
		return false, ErrPasswordsDoNotMatch
	}
	if len(info.Password) >= 72 {
		return false, fmt.Errorf("password length must be less than 72")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}

	account := &models.Account{
		Name:     info.Name,
		LastName: info.LastName,
		Email:    info.Email,
		Password: string(password),
		Role:     models.ACCOUNT_ROLE_USER,
	}
	result, _ := me.AccountRepository.GetByEmail(account.Email)
	if result != nil {
		log.Warnf("Account already exists %v", account.Email)
		return false, ErrExistingAccount
	}

	log.Infof("result: %v", result)

	err = me.AccountRepository.Create(account)
	if err != nil {
		return false, err
	}
	// Retrieve account id with new data
	account, err = me.AccountRepository.GetByEmail(account.Email)
	if err != nil {
		return false, err
	}

	if err := me.CreateActivationLink(account); err != nil {
		return false, err
	}
	return true, nil
}

func (me *AccountService) ResendActivationLink(email string) error {
	account, err := me.AccountRepository.GetByEmail(email)
	if err != nil {
		return err
	}

	if account.IsActive {
		return fmt.Errorf("account has already been activated")
	}

	if err := me.AccountRepository.DeleteActionTokensByAccountId(account.Id, models.ACTION_TOKEN_ACTIVATION); err != nil {
		return err
	}

	return me.CreateActivationLink(account)
}

func (me *AccountService) CreateActivationLink(account *models.Account) error {
	activationToken := utils.GenerateRandomToken()
	now := time.Now()
	expiration := now.Add(24 * time.Hour)

	log.Infof("Activation token for %v is %v", account.Email, activationToken)

	err := me.AccountRepository.CreateActionToken(account.Id, activationToken, expiration, models.ACTION_TOKEN_ACTIVATION)
	if err != nil {
		return err
	}

	activationLink := fmt.Sprintf("%v/activate?activationToken=%v", me.Settings.FrontEndBaseUrl, activationToken)
	emailMessage := fmt.Sprintf("This is your activation link: %v", activationLink)

	err = me.EmailProvider.SendEmail(account.Email, "Welcome to Time To Get A Job!", emailMessage)
	if err != nil {
		log.Errorf("Failed to send activation link: %v", err.Error())
	}
	log.Info("Activation link created successfully")
	return nil
}

func (me *AccountService) CreateResetPasswordLink(account *models.Account) error {
	resetToken := utils.GenerateRandomToken()
	now := time.Now()
	expiration := now.Add(24 * time.Hour)

	log.Infof("Password reset token for %v is %v", account.Email, resetToken)

	err := me.AccountRepository.CreateActionToken(account.Id, resetToken, expiration, models.ACTION_TOKEN_PASSWORD_RESET)
	if err != nil {
		return err
	}

	activationLink := fmt.Sprintf("%v/reset?token=%v", me.Settings.FrontEndBaseUrl, resetToken)
	emailMessage := fmt.Sprintf("This is your password reset link: %v", activationLink)

	err = me.EmailProvider.SendEmail(account.Email, "Reset Password", emailMessage)
	if err != nil {
		log.Errorf("Sending reset password email failed: %v", err.Error())
		// return err
	}
	log.Info("Reset password link created successfully.")
	return nil
}

func (me *AccountService) Login(info *LoginInfo) (*models.Account, error) {
	err := utils.Validate.Struct(info)
	if err != nil {
		log.Warnf("SignUp validation error: %v", err.Error())
		return nil, ErrIncorrectFields
	}

	account, err := me.AccountRepository.GetByEmail(info.Email)
	if err != nil {
		log.Warnf("Login failed, %v does not exist", info.Email)
		// for security always send invalid credentials
		// we don't want to make enumeration easy for an attacker
		return nil, ErrInvalidCredentials
	}

	if !account.IsActive {
		log.Warnf("Login failed, %v is not active yet", info.Email)
		return nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(info.Password))
	if err != nil {
		log.Warnf("Login failed, %v password does not match the one in database", info.Email)
		return nil, ErrInvalidCredentials
	}
	log.Infof("Login successfull: %v", info.Email)

	return account, nil
}

func (me *AccountService) GetProfile(acocuntId int) (*models.Account, error) {
	return nil, nil
}

func (me *AccountService) UpdateProfile(account *models.Account) error {
	return nil
}

func (me *AccountService) Activate(token string) error {
	actionToken, err := me.AccountRepository.GetActionToken(token, models.ACTION_TOKEN_ACTIVATION)
	if err != nil {
		return err
	}

	now := time.Now()
	diff := actionToken.ExpiresAt.Sub(now)

	if diff < 0 {
		return fmt.Errorf("activation token expired")
	}

	account, err := me.AccountRepository.Get(actionToken.AccountId)
	if err != nil {
		return err
	}

	account.IsActive = true
	log.Infof("Activating %v account", account.Email)
	err = me.AccountRepository.Update(account)
	if err != nil {
		return err
	}

	log.Infof("Deleting activation token %v", token)
	err = me.AccountRepository.DeleteActionToken(token, models.ACTION_TOKEN_ACTIVATION)
	if err != nil {
		return err
	}
	log.Infof("Activation of account %v was successful", account.Email)
	return nil
}

func (me *AccountService) SendResetPassword(email string) error {
	account, err := me.AccountRepository.GetByEmail(email)
	if err != nil {
		log.Warnf("Failed to get account %v: %v", email, err.Error())
		return err
	}
	if err := me.CreateResetPasswordLink(account); err != nil {
		return err
	}
	log.Infof("Reset password link sent successfully")
	return nil
}

func (me *AccountService) ResetPassword(info *ResetPasswordInfo) error {
	if info.Password != info.ConfirmaPassword {
		return fmt.Errorf("Reset passwords do not match")
	}
	actionToken, err := me.AccountRepository.GetActionToken(info.ResetPasswordToken, models.ACTION_TOKEN_PASSWORD_RESET)
	if err != nil {
		return err
	}
	account, err := me.AccountRepository.Get(actionToken.AccountId)
	if err != nil {
		return err
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	account.Password = string(newPassword)

	if err := me.AccountRepository.Update(account); err != nil {
		return err
	}
	log.Infof("Password reset successfully for %v with reset token %v", account.Email, info.ResetPasswordToken)
	return nil
}
