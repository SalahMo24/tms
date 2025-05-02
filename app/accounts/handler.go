package accounts

import (
	"net/http"
	"tms/app/users"

	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	userService     users.UserService
	accountsService AccountService
}

func NewAccountHandler(userService users.UserService, accountServce AccountService) *AccountHandler {
	return &AccountHandler{
		userService:     userService,
		accountsService: accountServce,
	}
}

func (h *AccountHandler) AccountCreate(c echo.Context) error {
	var req users.UserCreate
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	// Call service layer
	createdUserId, err := h.userService.CreateUser(req)
	var accountId string
	if err != nil {
		switch err.Error() {
		case "user already exists":
			accountId, err = h.accountsService.Create(createdUserId)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": err.Error(),
				})
			}

		default:
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
	} else {
		accountId, err = h.accountsService.Create(createdUserId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

	}

	// Convert to response
	response := map[string]string{"user_id": createdUserId, "account_id": accountId}

	return c.JSON(http.StatusCreated, response)
}
