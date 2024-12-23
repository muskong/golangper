package impl

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"blackapp/internal/service/dto"
)

type MockBlacklistRepository struct {
	mock.Mock
}

func (m *MockBlacklistRepository) Create(ctx context.Context, blacklist *entity.Blacklist) error {
	args := m.Called(ctx, blacklist)
	return args.Error(0)
}

func (m *MockBlacklistRepository) Update(ctx context.Context, blacklist *entity.Blacklist) error {
	args := m.Called(ctx, blacklist)
	return args.Error(0)
}

func (m *MockBlacklistRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockBlacklistRepository) FindByID(ctx context.Context, id uint) (*entity.Blacklist, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Blacklist), args.Error(1)
}

func (m *MockBlacklistRepository) CheckByPhone(ctx context.Context, phone string) (*entity.Blacklist, error) {
	args := m.Called(ctx, phone)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Blacklist), args.Error(1)
}

func TestBlacklistService_Create(t *testing.T) {
	mockRepo := new(MockBlacklistRepository)
	service := NewBlacklistService(mockRepo)

	ctx := context.Background()
	req := &dto.CreateBlacklistDTO{
		Name:       "Test User",
		Phone:      "12345678901",
		IDCard:     "123456789012345678",
		MerchantID: 1,
	}

	mockRepo.On("Create", ctx, mock.AnythingOfType("*entity.Blacklist")).Return(nil)

	err := service.Create(ctx, req)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBlacklistService_Check(t *testing.T) {
	mockRepo := new(MockBlacklistRepository)
	service := NewBlacklistService(mockRepo)

	ctx := context.Background()
	req := &dto.CheckBlacklistDTO{
		Phone: "12345678901",
	}

	blacklist := &entity.Blacklist{
		ID:    1,
		Phone: "12345678901",
	}

	mockRepo.On("CheckByPhone", ctx, req.Phone).Return(blacklist, nil)

	exists, err := service.Check(ctx, req)
	assert.NoError(t, err)
	assert.True(t, exists)
	mockRepo.AssertExpectations(t)
}
