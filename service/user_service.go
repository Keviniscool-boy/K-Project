package service

import (
	"auth-center/dao"
	"auth-center/models"
	"database/sql"
	"errors"
)

// Register 核心注册业务逻辑
func Register(db *sql.DB, req models.RegisterRequest) error {
	// 1. 逻辑校验：检查用户名是否重复
	existing, _ := dao.GetUserByUsername(db, req.Username)
	if existing != nil {
		return errors.New("conflict: username already exists")
	}

	// 2. 转换模型
	newUser := &models.User{
		Username: req.Username,
		Password: req.Password, // 实际开发建议此处进行 bcrypt 加密
	}

	// 3. 调用 DAO 落地数据
	return dao.CreateUser(db, newUser)
}
