/*
Package data contains all subpackages related to interacting with the database, including the functions
needed to create a new database connection, models, and repositories.
*/
package data

import (
	"github.com/JoRyGu/solo_stylist/data/repositories"
)

// AppContext is a centralized structure that exposes all of the repositories package members.
type AppContext struct {
	Accounts *repositories.AccountsRepository
}

// NewAppContext generates a new database connection and injects it into a reference to an instance of the AppContext
// struct before returning the new AppContext instance.
func NewAppContext() (*AppContext, error) {
	db, err := GetDatabaseConnection()
	if err != nil {
		return nil, err
	}

	accountsRepo := repositories.NewAccountsRepository(db)

	context := AppContext{accountsRepo}

	return &context, nil
}
