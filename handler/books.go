package handler

import (
	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type booksReq struct {
	Books []*Book `josn:"books"`
}

func HanbleBooks(c echo.Context) error {
	req := &booksReq{}

	rows, err := dbutil.DB.Query("select id, author, title from books")
	if err != nil {
		return err
	}
	defer rows.Close() // 推迟最后执行并退出

	// 遍历多条数据库
	for rows.Next() {
		bk := &Book{}
		err = rows.Scan(&bk.Id, &bk.Author, &bk.Title)
		if err != nil {
			return err
		}
		req.Books = append(req.Books, bk)
	}

	return c.JSON(200, req)
}
