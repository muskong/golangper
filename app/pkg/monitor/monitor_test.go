package monitor

import (
	"testing"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetCPUInfo(t *testing.T) {
	info, err := GetCPUInfo()
	if err != nil {
		t.Fatalf("获取CPU信息失败: %v", err)
	}
	if info.Cores <= 0 {
		t.Error("CPU核心数应大于0")
	}
}

func TestGetMemoryInfo(t *testing.T) {
	info, err := GetMemoryInfo()
	if err != nil {
		t.Fatalf("获取内存信息失败: %v", err)
	}
	if info.Total <= 0 {
		t.Error("总内存应大于0")
	}
}

func TestGetRedisInfo(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	info, err := GetRedisInfo(rdb)
	if err != nil {
		t.Fatalf("获取Redis信息失败: %v", err)
	}
	if !info.Connected {
		t.Error("Redis应连接成功")
	}
}

func TestGetPostgresInfo(t *testing.T) {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=password dbname=blackapp port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		t.Fatalf("初始化数据库失败: %v", err)
	}
	info, err := GetPostgresInfo(db)
	if err != nil {
		t.Fatalf("获取PostgreSQL信息失败: %v", err)
	}
	if !info.Connected {
		t.Error("PostgreSQL应连接成功")
	}
}
