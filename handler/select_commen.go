package handler

import (
	"strconv"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type selectCommentsReq struct {
	Username string `josn:"username"`
	Movieid  int64  `josn:"movieid"`
	Score    int64  `josn:"score"`
	Comment  string `josn:"comment"`
}

func HandlSltCommentsRes(c echo.Context) error {
	//用一个匿名函数来接收id, 一个长度为10容量为64的变量
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	// 初始化Movie
	req := &selectCommentsReq{}

	// 查询数据库里的信息并赋值
	err := dbutil.DB.QueryRow("select username, movieid, score, comment from comments where id = ?", id).
		Scan(&req.Username, &req.Movieid, &req.Score, req.Comment)
	if err != nil {
		return err
	}
	// 把Map中的id赋值给res
	res := req
	return c.JSON(200, res)
}
