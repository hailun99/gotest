package handler

import (
	"log"
	"time"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)
// 定义一个电影参数
type addTickekReq struct {
	Ciname   string `json:"ciname"`   // 电影院
	Movie    string `json:"movie"`    // 电影
	Type     string `json:"type"`     // 类型
	Seat     string `json:"seat"`     // 座位
	Price    string `json:"price"`    // 价格
	Comsumer string `json:"comsumer"` // 顾客
}

// 返回一个相应结果
type addTickekRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 定义一个接收参数并返回一个错误结果
func HandleAddTickek(c echo.Context) error {
	//取其地址赋值给req
	req := &addTickekReq{}

	//绑定请求参数
	c.Bind(req)

	res := &addTickekRes{}

	// 向数据库里增加数据
	r, err := dbutil.DB.Exec("INSERT INTO tickek (ciname, movie, type, seat, price, comsumer, created) VALUES (?, ?, ?, ?, ?, ?, ?)",
		req.Ciname, req.Movie, req.Type, req.Seat, req.Price, req.Comsumer, time.Now().Unix())
	//判断是否正确，正确返回结果,错误返回f
	if err != nil {
		res.Code = 100010
		res.Msg = err.Error()
		return c.JSON(200, res)
	} else {
		log.Print(r)
	}

	// moviesMap[3] = &Movie{
	// 	Id:          3,
	// 	Title:       req.Title,
	// 	Description: req.Description,
	// }

	res.Code = 0
	res.Msg = "ok"
	//正确返回相应结果
	return c.JSON(200, res)
}
