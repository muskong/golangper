package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"blacklist/internal/model"
	"blacklist/internal/pkg/database"
	"blacklist/internal/pkg/redis"
)

type BlacklistRepository struct {
	db *database.PostgresDB
}

func NewBlacklistRepository(db *database.PostgresDB) *BlacklistRepository {
	return &BlacklistRepository{db: db}
}

func (r *BlacklistRepository) Create(user *model.BlacklistUser) error {
	return r.db.Create(user).Error
}

func (r *BlacklistRepository) GetByID(id uint) (*model.BlacklistUser, error) {
	var user model.BlacklistUser

	// 先从Redis缓存中获取
	cacheKey := fmt.Sprintf("blacklist:user:%d", id)
	if data, err := redis.RDB.Get(context.Background(), cacheKey).Result(); err == nil {
		if err := json.Unmarshal([]byte(data), &user); err == nil {
			return &user, nil
		}
	}

	// 从数据库中获取
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	// 存入Redis缓存
	if data, err := json.Marshal(user); err == nil {
		redis.RDB.Set(context.Background(), cacheKey, data, 24*time.Hour)
	}

	return &user, nil
}

func (r *BlacklistRepository) Update(user *model.BlacklistUser) error {
	// 更新数据库
	if err := r.db.Save(user).Error; err != nil {
		return err
	}

	// 更新缓存
	cacheKey := fmt.Sprintf("blacklist:user:%d", user.ID)
	if data, err := json.Marshal(user); err == nil {
		redis.RDB.Set(context.Background(), cacheKey, data, 24*time.Hour)
	}

	return nil
}

func (r *BlacklistRepository) Delete(id uint) error {
	// 删除数据库记录
	if err := r.db.Delete(&model.BlacklistUser{}, id).Error; err != nil {
		return err
	}

	// 删除缓存
	cacheKey := fmt.Sprintf("blacklist:user:%d", id)
	redis.RDB.Del(context.Background(), cacheKey)

	return nil
}

// 添加查询参数结构体
type BlacklistUserQuery struct {
	Name    string
	Phone   string
	IDCard  string
	Email   string
	Address string
	Remark  string
	Page    int
	Size    int
}

// 修改 List 方法
func (r *BlacklistRepository) List(query *BlacklistUserQuery) ([]model.BlacklistUser, int64, error) {
	var users []model.BlacklistUser
	var total int64

	db := r.db.Model(&model.BlacklistUser{})

	// 添加查询条件
	if query.Name != "" {
		db = db.Where("name ILIKE ?", "%"+query.Name+"%")
	}
	if query.Phone != "" {
		db = db.Where("phone ILIKE ?", "%"+query.Phone+"%")
	}
	if query.IDCard != "" {
		db = db.Where("id_card ILIKE ?", "%"+query.IDCard+"%")
	}
	if query.Email != "" {
		db = db.Where("email ILIKE ?", "%"+query.Email+"%")
	}
	if query.Address != "" {
		db = db.Where("address ILIKE ?", "%"+query.Address+"%")
	}
	if query.Remark != "" {
		db = db.Where("remark ILIKE ?", "%"+query.Remark+"%")
	}

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	if err := db.Offset((query.Page - 1) * query.Size).
		Limit(query.Size).
		Order("id DESC").
		Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// ExistsByPhone 检查指定手机号的用户是否存在
func (r *BlacklistRepository) ExistsByPhone(phone string) (bool, error) {
	var count int64
	err := r.db.Model(&model.BlacklistUser{}).
		Where("phone = ?", phone).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetByPhone 根据手机号获取用户信息
func (r *BlacklistRepository) GetByPhone(phone string) (*model.BlacklistUser, error) {
	var user model.BlacklistUser

	// 先从Redis缓存中获取
	cacheKey := fmt.Sprintf("blacklist:user:phone:%s", phone)
	if data, err := redis.RDB.Get(context.Background(), cacheKey).Result(); err == nil {
		if err := json.Unmarshal([]byte(data), &user); err == nil {
			return &user, nil
		}
	}

	// 从数据库中获取
	if err := r.db.Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}

	// 存入Redis缓存
	if data, err := json.Marshal(user); err == nil {
		redis.RDB.Set(context.Background(), cacheKey, data, 24*time.Hour)
	}

	return &user, nil
}

// ExistsQuery 存在性检查的查询参数
type ExistsQuery struct {
	Phone  string
	IDCard string
	Name   string
}

// CheckExists 检查用户是否存在
func (r *BlacklistRepository) CheckExists(query *ExistsQuery) (bool, error) {
	var count int64
	db := r.db.Model(&model.BlacklistUser{})

	// 构建查询条件
	conditions := make([]string, 0)
	args := make([]interface{}, 0)

	if query.Phone != "" {
		conditions = append(conditions, "phone = ?")
		args = append(args, query.Phone)
	}
	if query.IDCard != "" {
		conditions = append(conditions, "id_card = ?")
		args = append(args, query.IDCard)
	}
	if query.Name != "" {
		conditions = append(conditions, "name = ?")
		args = append(args, query.Name)
	}

	if len(conditions) > 0 {
		db = db.Where(strings.Join(conditions, " AND "), args...)
	} else {
		return false, nil
	}

	// 检查总体是否存在
	if err := db.Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

// GetByIDCard 根据身份证号获取用户信息
func (r *BlacklistRepository) GetByIDCard(idCard string) (*model.BlacklistUser, error) {
	var user model.BlacklistUser

	// 先从Redis缓存中获取
	cacheKey := fmt.Sprintf("blacklist:user:idcard:%s", idCard)
	if data, err := redis.RDB.Get(context.Background(), cacheKey).Result(); err == nil {
		if err := json.Unmarshal([]byte(data), &user); err == nil {
			return &user, nil
		}
	}

	// 从数据库中获取
	if err := r.db.Where("id_card = ?", idCard).First(&user).Error; err != nil {
		return nil, err
	}

	// 存入Redis缓存
	if data, err := json.Marshal(user); err == nil {
		redis.RDB.Set(context.Background(), cacheKey, data, 24*time.Hour)
	}

	return &user, nil
}

// GetByName 根据姓名获取用户信息列表
func (r *BlacklistRepository) GetByName(name string) ([]model.BlacklistUser, error) {
	var users []model.BlacklistUser

	// 从数据库中获取
	if err := r.db.Where("name = ?", name).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
