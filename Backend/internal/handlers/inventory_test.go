package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/handlers"
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTransactionRepository
type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Create(transaction *models.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) GetAll() ([]models.Transaction, error) {
	args := m.Called()
	return args.Get(0).([]models.Transaction), args.Error(1)
}

func TestPurchaseSweet_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSweetRepo := new(MockSweetRepository)
	mockTransactionRepo := new(MockTransactionRepository)

	sweet := &models.Sweet{
		Name:     "Gulab Jamun",
		Price:    10.0,
		Quantity: 50,
	}
	sweet.ID = 1

	// Setup Expectations
	mockSweetRepo.On("GetByID", uint(1)).Return(sweet, nil)
	mockSweetRepo.On("Update", mock.MatchedBy(func(s *models.Sweet) bool {
		return s.Quantity == 45 // 50 - 5
	})).Return(nil)

	mockTransactionRepo.On("Create", mock.AnythingOfType("*models.Transaction")).Return(nil)

	h := handlers.NewHandler(mockSweetRepo, nil, mockTransactionRepo)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Simulate Auth Middleware setting userId
	c.Set("userId", float64(123)) // JWT often parses as float64
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	input := handlers.InventoryInput{Quantity: 5}
	jsonValue, _ := json.Marshal(input)
	c.Request, _ = http.NewRequest("POST", "/sweets/1/purchase", bytes.NewBuffer(jsonValue))

	h.PurchaseSweet(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockSweetRepo.AssertExpectations(t)
	mockTransactionRepo.AssertExpectations(t)
}

func TestPurchaseSweet_InsufficientStock(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSweetRepo := new(MockSweetRepository)

	sweet := &models.Sweet{
		Name:     "Gulab Jamun",
		Price:    10.0,
		Quantity: 2,
	}
	sweet.ID = 1

	mockSweetRepo.On("GetByID", uint(1)).Return(sweet, nil)

	h := handlers.NewHandler(mockSweetRepo, nil, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	input := handlers.InventoryInput{Quantity: 5}
	jsonValue, _ := json.Marshal(input)
	c.Request, _ = http.NewRequest("POST", "/sweets/1/purchase", bytes.NewBuffer(jsonValue))

	h.PurchaseSweet(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockSweetRepo.AssertExpectations(t)
}
