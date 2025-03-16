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
	// 查看登录
	e.POST("/api/login", handler.HandleLogin)
	// 添加用户
	e.POST("/api/ragiser", handler.HanbleRagiser)
	// 修改用户密码
	e.POST("/api/modifypassword", handler.HandleModifyRagiser)

	// 查询电影列表
	e.GET("/api/movies", handler.HandleMovies)
	// 添加电影
	e.POST("/api/add_movies", handler.HandleAddMovie)
	// 通过id查询某部电影
	e.GET("/api/queryRow_movies/:id", handler.HandleMovie)
	// 通过id修改某部电影
	e.PUT("/api/modify_movies/:id", handler.HandleEditMovie)
	// 通过id删除某部电影
	e.DELETE("/api/movies/:id", handler.HandleDeleteMovie)

	// 添加评论
	e.POST("/api/add_comment", handler.HanbleComments)
	// 删除评论（根据评论id）
	e.DELETE("/api/comment/:id", handler.HandleDeleteComment)
	// 修改评分（根据 id 修改评分）
	e.POST("/api/modify_comment", handler.HandlemodifyComment)
	// 查询一个电影的所有评论（根据电影 id）
	e.GET("/api/movie_comments", handler.HandlMovieComments)
	// 查询一个用户发表的所有评论
	e.GET("/api/username_movie_comments", handler.HandlUsernameCommentsRes)

	// 添加电影票
	e.POST("/api/add_tickek", handler.HandleAddTickek)
	//通过id删除电影票
	e.DELETE("/api/tickek/:id", handler.HandleDeleteTickek)
	// 通过id修改电影
	e.POST("/api/modify_tickek/:id", handler.HandleEditTickek)
	// 通过顾客id查看电影票
	e.GET("/api/tickek/:id", handler.HandleTickek)
	//查询电影票
	e.GET("/api/tickeks", handler.HandleTickeks)

	// //通过用户名查询他的电影票记录
	// e.GET("/api/recovds_tickek", handler.HandleRecovds)
	//

	// e调用方法，开始启动监听下面的端口号
	e.Logger.Fatal(e.Start("0.0.0.0:1323"))

}
