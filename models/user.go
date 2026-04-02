package models

//User定义用户模型，用于数据库交互
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//RegisterRequest //定义一个结构体，用来接收注册请求的数据
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6"`
}
