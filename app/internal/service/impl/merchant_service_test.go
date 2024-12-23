package impl

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"blackapp/internal/domain/entity"
	"blackapp/internal/service/dto"
)

type MockMerchantRepository struct {
	mock.Mock
}

func (m *MockMerchantRepository) Create(ctx context.Context, merchant *entity.Merchant) error {
	args := m.Called(ctx, merchant)
	return args.Error(0)
}

func (m *MockMerchantRepository) Update(ctx context.Context, merchant *entity.Merchant) error {
	args := m.Called(ctx, merchant)
	return args.Error(0)
}

func (m *MockMerchantRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockMerchantRepository) FindByID(ctx context.Context, id uint) (*entity.Merchant, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Merchant), args.Error(1)
}

func (m *MockMerchantRepository) FindByAPIKey(ctx context.Context, apiKey string) (*entity.Merchant, error) {
	args := m.Called(ctx, apiKey)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Merchant), args.Error(1)
}

func TestMerchantService_Create(t *testing.T) {
	mockRepo := new(MockMerchantRepository)
	service := NewMerchantService(mockRepo)

	ctx := context.Background()
	req := &dto.CreateMerchantDTO{
		Name:          "Test Merchant",
		Address:       "Test Address",
		ContactPerson: "Test Person",
		ContactPhone:  "12345678901",
	}

	mockRepo.On("Create", ctx, mock.AnythingOfType("*entity.Merchant")).Return(nil)
	mockRepo.On("FindByID", ctx, mock.AnythingOfType("uint")).Return(&entity.Merchant{}, nil)
	mockRepo.On("Update", ctx, mock.AnythingOfType("*entity.Merchant")).Return(nil)

	err := service.Create(ctx, req)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMerchantService_Login(t *testing.T) {
	mockRepo := new(MockMerchantRepository)
	service := NewMerchantService(mockRepo)

	ctx := context.Background()
	apiKey := "test_key"
	apiSecret := "test_secret"

	merchant := &entity.Merchant{
		ID:        1,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}

	mockRepo.On("FindByAPIKey", ctx, apiKey).Return(merchant, nil)
	mockRepo.On("UpdateToken", ctx, uint(1), mock.AnythingOfType("string"), mock.AnythingOfType("time.Time")).Return(nil)

	token, err := service.Login(ctx, apiKey, apiSecret)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}
