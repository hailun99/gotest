package handler

import (
	"log"
	"strconv"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type editBookReq struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// 返回一个相应结果
type editBookRes struct {
	Success bool `json:"success"`
}

// 定义一个接收参数，并返回结果
func HandleEditBook(c echo.Context) error {
	req := &editBookReq{}

	// 绑定请求参数
	c.Bind(req)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	// 更新数据
	r, err := dbutil.DB.Exec("update books set Author = ?, Title = ? where id = ?",
		req.Author, req.Title, id)
	if err != nil {
		return err
	} else {
		log.Print(r)
	}

	res := &editBookRes{
		Success: true,
	}
	//正确返回相应结果
	return c.JSON(200, res)
}
