package handler

import (
	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

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

	rows, err := dbutil.DB.Query("select id, title, description, created from movies")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		mov := &Movie{}
		err = rows.Scan(&mov.Id, &mov.Title, &mov.Description, &mov.Created)
		if err != nil {
			return err
		}
		res.Movies = append(res.Movies, mov)
	}

	// // 遍历Map
	// for _, value := range moviesMap {
	// 	res.Movies = append(res.Movies, value)
	// }

	return c.JSON(200, res)
}
