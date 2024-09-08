package handler

import (
	"log"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type editCommentReq struct {
	Id    int64 `json:"id"`
	Score int64 `json:"score"` // 评分
}

// 返回一个相应结果
type editCommenteRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 定义一个接收参数，并返回结果
func HandlemodifyComment(c echo.Context) error {
	req := &editCommentReq{}

	// 绑定请求参数
	c.Bind(req)

	res := &editCommenteRes{}

	// 评分不能大于5且不能小于等于0
	if req.Score > 5 || req.Score <= 0 {
		res.Code = 10001
		res.Msg = "分数不能大于5 或 小于等于0"
		return c.JSON(200, res)
	}

	// 通过id更新数据
	r, err := dbutil.DB.Exec("update comments set score = ? where id = ?",
		req.Score, req.Id)
	if err != nil {
		res.Code = 100010
		res.Msg = err.Error()
		return c.JSON(200, res)
	} else {
		log.Print(r)
	}

	res.Code = 0
	res.Msg = "ok"
	//正确返回相应结果
	return c.JSON(200, res)
}
