package handler

import (
	"log"
	"strconv"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type deleteMovieReq struct {
}

type deleteMovieRes struct {
	Success bool `json:"success"`
}

func HandleDeleteMovie(c echo.Context) error {
	req := &deleteMovieReq{}

	c.Bind(req)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	r, err := dbutil.DB.Exec("delete from movies where id = ?", id)
	if err != nil {
		return err
	} else {
		log.Print(r)
	}

	res := &deleteMovieRes{
		Success: true,
	}
	//正确返回相应结果
	return c.JSON(200, res)
}
