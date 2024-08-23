package handler

import (
	"errors"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

// 定义一个结构体接收请求参数
type loginReq struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

// 定义一个结构体返回响应结果
type loginRes struct {
	Token string `json:"token"`
}

func HandleLogin(c echo.Context) error {
	req := &loginReq{}

	// 绑定请求参数值
	c.Bind(req)
	// fmt.Println(req.Username)
	// fmt.Println(req.Password)

	res := &loginRes{}

	user := &loginReq{}
	err := dbutil.DB.QueryRow("select username, password from users where username = ?", req.Username).
		Scan(&user.Username, &user.Password)
	if err != nil {
		return err
	}

	if req.Username == user.Username && req.Password == user.Password {

		// 正确返回结果
		res.Token = "abcd"
		return c.JSON(200, res)
	} else {
		// 否则返回错误
		return errors.New("错误")
	}

}

// 	// // 判断用户名密码是否正确
// 	// if req.Username == "root" && req.Password == "123456" {

// 	// }

// ex, err1 := dbutil.DB.Exec("insert into users set username = ?, password = ?",
// 		&req.Username, &req.Password, time.Now().Unix())
// 	if err1 != nil {
// 		return err1
// 	} else {
// 		log.Println("用户已处在", ex)
// 	}

// sql := "select * from users"
// rows, err := dbutil.DB.Query(sql)
// if err != nil {
// 	return err
// }

// for rows.Next(){
// 	var users loginReq
// 	err :=rows.Scan(&users.Username, &users.Password)
// 	if err != nil {
// 		return err
// 	}
// }
