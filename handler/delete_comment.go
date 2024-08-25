package handler

import (
	"log"
	"strconv"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type deleteCommentReq struct {
}

type deleteCommentRes struct {
	Success bool `json:"success"`
}

// 定义一个接收参数，并返回结果
func HandleDeleteComment(c echo.Context) error {
	req := &deleteCommentReq{}

	// 绑定请求参数
	c.Bind(req)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	// 删除数据库里的数据
	r, err := dbutil.DB.Exec("delete from comments where id = ?", id)
	if err != nil {
		log.Print(err)
		return err
	} else {
		log.Print(r)
	}

	res := &deleteCommentRes{
		Success: true,
	}
	//正确返回相应结果
	return c.JSON(200, res)
}
