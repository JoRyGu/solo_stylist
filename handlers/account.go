/*
Package handlers provides all handler functions used by the fiber app instance.
*/
package handlers

import (
	"github.com/JoRyGu/solo_stylist/data/models"
	"github.com/JoRyGu/solo_stylist/services"
	"github.com/gofiber/fiber/v2"
)

// AccountController models a structure that contains all handlers for the account resource.
type AccountController struct {
	accountService *services.AccountService
}

// NewAccountController injects a new instance of AccountController with an instance of *services.AccountService
// and returns a pointer to the new AcccountController
func NewAccountController(accountService *services.AccountService) *AccountController {
	return &AccountController{accountService}
}

// GetAllAccounts is a handler that retrieves all account records from the database and returns
// a JSON representation of them to the API consumer.
func (ac *AccountController) GetAllAccounts(c *fiber.Ctx) error {
	accounts, err := ac.accountService.GetAllAccounts()
	if err != nil {
		return c.Status(500).JSON(&models.HttpError{
			StatusCode:    500,
			StatusMessage: "Internal Server Error",
			Message:       "Error retrieving accounts from database.",
		})
	}

	return c.JSON(accounts)
}

// CreateNewAccount is a handler that receives information required to create a new record on the "users" table,
// passes it on to the *services.AccountService, then returns a JSON representation of the newly created account.
func (ac *AccountController) CreateNewAccount(c *fiber.Ctx) error {
	a := &models.Account{}

	if err := c.BodyParser(a); err != nil {
		return err
	}

	if err := ac.accountService.CreateNewAccount(a); err != nil {
		return err
	}

	return c.JSON(a)
}
