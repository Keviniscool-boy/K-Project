package main

import (
	"auth-center/models"
	"auth-center/service"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 确保已 go get 这个驱动
)

func main() {
	// 1. 配置数据库连接
	// 格式: 用户名:密码@tcp(地址:端口)/数据库名
	dsn := "root:123456@tcp(127.0.0.1:3306)/auth_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 2. 模拟一个正常的注册流程
	testReq := models.RegisterRequest{
		Username: "raotao_dev",
		Password: "secure_password",
	}

	fmt.Println("开始执行注册测试...")
	err = service.Register(db, testReq)
	if err != nil {
		fmt.Printf("测试失败: %v\n", err)
	} else {
		fmt.Println("测试通过: 用户注册成功！")
	}

	// 3. 再次执行相同请求，测试“重复注册”逻辑
	fmt.Println("开始执行重复注册测试...")
	err = service.Register(db, testReq)
	if err != nil {
		fmt.Printf("预期内的失败 (逻辑正确): %v\n", err)
	}
}
