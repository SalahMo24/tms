package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser handles user creation
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User creation data"
// @Success 201 {object} CreateUserResponse
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (h *UserHandler) CreateUser(c echo.Context) error {
	var req UserCreate
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	// Call service layer
	createdUserId, err := h.userService.CreateUser(req)
	if err != nil {
		switch err.Error() {
		case "username already exists":
			return c.JSON(http.StatusConflict, map[string]string{
				"error": err.Error(),
			})
		default:
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
	}

	// Convert to response
	response := map[string]string{"id": createdUserId}

	return c.JSON(http.StatusCreated, response)
}
