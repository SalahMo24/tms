package transactionlogs

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"tms/app/types"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mocks ---

type MockTransactionLogService struct {
	mock.Mock
}

func (m *MockTransactionLogService) Create(req TransactionLogCreate) (string, error) {
	args := m.Called(req)
	return args.String(0), args.Error(1)
}

func (m *MockTransactionLogService) UpdateTransactionLogStatus(status types.Status, id string) (string, error) {
	args := m.Called(status, id)
	return args.String(0), args.Error(1)
}

// --- Handler Test ---

func TestTransactionLogHandler_Create_Success(t *testing.T) {
	e := echo.New()

	mockService := new(MockTransactionLogService)
	handler := &TransactionLogHandler{transactionLogService: mockService}

	reqBody := `{"amount": 100, "transaction_type": "debit", "account_id": "2fbcd408-22bc-4b3f-a7df-c11c74a19c9d"}`
	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBufferString(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	expectedReq := TransactionLogCreate{Amount: 100, TransactionType: "debit", AccountId: "2fbcd408-22bc-4b3f-a7df-c11c74a19c9d"}
	mockService.On("Create", expectedReq).Return("2fbcd408-22bc-4b3f-a7df-c11c74a19c9d", nil)

	err := handler.Create(c)
	log.Println(err)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.JSONEq(t, `{"id":"2fbcd408-22bc-4b3f-a7df-c11c74a19c9d"}`, rec.Body.String())
	mockService.AssertExpectations(t)
}

func TestTransactionLogHandler_Create_BadRequest(t *testing.T) {
	e := echo.New()

	mockService := new(MockTransactionLogService)
	handler := &TransactionLogHandler{transactionLogService: mockService}

	// Invalid JSON
	reqBody := `{"amount": "oops"}`
	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBufferString(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err := handler.Create(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.JSONEq(t, `{"error":"Invalid request payload"}`, rec.Body.String())
}

func TestTransactionLogHandler_Create_ServiceError(t *testing.T) {
	e := echo.New()

	mockService := new(MockTransactionLogService)
	handler := &TransactionLogHandler{transactionLogService: mockService}

	reqBody := `{"amount": 100, "note": "test"}`
	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBufferString(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	expectedReq := TransactionLogCreate{Amount: 100}
	mockService.On("Create", expectedReq).Return("", errors.New("service failed"))

	err := handler.Create(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.JSONEq(t, `{"error":"service failed"}`, rec.Body.String())
}
