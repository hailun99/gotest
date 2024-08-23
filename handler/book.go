package handler

import (
	"strconv"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type Book struct {
	Id      int64  `json:"id"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Created int64  `json:"created"`
}

func HanbleBook(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	req := &Book{}

	err := dbutil.DB.QueryRow("selet id, author, title from books where id = ?", id).
		Scan(&req.Id, &req.Author, &req.Title)
	if err != nil {
		return err
	}
	book := req
	return c.JSON(200, book)
}
