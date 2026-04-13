package entity

import "errors"

var ErrInvalidEmail = errors.New("invalid email address")
var ErrPasswordTooShort = errors.New("password too short")

type User struct {
	ID           string
	Email        string
	Name         string
	PasswordHash string
	SessionToken string
}

func NewUser(id, email, name, password string) (*User, error) {
	if email == "" {
		return nil, ErrInvalidEmail
	}
	if len(password) < 8 {
		return nil, ErrPasswordTooShort
	}
	return &User{
		ID:           id,
		Email:        email,
		Name:         name,
		PasswordHash: password,
	}, nil
}
