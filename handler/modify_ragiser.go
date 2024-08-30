package handler

import (
	"log"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type editRagiserReq struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	Oldpassword string `json:"oldpassword"` //旧密码
}

// 返回一个相应结果
type editRagiserRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 定义一个接收参数，并返回结果
func HandleModifyRagiser(c echo.Context) error {
	req := &editRagiserReq{}

	// 绑定请求参数
	c.Bind(req)

	res := &editRagiserRes{}

	// 更新数据
	r, err := dbutil.DB.Exec("update users set password = ? where username = ? AND password = ?",
		req.Password, req.Username, req.Oldpassword)
	if err != nil {
		res.Code = 100010
		res.Msg = err.Error()
		return c.JSON(200, res)
	} else {
		log.Print(r)
	}

	res.Code = 0
	res.Msg = "ok"
	//正确返回相应结果
	return c.JSON(200, res)
}
