package handler

import (
	"strconv"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type Tickek struct {
	Id      int64  `json:"id"`      // 电影id
	Ciname  string `json:"ciname"`  // 电影院
	Movie   string `json:"movie"`   // 电影
	Type    string `json:"type"`    // 类型
	Seat    string `json:"seat"`    // 座位
	Price   string `json:"price"`   // 价格
	Created string `json:"created"` // 时间戳
}

type tickekRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 定义一个方法       参数         返回值
func HandleTickek(c echo.Context) error {
	//用一个匿名函数来接收id, 一个长度为10容量为64的变量
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	// 初始化Movie
	req := &Tickek{}

	res := &tickekRes{}

	// 查询数据库里的信息并赋值
	err := dbutil.DB.QueryRow("select id, ciname, movie, type, seat, price, Created from tickek where id = ?", id).
		Scan(&req.Id, &req.Ciname, &req.Movie, &req.Type, &req.Seat, &req.Price, &req.Created)
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
