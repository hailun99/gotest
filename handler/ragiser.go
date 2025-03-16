package handler

import (
	"log"
	"time"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

// 定义一个结构体接收请求参数
type ragiserReq struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Gender    string `json:"gender"`    // 性别
	Signature string `json:"signature"` // 个人标签
	Vip       string `json:"vip"`       // 会员
}

// 定义一个结构体返回响应
type ragiserRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func HanbleRagiser(c echo.Context) error {
	req := &ragiserReq{}

	c.Bind(req)

	res := &ragiserRes{}

	r, err := dbutil.DB.Exec("INSERT INTO users(username, password, gender, signature, vip, created) VALUES (?, ?, ?, ?, ?, ?)",
		req.Username, req.Password, req.Gender, req.Signature, req.Vip, time.Now().Unix())
	if err != nil {
		res.Code = 100010
		res.Msg = err.Error()
		return c.JSON(200, res)
	} else {
		log.Println(r)
	}

	res.Code = 0
	res.Msg = "ok"
	return c.JSON(200, res)

}
