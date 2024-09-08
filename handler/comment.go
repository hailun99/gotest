package handler

import (
	"log"
	"time"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

// 定义一个电影评论的请求参数
type CommentsReq struct {
	Username string `json:"username"` // 昵称
	Movieid  int64  `json:"movieid"`  // 电影id
	Score    int64  `json:"score"`    // 评分
	Comment  string `json:"comment"`  // 评论
}

// 返回结果
type CommentsRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 定义一个接收器并返回一个错误
func HanbleComments(c echo.Context) error {
	// 使用指针的方法进行赋值
	req := &CommentsReq{}
	// 绑定参数
	c.Bind(req)

	res := &CommentsRes{}
	// 评分不能大于5且不能小于等于0
	if req.Score > 5 || req.Score <= 0 {
		res.Code = 10001
		res.Msg = "分数不能大于5 或 小于等于0"
		return c.JSON(200, res)
	}
	// 向comments里添加数据
	r, err := dbutil.DB.Exec("INSERT INTO comments(username, movieid, score, comment, created) VALUES (?, ?, ?, ?, ?)",
		req.Username, req.Movieid, req.Score, req.Comment, time.Now().Unix())
	if err != nil {
		res.Code = 100010
		res.Msg = err.Error()
		return c.JSON(200, res)
	} else {
		log.Println(r)
	}

	res.Code = 0
	res.Msg = "ok"
	return c.JSON(200, res)

}
