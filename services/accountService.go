/*
Package services implements all business logic for the application.
*/
package services

import (
	"time"

	"github.com/JoRyGu/solo_stylist/data"
	"github.com/JoRyGu/solo_stylist/data/models"
)

// AccountService models the service used to implement the business logic for accounts for the application.
type AccountService struct {
	appContext  *data.AppContext
	authService *AuthService
}

// NewAccountService injects *data.AppContext and *AuthService into a new instance of AccountService
// and returns a pointer to the new AccountService.
func NewAccountService(appContext *data.AppContext, authService *AuthService) *AccountService {
	return &AccountService{appContext, authService}
}

// CreateNewAccount adds CreatedOn and UpdatedOn timestamps to an instance of models.Account, modifies the
// provided password with a hashing function from the AuthService, then saves the record to the database.
// The result will be a fully complete models.Account instance ready for user consumption.
func (as *AccountService) CreateNewAccount(a *models.Account) error {
	nowUtc := time.Now().UTC()
	a.CreatedOn = &nowUtc
	a.UpdatedOn = &nowUtc

	hashedPass, err := as.authService.HashPassword(a.Password)
	if err != nil {
		return err
	}

	a.Password = hashedPass

	err = as.appContext.Accounts.Create(a)
	if err != nil {
		return err
	}

	return nil
}

// GetAccountById retrieves account with the passed id, then adjusts the CreatedOn and UpdatedOn
// properties to UTC, the returns the modified models.Account.
func (as *AccountService) GetAccountById(id int) (*models.AccountDto, error) {
	a, err := as.appContext.Accounts.GetById(id)
	if err != nil {
		return nil, err
	}

	createdUtc := a.CreatedOn.UTC()
	updatedUtc := a.UpdatedOn.UTC()

	a.CreatedOn = &createdUtc
	a.UpdatedOn = &updatedUtc

	dto := a.ToDto()
	dto.Password = nil

	return dto, nil
}

// GetAllAccounts retrieves all accounts from the database through data.AppContext, formats the
// CreatedOn and UpdatedOn dates to UTC, then returns the modified accounts slice.
func (as *AccountService) GetAllAccounts() ([]*models.AccountDto, error) {
	accounts, err := as.appContext.Accounts.GetAll()
	accountDtos := []*models.AccountDto{}

	if err != nil {
		return nil, err
	}

	for _, a := range accounts {
		createdUtc := a.CreatedOn.UTC()
		updatedUtc := a.UpdatedOn.UTC()

		a.CreatedOn = &createdUtc
		a.UpdatedOn = &updatedUtc

		dto := a.ToDto()

		dto.Password = nil

		accountDtos = append(accountDtos, dto)
	}

	return accountDtos, nil
}

// func (as *AccountService) UpdateAccount(id int, account *models.Account) error {
// 	updatedOn := time.Now().UTC()
// 	hashedPass, err := as.authService.HashPassword(account.Password.String)
// 	if err != nil {
// 		return err
// 	}

// 	account.UpdatedOn = &updatedOn
// 	account.Password = sql.NullString{String: hashedPass}

// }
