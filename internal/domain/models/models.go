package models

import "time"

const (
	ACCOUNT_ROLE_USER     = "user"
	ACCOUNT_ROLE_ADMIN    = "admin"
	ACCOUNT_ROLE_EMPLOYER = "employer"
)

type Account struct {
	Id        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `db:"password"`
	Name      string    `json:"name" db:"name"`
	LastName  string    `json:"lastName" db:"last_name"`
	Role      string    `db:"role"`
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
}

type AccountActionToken struct {
	Id        int       `db:"id"`
	AccountId int       `db:"account_id"`
	Token     string    `db:"token"`
	Action    string    `db:"action"`
	ExpiresAt time.Time `db:"expires_at"`
	CreatedAt time.Time `db:"created_at"`
}

type JobPost struct {
	Id          int    `json:"id"`
	AccountId   int    `json:"accountId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   time.Time
}

type JobPostApplication struct {
	Id              int    `json:"id"`
	JobPostId       int    `json:"postId"`
	ApplicantId     int    `json:"applicantId"`
	ApplicationText string `json:"applicationText"`
	AppliedAt       time.Time
}
