package handler

import (
	"log"
	"time"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type CommentsReq struct {
	Username string `json:"username"`
	Movieid  int64  `json:"movieid"`
	Score    int64  `json:"score"`
	Comment  string `json:"comment"`
}

type CommentsRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func HanbleComments(c echo.Context) error {
	req := &CommentsReq{}

	c.Bind(req)

	res := &CommentsRes{}

	if req.Score > 5 || req.Score <= 0 {
		res.Code = 10001
		res.Msg = "分数不能大于5 或 小于等于0"
		return c.JSON(200, res)
	}

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
