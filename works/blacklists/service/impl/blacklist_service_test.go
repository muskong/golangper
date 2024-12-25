package impl

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"blacklists/domain/entity"
	"blacklists/service/dto"
)

type MockBlacklistRepository struct {
	mock.Mock
}

func (m *MockBlacklistRepository) Create(ctx *gin.Context, blacklist *entity.Blacklist) error {
	args := m.Called(ctx, blacklist)
	return args.Error(0)
}

func (m *MockBlacklistRepository) Update(ctx *gin.Context, blacklist *entity.Blacklist) error {
	args := m.Called(ctx, blacklist)
	return args.Error(0)
}

func (m *MockBlacklistRepository) Delete(ctx *gin.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockBlacklistRepository) FindByID(ctx *gin.Context, id int) (*entity.Blacklist, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Blacklist), args.Error(1)
}

func (m *MockBlacklistRepository) CheckByPhone(ctx *gin.Context, phone string) (*entity.Blacklist, error) {
	args := m.Called(ctx, phone)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Blacklist), args.Error(1)
}

func (m *MockBlacklistRepository) CheckByIDCard(ctx *gin.Context, idCard string) (*entity.Blacklist, error) {
	args := m.Called(ctx, idCard)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Blacklist), args.Error(1)
}

func (m *MockBlacklistRepository) CheckByName(ctx *gin.Context, name string) (*entity.Blacklist, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Blacklist), args.Error(1)
}

func (m *MockBlacklistRepository) List(ctx *gin.Context, page, size int) ([]*entity.Blacklist, int64, error) {
	args := m.Called(ctx, page, size)
	return args.Get(0).([]*entity.Blacklist), args.Get(1).(int64), args.Error(2)
}

func (m *MockBlacklistRepository) UpdateStatus(ctx *gin.Context, id int, status int) error {
	args := m.Called(ctx, id, status)
	return args.Error(0)
}

type MockQueryLogRepository struct {
	mock.Mock
}

func (m *MockQueryLogRepository) Create(ctx *gin.Context, log *entity.QueryLog) error {
	args := m.Called(ctx, log)
	return args.Error(0)
}

func (m *MockQueryLogRepository) List(ctx *gin.Context, merchantID int, page, size int) ([]*entity.QueryLog, int64, error) {
	args := m.Called(ctx, merchantID, page, size)
	return args.Get(0).([]*entity.QueryLog), args.Get(1).(int64), args.Error(2)
}

func TestBlacklistService_Create(t *testing.T) {
	mockRepo := new(MockBlacklistRepository)
	mockQueryLogRepo := new(MockQueryLogRepository)
	service := NewBlacklistService(mockRepo, mockQueryLogRepo)

	ctx := &gin.Context{}
	req := &dto.CreateBlacklistDTO{
		Name:   "Test User",
		Phone:  "12345678901",
		IDCard: "123456789012345678",
	}

	mockRepo.On("Create", ctx, mock.AnythingOfType("*entity.Blacklist")).Return(nil)

	err := service.Create(ctx, req)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBlacklistService_Check(t *testing.T) {
	mockRepo := new(MockBlacklistRepository)
	mockQueryLogRepo := new(MockQueryLogRepository)
	service := NewBlacklistService(mockRepo, mockQueryLogRepo)

	ctx := &gin.Context{}
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
