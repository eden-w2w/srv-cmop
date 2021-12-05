package errors

import (
	"github.com/eden-framework/courier/status_error"
)

func init() {
	status_error.StatusErrorCodes.Register("BadRequest", 400103000, "请求参数错误", "", false)
	status_error.StatusErrorCodes.Register("NotFound", 404103000, "未找到", "", false)
	status_error.StatusErrorCodes.Register("AdminNotFound", 404103001, "管理员未找到", "", true)
	status_error.StatusErrorCodes.Register("BookingStatusForbidComplete", 403103002, "预售状态必须为进行中才能执行完成操作", "", true)
	status_error.StatusErrorCodes.Register("Conflict", 409103000, "操作冲突", "", true)
	status_error.StatusErrorCodes.Register("InternalError", 500103000, "内部处理错误", "", false)
	status_error.StatusErrorCodes.Register("BookingStatusForbidStart", 403103001, "预售状态必须为就绪才能执行开始操作", "", true)
	status_error.StatusErrorCodes.Register("BadGateway", 502103000, "上游错误", "", false)
	status_error.StatusErrorCodes.Register("InvalidUserNamePassword", 401103001, "用户名或密码错误", "", true)
	status_error.StatusErrorCodes.Register("Forbidden", 403103000, "不允许操作", "", true)
	status_error.StatusErrorCodes.Register("Unauthorized", 401103000, "未授权", "", true)
	status_error.StatusErrorCodes.Register("TemplateForbidDelete", 403103003, "存在已关联的商品，需要解除关联才能删除运费模板", "", true)

}
