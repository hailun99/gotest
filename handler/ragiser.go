package handler

import (
	"log"
	"time"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

// 定义一个结构体接收请求参数
type ragiserReq struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码

}

// 定义一个结构体返回响应结果
type ragiserRes struct {
	Success bool `json:"success"`
}

func HanbleRagiser(c echo.Context) error {
	req := &ragiserReq{}

	c.Bind(req)

	r, err := dbutil.DB.Exec("INSERT INTO users(username, password, created) VALUES (?, ?, ?)",
		req.Username, req.Password, time.Now().Unix())
	if err != nil {
		return err
	} else {
		log.Println(r)
	}

	res := &ragiserRes{
		Success: true,
	}
	return c.JSON(200, res)

}
