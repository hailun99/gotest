package handler

import (
	"strconv"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

// 定义一个结构体
type Movie struct {
	Id          int64  `json:"id"`          // 电影id
	Title       string `json:"title"`       // 电影标题
	Description string `json:"description"` // 电影介绍
	Created     int64  `json:"created"`     // 插入时间戳
	Directo     string `json:"directo"`     //导演
	Performer   string `json:"performer"`   // 演员

}

type movieRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 定义一个方法       参数         返回值
func HandleMovie(c echo.Context) error {
	//用一个匿名函数来接收id, 一个长度为10容量为64的变量
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	// 初始化Movie
	mov := &Movie{}

	res := &movieRes{}

	// 查询数据库里的信息并赋值
	err := dbutil.DB.QueryRow("select id, title, description, created, directo, performer from movies where id = ?", id).
		Scan(&mov.Id, &mov.Title, &mov.Description, &mov.Created, &mov.Directo, &mov.Performer)
	if err != nil {
		res.Code = 100020
		res.Msg = err.Error()
		return c.JSON(200, res)
	}
	// 把Map中的id赋值给res
	res.Code = 0
	res.Msg = "ok"
	return c.JSON(200, res)
}
