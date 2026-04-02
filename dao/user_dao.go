package dao

import (
	"auth-center/models"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// CreateUser将新用户存入数据库
func CreateUser(db *sql.DB, user *models.User) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	//value使用？？占位符，防止SQL注入攻击。便于数据库驱动程序正确处理输入数据。
	_, err := db.Exec(query, user.Username, user.Password)
	return err

}

// 建立函数GetUserByUsername从数据库中根据用户名查询用户信息，用于登录验证
func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = ?"
	var user models.User
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
