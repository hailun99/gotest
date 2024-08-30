package handler

import (
	"log"
	"strconv"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

// 增加电影导演
// 1. 数据库要加字段，学习 mysql 加字段
// 2. 所有涉及到电影相关的接口都要加导演字段

// 定义一个请求参数
type editMovieReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Directo     string `json:"directo"` //导演
}

// 返回一个相应结果
type editMovieRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 定义一个接收参数，并返回结果
func HandleEditMovie(c echo.Context) error {
	req := &editMovieReq{}

	// 绑定请求参数
	c.Bind(req)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	res := &editMovieRes{}

	// 更新数据
	r, err := dbutil.DB.Exec("update movies set title = ?, description = ?, directo = ? where id = ?",
		&req.Title, &req.Description, &req.Directo, id)
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
