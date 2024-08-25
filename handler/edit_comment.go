package handler

import (
	"log"
	"strconv"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type editCommentReq struct {
	Score    int64  `josn:"score"` // 评分
	Oldscore string `json:"oldscore"`
}

// 返回一个相应结果
type editCommenteRes struct {
	Success bool `json:"success"`
}

// 定义一个接收参数，并返回结果
func HandleEditComment(c echo.Context) error {
	req := &editCommentReq{}

	// 绑定请求参数
	c.Bind(req)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	// 更新数据
	r, err := dbutil.DB.Exec("update comments set score = ?, score = ? where id = ?",
		req.Score, req.Oldscore, id)
	if err != nil {
		log.Println("无法查询")
		return err
	} else {
		log.Print(r)
	}

	res := &editCommenteRes{
		Success: true,
	}
	//正确返回相应结果
	return c.JSON(200, res)
}
