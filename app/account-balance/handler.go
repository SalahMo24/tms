package accountbalance

import (
	"net/http"
	transactionlogs "tms/app/transaction-logs"
	"tms/app/types"

	"github.com/labstack/echo/v4"
)

type AccountBalanceHandler struct {
	transactionLogService transactionlogs.TransactionLogService
	accountBalanceService AccountBalanceService
}

func NewAccountBalanceHandler(accountBalanceService AccountBalanceService, transactionLogService transactionlogs.TransactionLogService) *AccountBalanceHandler {
	return &AccountBalanceHandler{
		transactionLogService: transactionLogService,
		accountBalanceService: accountBalanceService,
	}
}

func (h *AccountBalanceHandler) Create(c echo.Context) error {
	var req transactionlogs.TransactionLogCreate
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
	newBalanceId, err := h.accountBalanceService.Create(AccountBalanceCreate{
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionId:   transactionId,
		AccountId:       req.AccountId,
	})

	if err != nil {
		transactionId, err := h.transactionLogService.UpdateTransactionLogStatus(types.Rejected, transactionId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error":         err.Error(),
				"transactionId": transactionId,
			})

		}
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":         "transaction rejected",
			"transactionId": transactionId,
		})

	}
	transactionId, err = h.transactionLogService.UpdateTransactionLogStatus(types.Processed, transactionId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})

	}

	// Convert to response
	response := map[string]string{"transactionId": transactionId, "newBalanceId": newBalanceId}

	return c.JSON(http.StatusCreated, response)
}
