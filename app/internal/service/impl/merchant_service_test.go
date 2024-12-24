package impl

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"blackapp/internal/domain/entity"
	"blackapp/internal/service/dto"
)

type MockMerchantRepository struct {
	mock.Mock
}

func (m *MockMerchantRepository) Create(ctx *gin.Context, merchant *entity.Merchant) error {
	args := m.Called(ctx, merchant)
	return args.Error(0)
}

func (m *MockMerchantRepository) Update(ctx *gin.Context, merchant *entity.Merchant) error {
	args := m.Called(ctx, merchant)
	return args.Error(0)
}

func (m *MockMerchantRepository) Delete(ctx *gin.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockMerchantRepository) FindByID(ctx *gin.Context, id int) (*entity.Merchant, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Merchant), args.Error(1)
}

func (m *MockMerchantRepository) FindByAPIKey(ctx *gin.Context, apiKey string) (*entity.Merchant, error) {
	args := m.Called(ctx, apiKey)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Merchant), args.Error(1)
}

func (m *MockMerchantRepository) List(ctx *gin.Context, page, size int) ([]*entity.Merchant, int64, error) {
	args := m.Called(ctx, page, size)
	return args.Get(0).([]*entity.Merchant), args.Get(1).(int64), args.Error(2)
}

func (m *MockMerchantRepository) UpdateToken(ctx *gin.Context, id int, token string, expireTime time.Time) error {
	args := m.Called(ctx, id, token, expireTime)
	return args.Error(0)
}

func (m *MockMerchantRepository) UpdateStatus(ctx *gin.Context, id int, status int) error {
	args := m.Called(ctx, id, status)
	return args.Error(0)
}

type MockLoginLogRepository struct {
	mock.Mock
}

func (m *MockLoginLogRepository) Create(ctx *gin.Context, log *entity.LoginLog) error {
	args := m.Called(ctx, log)
	return args.Error(0)
}

func (m *MockLoginLogRepository) List(ctx *gin.Context, userType int, page, size int) ([]*entity.LoginLog, int64, error) {
	args := m.Called(ctx, userType, page, size)
	return args.Get(0).([]*entity.LoginLog), args.Get(1).(int64), args.Error(2)
}

func TestMerchantService_Create(t *testing.T) {
	mockRepo := new(MockMerchantRepository)
	mockLoginLogRepo := new(MockLoginLogRepository)
	service := NewMerchantService(mockRepo, mockLoginLogRepo, "test-secret", 24*time.Hour)

	ctx := &gin.Context{}
	req := &dto.CreateMerchantDTO{
		Name:          "Test Merchant",
		Address:       "Test Address",
		ContactPerson: "Test Person",
		ContactPhone:  "12345678901",
	}

	mockRepo.On("Create", ctx, mock.AnythingOfType("*entity.Merchant")).Return(nil)
	mockRepo.On("FindByID", ctx, mock.AnythingOfType("int")).Return(&entity.Merchant{}, nil)
	mockRepo.On("Update", ctx, mock.AnythingOfType("*entity.Merchant")).Return(nil)

	err := service.Create(ctx, req)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMerchantService_Login(t *testing.T) {
	mockRepo := new(MockMerchantRepository)
	mockLoginLogRepo := new(MockLoginLogRepository)
	service := NewMerchantService(mockRepo, mockLoginLogRepo, "test-secret", 24*time.Hour)

	ctx := &gin.Context{}
	apiKey := "test_key"
	apiSecret := "test_secret"
	req := &dto.MerchantLoginDTO{
		APIKey:    apiKey,
		APISecret: apiSecret,
	}

	merchant := &entity.Merchant{
		ID:        1,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}

	mockRepo.On("FindByAPIKey", ctx, apiKey).Return(merchant, nil)
	mockRepo.On("UpdateToken", ctx, int(1), mock.AnythingOfType("string"), mock.AnythingOfType("time.Time")).Return(nil)

	token, err := service.Login(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}
