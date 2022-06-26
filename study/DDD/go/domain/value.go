package domain

import (
	"errors"
)

type FirstName string
type LastName string
type UserName struct {
	FirstName FirstName
	LastName  LastName
}

func newUserName(firstname FirstName, lastname LastName) (*UserName, error) {
	if firstname == "" || lastname == "" {
		return nil, errors.New("need FirstName and LastName")
	}
	if len(firstname) > 10 || len(lastname) > 10 {
		return nil, errors.New("name must be 10 characters or less")
	}

	return &UserName{FirstName: firstname, LastName: lastname}, nil
}

type UserId int64
