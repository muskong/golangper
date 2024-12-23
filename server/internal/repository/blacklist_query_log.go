package repository

import (
	"blacklist/internal/model"
	"blacklist/internal/pkg/database"
	"context"

	"gorm.io/gorm"
)

type BlacklistQueryLogRepository interface {
	Create(ctx context.Context, log *model.BlacklistQueryLog) error
	FindByMerchantID(ctx context.Context, merchantID uint, page, pageSize int) ([]model.BlacklistQueryLog, int64, error)
	FindByPhone(ctx context.Context, phone string, page, pageSize int) ([]model.BlacklistQueryLog, int64, error)
	DB() *gorm.DB
}

type blacklistQueryLogRepository struct {
	db *database.PostgresDB
}

func NewBlacklistQueryLogRepository(db *database.PostgresDB) BlacklistQueryLogRepository {
	return &blacklistQueryLogRepository{db: db}
}

func (r *blacklistQueryLogRepository) Create(ctx context.Context, log *model.BlacklistQueryLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *blacklistQueryLogRepository) FindByMerchantID(ctx context.Context, merchantID uint, page, pageSize int) ([]model.BlacklistQueryLog, int64, error) {
	var logs []model.BlacklistQueryLog
	var total int64

	offset := (page - 1) * pageSize

	err := r.db.WithContext(ctx).Model(&model.BlacklistQueryLog{}).
		Where("merchant_id = ?", merchantID).
		Count(&total).
		Order("query_time DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&logs).Error

	return logs, total, err
}

func (r *blacklistQueryLogRepository) FindByPhone(ctx context.Context, phone string, page, pageSize int) ([]model.BlacklistQueryLog, int64, error) {
	var logs []model.BlacklistQueryLog
	var total int64

	offset := (page - 1) * pageSize

	err := r.db.WithContext(ctx).Model(&model.BlacklistQueryLog{}).
		Where("phone = ?", phone).
		Count(&total).
		Order("query_time DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&logs).Error

	return logs, total, err
}

func (r *blacklistQueryLogRepository) DB() *gorm.DB {
	return r.db.DB
}
