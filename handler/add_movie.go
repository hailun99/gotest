package handler

import "github.com/labstack/echo/v4"

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
	moviesMap[3] = &Movie{
		Id:          3,
		Title:       req.Title,
		Description: req.Description,
	}

	res := &addMovieRes{
		Success: true,
	}
	return c.JSON(200, res)
}
