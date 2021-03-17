/*
Package handlers provides all handler functions used by the fiber app instance.
*/
package handlers

import (
	"fmt"
	"strconv"

	"github.com/JoRyGu/solo_stylist/data/models"
	"github.com/JoRyGu/solo_stylist/services"
	"github.com/JoRyGu/solo_stylist/util"
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

func (ac *AccountController) GetAccountById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return util.SendError(c, 400, "Invalid ID")
	}

	account, err := ac.accountService.GetAccountById(id)
	if err != nil {
		return util.SendError(c, 404, fmt.Sprintf("Could not find account with ID %d", id))
	}

	return c.JSON(account)
}

// GetAllAccounts is a handler that retrieves all account records from the database and returns
// a JSON representation of them to the API consumer.
func (ac *AccountController) GetAllAccounts(c *fiber.Ctx) error {
	accounts, err := ac.accountService.GetAllAccounts()
	if err != nil {
		return util.SendError(c, 500, "Error retrieving accounts from database.")
	}

	return c.JSON(accounts)
}

// CreateNewAccount is a handler that receives information required to create a new record on the "users" table,
// passes it on to the *services.AccountService, then returns a JSON representation of the newly created account.
func (ac *AccountController) CreateNewAccount(c *fiber.Ctx) error {
	a := &models.Account{}

	if err := c.BodyParser(a); err != nil {
		return util.SendError(c, 500, "Error parsing request body.")
	}

	if err := ac.accountService.CreateNewAccount(a); err != nil {
		return util.SendError(c, 500, "Error while creating new account.")
	}

	return c.JSON(a)
}
