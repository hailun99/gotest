package handler

import (
	"log"
	"strconv"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type editMovieReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type editMovieRes struct {
	Success bool `json:"success"`
}

func HandleEditMovie(c echo.Context) error {
	req := &editMovieReq{}

	c.Bind(req)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	r, err := dbutil.DB.Exec("update movies set title = ?, description = ? where id = ?",
		req.Title, req.Description, id)
	if err != nil {
		return err
	} else {
		log.Print(r)
	}

	res := &editMovieRes{
		Success: true,
	}
	//正确返回相应结果
	return c.JSON(200, res)
}
