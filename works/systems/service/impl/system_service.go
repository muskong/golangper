package impl

import (
	"runtime"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/muskong/gopermission/works/systems/service/dto"

	"github.com/muskong/gopermission/works/pkgs/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type systemService struct {
	rdb *redis.Client
	db  *gorm.DB
}

func NewSystemService(rdb *redis.Client, db *gorm.DB) *systemService {
	return &systemService{rdb: rdb, db: db}
}

func (s *systemService) GetSystemMetrics(ctx *gin.Context) (*dto.SystemMetrics, error) {
	metrics := &dto.SystemMetrics{}

	// 获取CPU信息
	cpuInfo, err := getCPUInfo()
	if err != nil {
		logger.Logger.Error("获取CPU信息失败", zap.Error(err))
		return nil, err
	}
	metrics.CPU = *cpuInfo

	// 获取内存信息
	memInfo, err := getMemoryInfo()
	if err != nil {
		logger.Logger.Error("获取内存信息失败", zap.Error(err))
		return nil, err
	}
	metrics.Memory = *memInfo

	// 获取Redis信息
	redisInfo, err := getRedisInfo(s.rdb)
	if err != nil {
		logger.Logger.Error("获取Redis信息失败", zap.Error(err))
		return nil, err
	}
	metrics.Redis = *redisInfo

	// 获取PostgreSQL信息
	pgInfo, err := getPostgresInfo(s.db)
	if err != nil {
		logger.Logger.Error("获取PostgreSQL信息失败", zap.Error(err))
		return nil, err
	}
	metrics.Postgres = *pgInfo

	return metrics, nil
}

// GetCPUInfo 获取CPU信息
func getCPUInfo() (*dto.CPUInfo, error) {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	info := &dto.CPUInfo{
		Usage:     cpuPercent[0],
		Cores:     runtime.NumCPU(),
		ModelName: "",
	}
	if len(cpuInfo) > 0 {
		info.ModelName = cpuInfo[0].ModelName
	}
	return info, nil
}

// GetMemoryInfo 获取内存信息
func getMemoryInfo() (*dto.MemoryInfo, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	return &dto.MemoryInfo{
		Total: memInfo.Total,
		Used:  memInfo.Used,
		Free:  memInfo.Free,
		Usage: memInfo.UsedPercent,
	}, nil
}

// GetRedisInfo 获取Redis信息
func getRedisInfo(rdb *redis.Client) (*dto.RedisInfo, error) {
	info := &dto.RedisInfo{
		Connected: false,
	}

	// 检查连接
	if err := rdb.Ping(rdb.Context()).Err(); err != nil {
		return info, nil
	}
	info.Connected = true

	// 获取内存使用情况
	memory, err := rdb.Info(rdb.Context(), "memory").Result()
	if err == nil {
		// 解析内存信息
		for _, line := range strings.Split(memory, "\n") {
			if strings.HasPrefix(line, "used_memory:") {
				fields := strings.Split(line, ":")
				if len(fields) == 2 {
					if value, err := strconv.ParseUint(strings.TrimSpace(fields[1]), 10, 64); err == nil {
						info.UsedMemory = value
					}
				}
				break
			}
		}
	}

	// 获取键数量
	keys, err := rdb.DBSize(rdb.Context()).Result()
	if err == nil {
		info.Keys = int(keys)
	}

	// 获取客户端数量
	if clients, err := rdb.ClientList(rdb.Context()).Result(); err == nil {
		info.Clients = len(strings.Split(clients, "\n")) - 1
	}

	return info, nil
}

// GetPostgresInfo 获取PostgreSQL信息
func getPostgresInfo(db *gorm.DB) (*dto.PostgresInfo, error) {
	info := &dto.PostgresInfo{
		Connected: false,
	}

	sqlDB, err := db.DB()
	if err != nil {
		return info, nil
	}

	// 检查连接
	if err := sqlDB.Ping(); err != nil {
		return info, nil
	}
	info.Connected = true

	// 获取版本信息
	var version string
	db.Raw("SELECT version()").Scan(&version)
	info.Version = version

	// 获取连接数
	var connections int
	db.Raw("SELECT count(*) FROM pg_stat_activity").Scan(&connections)
	info.Connections = connections

	// 获取数据库大小
	var dbSize string
	db.Raw("SELECT pg_size_pretty(pg_database_size(current_database()))").Scan(&dbSize)
	// TODO: 转换dbSize字符串为uint64

	return info, nil
}
