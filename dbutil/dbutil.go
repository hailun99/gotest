package dbutil

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //init()
)

var DB *sql.DB //声明全局变量

func Init() error {
	//链接数据库 用户名:密码@tcp(ip:端口号)/数据库的名字
	//不会检验账号密码是否正确
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gotest") //设置链接数据库的参数
	if err != nil {
		return err
	} else {
		DB = db
		return nil
	}
}
