package accountbalance

import (
	"encoding/json"
	"net/http"
	"time"
	transactionlogs "tms/app/transaction-logs"
	"tms/app/types"
	locking "tms/utils/lock"
	"tms/utils/queue"

	"github.com/labstack/echo/v4"
)

type AccountBalanceHandler struct {
	transactionLogService transactionlogs.TransactionLogService
	accountBalanceService AccountBalanceService
	queueService          *queue.InMemoryQueue
	lockService           *locking.InMemoryLock
}

var queueRunning = false

func NewAccountBalanceHandler(
	abs AccountBalanceService,
	tls transactionlogs.TransactionLogService,
) *AccountBalanceHandler {
	return &AccountBalanceHandler{
		transactionLogService: tls,
		accountBalanceService: abs,
		queueService:          queue.NewInMemoryQueue(),
		lockService:           locking.NewInMemoryLock(),
	}
}

func (h *AccountBalanceHandler) Create(c echo.Context) error {
	var req transactionlogs.TransactionLogCreate
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	// Serialize the request for the queue
	reqData, err := json.Marshal(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to process transaction",
		})
	}

	// Create initial transaction log (pending status)
	transactionId, err := h.transactionLogService.Create(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	// Enqueue the transaction for processing
	if err := h.queueService.EnqueueTransaction(transactionId, reqData); err != nil {
		h.transactionLogService.UpdateTransactionLogStatus(types.Rejected, transactionId)
		return c.JSON(http.StatusServiceUnavailable, map[string]string{
			"error": "Transaction queuing failed",
		})
	}

	// Start processing if not already running
	if !queueRunning {
		queueRunning = true
		go h.processQueuedTransactions()
	}

	return c.JSON(http.StatusAccepted, map[string]string{
		"message":       "Transaction accepted for processing",
		"transactionId": transactionId,
	})
}

func (h *AccountBalanceHandler) processQueuedTransactions() {
	for {
		// Get next transaction from queue
		transactionId, data, err := h.queueService.DequeueTransaction()
		if err != nil {
			// Queue is empty, wait a bit before checking again
			time.Sleep(100 * time.Millisecond)
			continue
		}

		var req transactionlogs.TransactionLogCreate
		if err := json.Unmarshal(data, &req); err != nil {
			h.transactionLogService.UpdateTransactionLogStatus(types.Rejected, transactionId)
			continue
		}

		// Acquire lock for the account
		lockKey := req.AccountId
		acquired, err := h.lockService.AcquireLock(lockKey, 30*time.Second)
		if err != nil || !acquired {
			// Couldn't get lock, requeue and try later
			h.queueService.EnqueueTransaction(transactionId, data)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		// Process the transaction
		_, err = h.accountBalanceService.Create(AccountBalanceCreate{
			Amount:          req.Amount,
			TransactionType: req.TransactionType,
			TransactionId:   transactionId,
			AccountId:       req.AccountId,
		})

		// Release the lock
		h.lockService.ReleaseLock(lockKey)

		// Update transaction status
		if err != nil {
			h.transactionLogService.UpdateTransactionLogStatus(types.Rejected, transactionId)
		} else {
			h.transactionLogService.UpdateTransactionLogStatus(types.Processed, transactionId)
		}
	}
}
