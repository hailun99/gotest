package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type Movie struct {
	Id          int64  `json:"id"`          // 电影id
	Title       string `json:"title"`       // 电影标题
	Description string `json:"description"` // 电影介绍
}

func HandleMovie(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	res := moviesMap[id]
	return c.JSON(200, res)
}
