package admins

type LoginParams struct {
	// 用户名
	UserName string `in:"body" json:"userName"`
	// 密码
	Password string `in:"body" json:"password"`
}
