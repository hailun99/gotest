package handler

import (
	"log"
	"strconv"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type editTickekReq struct {
	Movie string `json:"movie"` // 电影
	Type  string `json:"type"`  // 类型
	Seat  string `json:"seat"`  // 座位
	Price string `json:"price"` // 价格
}

// 返回一个相应结果
type editTickekRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 定义一个接收参数，并返回结果
func HandleEditTickek(c echo.Context) error {
	req := &editTickekReq{}

	// 绑定请求参数
	c.Bind(req)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	res := &editTickekRes{}

	// 通过id更新movies表里面的数据数据
	r, err := dbutil.DB.Exec("update tickek set  Movie = ?, Type = ?, Seat = ?, Price = ? where id = ?",
		&req.Movie, &req.Type, &req.Seat, &req.Price, id)
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
