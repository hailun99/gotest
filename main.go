package main

import (
	"github.com/hailun99/gotest/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	// 调用echo通过New赋值给e
	e := echo.New()

	e.Static("/", "static")

	e.POST("/api/login", handler.HandleLogin)

	//
	e.GET("/api/movies", handler.HandleMovies)
	e.POST("/api/movies", handler.HandleAddMovie)

	e.GET("/api/movies/:id", handler.HandleMovie)

	// e调用方法，开始启动监听下面的端口号
	e.Logger.Fatal(e.Start("0.0.0.0:1323"))
}
