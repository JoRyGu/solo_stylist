/*
Package models provides the basic entity models used for database and API intereactions throughout
the solo_stylist project.
*/
package models

import (
	"time"
)

// Account is used to model the "users" table in the datbase.
type Account struct {
	Id        int        `db:"id" json:"id,omitempty"`
	FirstName string     `db:"first_name" json:"firstName"`
	LastName  string     `db:"last_name" json:"lastName"`
	Email     string     `db:"email" json:"email"`
	Password  string     `db:"password" json:"password"`
	CreatedOn *time.Time `db:"created_on" json:"createdOn,omitempty"`
	UpdatedOn *time.Time `db:"updated_on" json:"updatedOn,omitempty"`
}

// NewAccount creates a basic Account reference and returns it.
func NewAccount(firstName, lastName, email, password string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}
}
