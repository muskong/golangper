package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"blacklist/internal/model"
	"blacklist/internal/pkg/database"
	"blacklist/internal/pkg/redis"
)

type MerchantRepository struct {
	db *database.PostgresDB
}

func NewMerchantRepository(db *database.PostgresDB) *MerchantRepository {
	return &MerchantRepository{db: db}
}

// Create 创建商户
func (r *MerchantRepository) Create(merchant *model.Merchant) error {
	return r.db.Create(merchant).Error
}

// Update 更新商户
func (r *MerchantRepository) Update(merchant *model.Merchant) error {
	// 更新数据库
	if err := r.db.Save(merchant).Error; err != nil {
		return err
	}

	// 更新缓存
	cacheKey := fmt.Sprintf("merchant:%d", merchant.ID)
	if data, err := json.Marshal(merchant); err == nil {
		redis.RDB.Set(context.Background(), cacheKey, data, 24*time.Hour)
	}

	return nil
}

// Delete 删除商户
func (r *MerchantRepository) Delete(id uint) error {
	// 软删除数据库记录
	if err := r.db.Delete(&model.Merchant{}, id).Error; err != nil {
		return err
	}

	// 删除缓存
	cacheKey := fmt.Sprintf("merchant:%d", id)
	redis.RDB.Del(context.Background(), cacheKey)

	return nil
}

// GetByID 根据ID获取商户
func (r *MerchantRepository) GetByID(id uint) (*model.Merchant, error) {
	var merchant model.Merchant

	// 先从Redis缓存中获取
	cacheKey := fmt.Sprintf("merchant:%d", id)
	if data, err := redis.RDB.Get(context.Background(), cacheKey).Result(); err == nil {
		if err := json.Unmarshal([]byte(data), &merchant); err == nil {
			return &merchant, nil
		}
	}

	// 从数据库中获取
	if err := r.db.First(&merchant, id).Error; err != nil {
		return nil, err
	}

	// 存入Redis缓存
	if data, err := json.Marshal(merchant); err == nil {
		redis.RDB.Set(context.Background(), cacheKey, data, 24*time.Hour)
	}

	return &merchant, nil
}

// GetByAPIKey 根据API Key获取商户
func (r *MerchantRepository) GetByAPIKey(apiKey string) (*model.Merchant, error) {
	var merchant model.Merchant

	// 先从Redis缓存中获取
	cacheKey := fmt.Sprintf("merchant:apikey:%s", apiKey)
	if data, err := redis.RDB.Get(context.Background(), cacheKey).Result(); err == nil {
		if err := json.Unmarshal([]byte(data), &merchant); err == nil {
			return &merchant, nil
		}
	}

	// 从数据库中获取
	if err := r.db.Where("api_key = ?", apiKey).First(&merchant).Error; err != nil {
		return nil, err
	}

	// 存入Redis缓存
	if data, err := json.Marshal(merchant); err == nil {
		redis.RDB.Set(context.Background(), cacheKey, data, 24*time.Hour)
	}

	return &merchant, nil
}

// List 获取商户列表
type MerchantQuery struct {
	Name   string
	Status int
	Page   int
	Size   int
}

func (r *MerchantRepository) List(query *MerchantQuery) ([]model.Merchant, int64, error) {
	var merchants []model.Merchant
	var total int64

	db := r.db.Model(&model.Merchant{})

	// 添加查询条件
	if query.Name != "" {
		db = db.Where("name ILIKE ?", "%"+query.Name+"%")
	}
	if query.Status > 0 {
		db = db.Where("status = ?", query.Status)
	}

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	if err := db.Offset((query.Page - 1) * query.Size).
		Limit(query.Size).
		Order("id DESC").
		Find(&merchants).Error; err != nil {
		return nil, 0, err
	}

	return merchants, total, nil
}

// UpdateStatus 更新商户状态
func (r *MerchantRepository) UpdateStatus(id uint, status int) error {
	return r.db.Model(&model.Merchant{}).Where("id = ?", id).Update("status", status).Error
}

// UpdateToken 更新商户Token
func (r *MerchantRepository) UpdateToken(id uint, token string, expiry time.Time) error {
	return r.db.Model(&model.Merchant{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"token":     token,
			"token_exp": expiry,
		}).Error
}
