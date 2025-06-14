package repository

import (
	"errors"
	"strconv"
	"time"

	"github.com/donkeysharp/time-to-get-a-job-backend/internal/domain/models"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

var ErrItemNotFound = errors.New("item not found")

type AccountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (me *AccountRepository) Get(id int) (*models.Account, error) {
	account := models.Account{}
	err := me.db.Get(&account, "SELECT * FROM account where id = $1", id)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (me *AccountRepository) GetByEmail(email string) (*models.Account, error) {
	account := models.Account{}
	err := me.db.Get(&account, "SELECT * FROM account where email = $1", email)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (me *AccountRepository) GetAll() ([]*models.Account, error) {
	accounts := []*models.Account{}
	err := me.db.Select(&accounts, "select * from account")
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (me *AccountRepository) Create(item *models.Account) error {
	sql := "insert into account(email, password, name, last_name, role) values ($1, $2, $3, $4, $5)"
	res := me.db.MustExec(sql, item.Email, item.Password, item.Name, item.LastName, item.Role)
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Infof("Insert account rows affected: %v", rowsAffected)
	return nil
}

func (me *AccountRepository) CreateActivation(accountId int, token string, expiration time.Time) error {
	sql := "insert into account_action_token(account_id, token, action, expires_at) values($1, $2, $3, $4)"
	res := me.db.MustExec(sql, strconv.Itoa(accountId), token, "activation", expiration)
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Warnf("Create Actionvation failed: %v", err.Error())
		return err
	}
	log.Infof("CreateActivation rows affected: %v", rowsAffected)

	return nil
}

func (me *AccountRepository) Update(item *models.Account) error {
	sql := "update account set password=$2, name=$3, last_name=$4 where id=$1"
	res := me.db.MustExec(sql, item.Id, item.Password, item.Name, item.LastName)
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Infof("Update account rows affected: %v", rowsAffected)
	return nil
}

func (me *AccountRepository) Delete(item *models.Account) error {
	sql := "delete from account where id=$"
	res := me.db.MustExec(sql, item.Id)
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Infof("Delete account rows affected: %v", rowsAffected)
	return nil
}
