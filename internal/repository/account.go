package repository

import "github.com/donkeysharp/time-to-get-a-job-backend/internal/domain/models"

type AccountRepository struct{}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (me *AccountRepository) Get(id int) (*models.Account, error) {
	return nil, nil
}

func (me *AccountRepository) GetAll() ([]*models.Account, error) {
	return nil, nil
}

func (me *AccountRepository) Create(item *models.Account) error {
	return nil
}

func (me *AccountRepository) Update(item *models.Account) error {
	return nil
}

func (me *AccountRepository) Delete(item *models.Account) error {
	return nil
}
