package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/handlers"
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockSweetRepository
type MockSweetRepository struct {
	mock.Mock
}

func (m *MockSweetRepository) Create(sweet *models.Sweet) error {
	args := m.Called(sweet)
	return args.Error(0)
}

func (m *MockSweetRepository) GetAll() ([]models.Sweet, error) {
	args := m.Called()
	return args.Get(0).([]models.Sweet), args.Error(1)
}

func (m *MockSweetRepository) GetByID(id uint) (*models.Sweet, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Sweet), args.Error(1)
}

func (m *MockSweetRepository) Update(sweet *models.Sweet) error {
	args := m.Called(sweet)
	return args.Error(0)
}

func (m *MockSweetRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockSweetRepository) Search(query string) ([]models.Sweet, error) {
	args := m.Called(query)
	return args.Get(0).([]models.Sweet), args.Error(1)
}

func TestListSweets(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(MockSweetRepository)
	sweets := []models.Sweet{
		{Name: "Laddu", Price: 10},
		{Name: "Jalebi", Price: 5},
	}
	mockRepo.On("GetAll").Return(sweets, nil)

	h := handlers.NewHandler(mockRepo, nil, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	h.ListSweets(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Sweet
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 2, len(response))
	assert.Equal(t, "Laddu", response[0].Name)
	mockRepo.AssertExpectations(t)
}

func TestAddSweet(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(MockSweetRepository)
	mockRepo.On("Create", mock.AnythingOfType("*models.Sweet")).Return(nil)

	h := handlers.NewHandler(mockRepo, nil, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	input := handlers.SweetInput{
		Name:     "Kaju Katli",
		Category: "Dry Fruit",
		Price:    50.0,
		Quantity: 100,
	}
	jsonValue, _ := json.Marshal(input)
	c.Request, _ = http.NewRequest("POST", "/sweets", bytes.NewBuffer(jsonValue))

	h.AddSweet(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestAddSweet_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(MockSweetRepository)
	mockRepo.On("Create", mock.AnythingOfType("*models.Sweet")).Return(errors.New("db error"))

	h := handlers.NewHandler(mockRepo, nil, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	input := handlers.SweetInput{
		Name:     "Kaju Katli",
		Category: "Dry Fruit",
		Price:    50.0,
		Quantity: 100,
	}
	jsonValue, _ := json.Marshal(input)
	c.Request, _ = http.NewRequest("POST", "/sweets", bytes.NewBuffer(jsonValue))

	h.AddSweet(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}
