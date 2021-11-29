package errors

import (
	"net/http"

	"github.com/eden-framework/courier/status_error"
)

//go:generate eden generate error
const ServiceStatusErrorCode = 103 * 1e3 // todo rename this

const (
	// 请求参数错误
	BadRequest status_error.StatusErrorCode = http.StatusBadRequest*1e6 + ServiceStatusErrorCode + iota
)

const (
	// 未找到
	NotFound status_error.StatusErrorCode = http.StatusNotFound*1e6 + ServiceStatusErrorCode + iota
	// @errTalk 管理员未找到
	AdminNotFound
)

const (
	// @errTalk 未授权
	Unauthorized status_error.StatusErrorCode = http.StatusUnauthorized*1e6 + ServiceStatusErrorCode + iota
	// @errTalk 用户名或密码错误
	InvalidUserNamePassword
)

const (
	// @errTalk 操作冲突
	Conflict status_error.StatusErrorCode = http.StatusConflict*1e6 + ServiceStatusErrorCode + iota
)

const (
	// @errTalk 不允许操作
	Forbidden status_error.StatusErrorCode = http.StatusForbidden*1e6 + ServiceStatusErrorCode + iota
	// @errTalk 预售状态必须为就绪才能执行开始操作
	BookingStatusForbidStart
	// @errTalk 预售状态必须为进行中才能执行完成操作
	BookingStatusForbidComplete
)

const (
	// 内部处理错误
	InternalError status_error.StatusErrorCode = http.StatusInternalServerError*1e6 + ServiceStatusErrorCode + iota
)

const (
	// 上游错误
	BadGateway status_error.StatusErrorCode = http.StatusBadGateway*1e6 + ServiceStatusErrorCode + iota
)
