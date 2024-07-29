package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// 定义一个结构体
type Movie struct {
	Id          int64  `json:"id"`          // 电影id
	Title       string `json:"title"`       // 电影标题
	Description string `json:"description"` // 电影介绍
}

// 定义一个方法       参数         返回值
func HandleMovie(c echo.Context) error {
	//用一个匿名函数与id来接收, 一个长度为10容量为64的变量
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	// 把Map中的id赋值给res
	res := moviesMap[id]
	return c.JSON(200, res)
}
