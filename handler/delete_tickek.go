package handler

import (
	"log"
	"strconv"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type deleteTickekRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 定义一个接收参数，并返回结果
func HandleDeleteTickek(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	res := &deleteTickekRes{}

	// id小于0无法查询
	if id < 0 {
		res.Code = 10001
		res.Msg = "无此id"
		return c.JSON(200, res)
	}

	// 删除数据库里的数据
	r, err := dbutil.DB.Exec("delete from tickek where id = ?", id)
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
