package handler

import (
	"log"
	"time"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type CommentsReq struct {
	Username string `josn:"username"`
	Movieid  int64  `josn:"movieid"`
	Score    int64  `josn:"score"`
	Comment  string `josn:"comment"`
}

type CommentsRes struct {
	Success bool `json:"success"`
}

func HanbleComments(c echo.Context) error {
	req := &CommentsReq{}

	c.Bind(req)

	r, err := dbutil.DB.Exec("INSERT INTO comments(username, movieid, score, comment, created) VALUES (?, ?, ?, ?, ?)",
		req.Username, req.Movieid, req.Score, req.Comment, time.Now().Unix())
	if err != nil {
		log.Print(err)
		return err
	} else {
		log.Println(r)
	}

	res := &CommentsRes{
		Success: true,
	}
	return c.JSON(200, res)

}
