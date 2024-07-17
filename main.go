package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func s1(c echo.Context) error {
	return c.String(http.StatusOK, "hello,world")
}

// 定义一个结构体接收请求参数
type loginReq struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

// 定义一个结构体返回响应结果
type loginRes struct {
	Token string `json:"token"`
}

func login(c echo.Context) error {
	req := &loginReq{}

	// 绑定请求参数值
	c.Bind(req)
	fmt.Println(req.Username)
	fmt.Println(req.Password)

	res := &loginRes{}

	// 判断用户名密码是否正确
	if req.Username == "root" && req.Password == "123456" {

		// 正确返回结果
		res.Token = "abcd"
		return c.JSON(200, res)

	} else {
		// 否则返回错误
		return errors.New("用户名密码错误")
	}
}

func main() {
	// 调用echo通过New赋值给e
	e := echo.New()

	e.Static("/", "static")

	e.POST("/api/login", login)
	//e使用echo里的GET方法，调用上面 s1 与 s2
	e.GET("/s1", s1)
	// e调用方法，开始启动监听下面的端口号
	e.Logger.Fatal(e.Start("0.0.0.0:1323"))
}
