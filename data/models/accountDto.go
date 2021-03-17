package models

import "time"

type AccountDto struct {
	Id        *int       `json:"id,omitempty"`
	FirstName *string    `json:"firstName,omitempty"`
	LastName  *string    `json:"lastName,omitempty"`
	Email     *string    `json:"email,omitempty"`
	Password  *string    `json:"password,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	UpdatedOn *time.Time `json:"updatedOn,omitempty"`
}
