package services

import (
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/domain/models"
	"github.com/donkeysharp/time-to-get-a-job-backend/internal/repository"
)

type AccountService struct {
	AccountRepository *repository.AccountRepository
}

type LoginInfo struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
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
	return false, nil
}

func (me *AccountService) Login(info *LoginInfo) (*models.Account, error) {
	return nil, nil
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
