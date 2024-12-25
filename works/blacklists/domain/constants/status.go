package constants

// 通用状态
const (
	StatusEnabled  = 1 // 启用
	StatusDisabled = 2 // 禁用
)

// 黑名单状态
const (
	BlacklistStatusPending  = 0 // 待审核
	BlacklistStatusApproved = 1 // 已通过
	BlacklistStatusRejected = 2 // 已拒绝
)

// 日志状态
const (
	LogStatusSuccess = 1 // 成功
	LogStatusFailed  = 2 // 失败
)
