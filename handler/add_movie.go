package handler

import (
	"log"
	"time"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

// 定义一个请求参数
type addMovieReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Directo     string `json:"directo"`
}

// 返回一个相应结果
type addMovieRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 定义一个接收参数并返回一个错误结果
func HandleAddMovie(c echo.Context) error {
	//取其地址赋值给req
	req := &addMovieReq{}

	//绑定请求参数
	c.Bind(req)

	res := &addMovieRes{}

	// 向数据库里增加数据
	r, err := dbutil.DB.Exec("INSERT INTO movies (title, description, directo, created) VALUES (?, ?, ?, ?)",
		req.Title, req.Description, req.Directo, time.Now().Unix())
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
