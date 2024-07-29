package handler

import (
	"log"
	"time"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type addMovieReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type addMovieRes struct {
	Success bool `json:"success"`
}

func HandleAddMovie(c echo.Context) error {
	req := &addMovieReq{}

	c.Bind(req)

	r, err := dbutil.DB.Exec("INSERT INTO movies (title, description, created) VALUES (?, ?, ?)",
		req.Title, req.Description, time.Now().Unix())
	if err != nil {
		return err
	} else {
		log.Print(r)
	}

	// moviesMap[3] = &Movie{
	// 	Id:          3,
	// 	Title:       req.Title,
	// 	Description: req.Description,
	// }

	res := &addMovieRes{
		Success: true,
	}
	return c.JSON(200, res)
}
