package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// HashPassword 使用MD5哈希密码
func HashPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

// GenerateRandomString 生成指定长度的随机字符串
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// ParseSize 解析字符串形式的大小为字节数
func ParseSize(size string) uint64 {
	size = strings.ToLower(strings.TrimSpace(size))
	if size == "" {
		return 0
	}

	var multiplier uint64 = 1
	if strings.HasSuffix(size, "kb") {
		multiplier = 1024
		size = size[:len(size)-2]
	} else if strings.HasSuffix(size, "mb") {
		multiplier = 1024 * 1024
		size = size[:len(size)-2]
	} else if strings.HasSuffix(size, "gb") {
		multiplier = 1024 * 1024 * 1024
		size = size[:len(size)-2]
	} else if strings.HasSuffix(size, "b") {
		size = size[:len(size)-1]
	}

	var value uint64
	_, err := fmt.Sscanf(size, "%d", &value)
	if err != nil {
		return 0
	}

	return value * multiplier
}
