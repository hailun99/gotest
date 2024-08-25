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
	e.POST("/api/ragiser", handler.HanbleRagiser)
	e.POST("/api/modifypassword", handler.HandleModifyRagiser)

	//
	e.GET("/api/movies", handler.HandleMovies)
	e.POST("/api/movies", handler.HandleAddMovie)
	e.GET("/api/movies/:id", handler.HandleMovie)
	e.PUT("/api/movies/:id", handler.HandleEditMovie)
	e.DELETE("/api/movies/:id", handler.HandleDeleteMovie)

	// 添加评论
	e.POST("/api/addcomment", handler.HanbleComments)

	// 删除评论（根据评论id）
	e.DELETE("/api/deletecomment/:id", handler.HandleDeleteComment)

	// 修改评论（根据 id 修改评分）
	e.POST("/api/HandleEditComment/:id", handler.HandleEditComment)

	// 查询一个电影的所有评论（根据电影 id）
	e.GET("/api/sltcomments/:id", handler.HandlSltCommentsRes)

	// e调用方法，开始启动监听下面的端口号
	e.Logger.Fatal(e.Start("0.0.0.0:1323"))

}
