package utils

import "regexp"

// IsPhoneNumber 验证手机号格式是否正确
// 支持以下格式:
// - 11位数字
// - 以1开头
// - 第二位是3-9
// - 后面9位是0-9的数字
func IsPhoneNumber(phone string) bool {
	if phone == "" {
		return false
	}
	pattern := `^1[3-9]\d{9}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(phone)
}

// IsIDCard 验证身份证号格式是否正确
// 支持15位和18位身份证号
func IsIDCard(idCard string) bool {
	if idCard == "" {
		return false
	}
	pattern := `(^\d{15}$)|(^\d{17}([0-9]|X)$)`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(idCard)
}

// IsEmail 验证邮箱格式是否正确
func IsEmail(email string) bool {
	if email == "" {
		return false
	}
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
