package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func s1(c echo.Context) error {
	return c.String(http.StatusOK, "hello,world")
}

// 定义一个结构体接收请求参数
type loginReq struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

// 定义一个结构体返回响应结果
type loginRes struct {
	Token string `json:"token"`
}

func login(c echo.Context) error {
	req := &loginReq{}

	// 绑定请求参数值
	c.Bind(req)
	fmt.Println(req.Username)
	fmt.Println(req.Password)

	res := &loginRes{}

	// 判断用户名密码是否正确
	if req.Username == "root" && req.Password == "123456" {

		// 正确返回结果
		res.Token = "abcd"
		return c.JSON(200, res)

	} else {
		// 否则返回错误
		return errors.New("用户名密码错误")
	}
}

type Movie struct {
	Id          int64  `json:"id"`          // 电影id
	Title       string `json:"title"`       // 电影标题
	Description string `json:"description"` // 电影介绍
}

type moviesRes struct {
	Movies []*Movie `json:"movies"` // 电影列表
}

var moviesMap = map[int64]*Movie{
	1: &Movie{
		Id:          1,
		Title:       "猪猪波",
		Description: "是头猪",
	},
	2: &Movie{
		Id:          2,
		Title:       "钢铁",
		Description: "一度量",
	},
}

func movies(c echo.Context) error {
	res := &moviesRes{}

	for _, value := range moviesMap {
		res.Movies = append(res.Movies, value)
	}

	/*
		res.Movies = append(res.Movies, &Movie{
			Id:          1,
			Title:       "猪猪波",
			Description: "是头猪",
		})
		res.Movies = append(res.Movies, &Movie{
			Id:          2,
			Title:       "钢铁",
			Description: "一度量",
		})
	*/

	return c.JSON(200, res)
}

func movie(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	res := moviesMap[id]
	return c.JSON(200, res)
}

func main() {
	// 调用echo通过New赋值给e
	e := echo.New()

	e.Static("/", "static")

	e.POST("/api/login", login)

	e.GET("/api/movies", movies)

	e.GET("/api/movies/:id", movie)
	//e使用echo里的GET方法，调用上面 s1 与 s2
	e.GET("/s1", s1)
	// e调用方法，开始启动监听下面的端口号
	e.Logger.Fatal(e.Start("0.0.0.0:1323"))
}
