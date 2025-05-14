package transactionlogs

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionLogHandler struct {
	transactionLogService TransactionLogServiceInterface
}

func NewTransactionLogHandler(transactionLogService TransactionLogServiceInterface) *TransactionLogHandler {
	return &TransactionLogHandler{
		transactionLogService: transactionLogService,
	}
}

func (h *TransactionLogHandler) Create(c echo.Context) error {

	var req TransactionLogCreate
	if err := c.Bind(&req); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	// Call service layer
	transactionId, err := h.transactionLogService.Create(req)
	if err != nil {
		log.Println(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})

	}

	// Convert to response
	response := map[string]string{"id": transactionId}

	return c.JSON(http.StatusCreated, response)
}
