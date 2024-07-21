package handler

import "github.com/labstack/echo/v4"

type moviesRes struct {
	Movies []*Movie `json:"movies"` // 电影列表
}

var moviesMap = map[int64]*Movie{
	1: &Movie{
		Id:          1,
		Title:       "猪猪波",
		Description: "是头猪",
	},
	2: &Movie{
		Id:          2,
		Title:       "钢铁",
		Description: "一度量",
	},
}

func HandleMovies(c echo.Context) error {
	res := &moviesRes{}

	for _, value := range moviesMap {
		res.Movies = append(res.Movies, value)
	}

	/*
		res.Movies = append(res.Movies, &Movie{
			Id:          1,
			Title:       "猪猪波",
			Description: "是头猪",
		})
		res.Movies = append(res.Movies, &Movie{
			Id:          2,
			Title:       "钢铁",
			Description: "一度量",
		})
	*/

	return c.JSON(200, res)
}
