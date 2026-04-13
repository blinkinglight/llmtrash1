package entity

import "errors"

var ErrInvalidEmail = errors.New("invalid email address")

type User struct {
	ID    string
	Email string
	Name  string
}

func NewUser(id, email, name string) (*User, error) {
	if email == "" {
		return nil, ErrInvalidEmail
	}
	return &User{
		ID:    id,
		Email: email,
		Name:  name,
	}, nil
}
