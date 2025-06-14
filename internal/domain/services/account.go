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
	Email    string `json:"email"`
	Password string `json:"password"`
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

	activationToken := utils.GenerateRandomToken()
	now := time.Now()
	expiration := now.Add(24 * time.Hour)

	log.Infof("Activation token for %v is %v", account.Email, activationToken)

	err = me.AccountRepository.CreateActivation(account.Id, activationToken, expiration)
	if err != nil {
		return false, err
	}

	activationLink := fmt.Sprintf("%v/activate?activationToken=%v", me.Settings.FrontEndBaseUrl, activationToken)
	emailMessage := fmt.Sprintf("This is your activation link: %v", activationLink)

	err = me.EmailProvider.SendEmail(account.Email, "Welcome to Time To Get A Job!", emailMessage)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (me *AccountService) Login(info *LoginInfo) (*models.Account, error) {
	account, err := me.AccountRepository.GetByEmail(info.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(info.Password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return account, nil
}

func (me *AccountService) GetProfile(acocuntId int) (*models.Account, error) {
	return nil, nil
}

func (me *AccountService) UpdateProfile(account *models.Account) error {
	return nil
}

func (me *AccountService) Activate(activationId string) error {
	return nil
}

func (me *AccountService) ResetPassword(info *ResetPasswordInfo) error {
	return nil
}
