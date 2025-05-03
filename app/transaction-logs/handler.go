package transactionlogs

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionLogHandler struct {
	transactionLogService TransactionLogService
}

func NewTransactionLogHandler(transactionLogService TransactionLogService) *TransactionLogHandler {
	return &TransactionLogHandler{
		transactionLogService: transactionLogService,
	}
}

func (h *TransactionLogHandler) Create(c echo.Context) error {
	var req TransactionLogCreate
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	// Call service layer
	transactionId, err := h.transactionLogService.Create(req)
	if err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})

	}

	// Convert to response
	response := map[string]string{"id": transactionId}

	return c.JSON(http.StatusCreated, response)
}
