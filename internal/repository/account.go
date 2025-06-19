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
	res, err := me.db.Exec(sql, item.Email, item.Password, item.Name, item.LastName, item.Role)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Infof("Insert account rows affected: %v", rowsAffected)
	return nil
}

func (me *AccountRepository) CreateActionToken(accountId int, token string, expiration time.Time, tokenType models.ActionType) error {
	sql := "insert into account_action_token(account_id, token, action, expires_at) values($1, $2, $3, $4)"
	res, err := me.db.Exec(sql, strconv.Itoa(accountId), token, tokenType, expiration)
	if err != nil {
		log.Warnf("Create Action token failed: %v", err.Error())
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Infof("Create action token rows affected: %v", rowsAffected)

	return nil
}

func (me *AccountRepository) Update(item *models.Account) error {
	sql := "update account set password=$2, name=$3, last_name=$4, is_active=$5 where id=$1"
	res, err := me.db.Exec(sql, item.Id, item.Password, item.Name, item.LastName, item.IsActive)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Infof("Update account rows affected: %v", rowsAffected)
	return nil
}

func (me *AccountRepository) Delete(item *models.Account) error {
	sql := "delete from account where id=$"
	res, err := me.db.Exec(sql, item.Id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Infof("Delete account rows affected: %v", rowsAffected)
	return nil
}

func (me *AccountRepository) DeleteActionTokensByAccountId(accountId int, tokenType models.ActionType) error {
	sql := "delete from account_action_token where account_id = $1 and action = $2"

	res, err := me.db.Exec(sql, accountId, tokenType)
	if err != nil {
		log.Errorf("Failed to delete activation token for an account: %v", err.Error())
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Infof("Deleted account activation token for account successfully rows affected: %v", rowsAffected)
	return nil
}

func (me *AccountRepository) GetActionToken(token string, tokenType models.ActionType) (*models.AccountActionToken, error) {
	sql := "select * from account_action_token where token = $1 and action = $2"

	actionToken := models.AccountActionToken{}
	err := me.db.Get(&actionToken, sql, token, tokenType)
	if err != nil {
		return nil, err
	}

	return &actionToken, err
}

func (me *AccountRepository) DeleteActionToken(token string, tokenType models.ActionType) error {
	sql := "delete from account_action_token where token = $1 and action = $2"
	res, err := me.db.Exec(sql, token, tokenType)
	if err != nil {
		log.Errorf("Failed to delete activation token %v", err.Error())
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Infof("Delete account activation token successfully rows affected: %v", rowsAffected)
	return nil
}
