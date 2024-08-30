package handler

import (
	"log"
	"strconv"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

// 返回一个相应结果
type deleteMovieRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 定义一个接收参数，并返回结果
func HandleDeleteMovie(c echo.Context) error {
	// 绑定请求参数
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	res := &deleteMovieRes{}

	// 删除数据库里的数据
	r, err := dbutil.DB.Exec("delete from movies where id = ?", id)
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
