package server

import (
	"net/http"
	"tms/app/accounts"
	"tms/app/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	e.GET("/", s.HelloWorldHandler)

	e.GET("/health", s.healthHandler)

	UserRoutes(e)
	AccountRoutes(e)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}

func UserRoutes(e *echo.Echo) {
	userRepo := users.NewRepository()

	// Initialize service
	userService := users.NewUserService(*userRepo)
	userHandler := users.NewUserHandler(*userService)

	// Group for versioned API routes
	api := e.Group("/api/v1")

	// User routes
	users := api.Group("/users")
	users.POST("", userHandler.CreateUser)

	// Add middleware specific to user routes if needed
	users.Use(middleware.Logger())
}
func AccountRoutes(e *echo.Echo) {
	userRepo := users.NewRepository()
	userService := users.NewUserService(*userRepo)

	accountRepo := accounts.NewRepository()
	accountRepoService := accounts.NewUserService(*accountRepo)
	accounthandler := accounts.NewAccountHandler(*userService, *accountRepoService)

	// Group for versioned API routes
	api := e.Group("/api/v1")

	// User routes
	users := api.Group("/accounts")
	users.POST("", accounthandler.AccountCreate)

	// Add middleware specific to user routes if needed
	users.Use(middleware.Logger())
}
