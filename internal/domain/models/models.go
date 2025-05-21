package models

type JobPost struct {
}

type JobPostApplication struct {
}

type Account struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string
}
