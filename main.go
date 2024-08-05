package main

import (
	"log"

	"github.com/hailun99/gotest/dbutil"
	"github.com/hailun99/gotest/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	err := dbutil.Init() //调用输出化数据库的函数
	if err != nil {
		log.Fatal(err)
	}
	err = dbutil.DB.Ping() //验证用户名与密码,尝试链接数据库
	if err != nil {
		log.Fatal(err)
	}

	// 调用echo通过New赋值给e
	e := echo.New()

	e.Static("/", "static")

	e.POST("/api/login", handler.HandleLogin)

	//
	e.GET("/api/movies", handler.HandleMovies)
	e.POST("/api/movies", handler.HandleAddMovie)

	e.GET("/api/movies/:id", handler.HandleMovie)

	e.PUT("/api/movies/:id", handler.HandleEditMovie)

	e.DELETE("/api/movies/:id", handler.HandleDeleteMovie)

	// e调用方法，开始启动监听下面的端口号
	e.Logger.Fatal(e.Start("0.0.0.0:1323"))
}
