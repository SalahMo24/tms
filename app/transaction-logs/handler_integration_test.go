package transactionlogs

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"tms/app/types"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransactionLog_Integration(t *testing.T) {
	// === Setup ===
	e := echo.New()
	// Initialize real repository, service, handler
	repo := NewRepository()
	svc := NewTransactionLogService(*repo)
	h := NewTransactionLogHandler(svc)

	// Register route
	e.POST("/transaction-logs", h.Create)

	// === Test Request ===
	requestBody := TransactionLogCreate{
		TransactionType: types.Credit, // Adjust to your enum/type
		Amount:          150.75,
		AccountId:       "2fbcd408-22bc-4b3f-a7df-c11c74a19c9d",
	}
	jsonBody, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, "/transaction-logs", bytes.NewBuffer(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// === Execute ===
	if assert.NoError(t, h.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["id"], "-") // Or whatever your ID pattern is
	}

	// Optional: Clean up DB
	// teardownTestDatabase(db)
}
