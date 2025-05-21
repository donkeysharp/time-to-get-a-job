package repository

import "github.com/donkeysharp/time-to-get-a-job-backend/internal/domain/models"

type AccountRepository struct{}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (me *AccountRepository) Get(id int) (*models.Account, error) {
	return nil, nil
}

func (me *AccountRepository) GetByEmail(email string) (*models.Account, error) {
	return &models.Account{
		Email:    "foo@bar.com",
		Id:       12345,
		Password: "$2a$04$H1dmvBU/5EX0BfnM.84ELuq62383a.5F3PQxTeAhApXC1viGxnjFG",
	}, nil
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
