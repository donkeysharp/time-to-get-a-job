package services

import (
	"errors"
	"fmt"

	"github.com/donkeysharp/time-to-get-a-job-backend/internal/domain/models"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/providers"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/repository"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/web"
	"golang.org/x/crypto/bcrypt"
)

// var InvalidParametersError = fmt.Errorf("Invalid parameters")
var ErrPasswordsDoNotMatch = errors.New("passwords do not match")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrExistingAccount = errors.New("account already exists")

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
	Name            string `json:"name"`
	LastName        string `json:"lastName"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type ResetPasswordInfo struct {
	ResetPasswordToken string `json:"resetToken"`
	Password           string `json:"password"`
	ConfirmaPassword   string `json:"confirmPassword"`
}

func NewAccountService(accountRepo *repository.AccountRepository) *AccountService {
	return &AccountService{
		AccountRepository: accountRepo,
	}
}

func (me *AccountService) SignUp(info *RegisterInfo) (bool, error) {
	if info.Password != info.ConfirmPassword {
		return false, ErrPasswordsDoNotMatch
	}
	account := &models.Account{
		Name:     info.Name,
		LastName: info.LastName,
		Email:    info.Email,
	}
	result, _ := me.AccountRepository.GetByEmail(account.Email)
	if result != nil {
		return false, ErrExistingAccount
	}

	err := me.AccountRepository.Create(account)
	if err != nil {
		return false, err
	}

	activationToken, err := me.AccountRepository.CreateActivation(account)
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
