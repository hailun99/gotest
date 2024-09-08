package handler

import (
	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type moviesRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Movies []*Movie `json:"movies"` // 电影列表
	} `json:"data"`
}

// var moviesMap = map[int64]*Movie{
// 	1: &Movie{
// 		Id:          1,
// 		Title:       "猪猪波",
// 		Description: "是头猪",
// 	},
// 	2: &Movie{
// 		Id:          2,
// 		Title:       "钢铁",
// 		Description: "一度量",
// 	},
// }

// 定义一个接收变量
func HandleMovies(c echo.Context) error {
	res := &moviesRes{}
	// 查询movies里的全部数据
	rows, err := dbutil.DB.Query("select id, title, description, created, directo, performer from movies")
	if err != nil {
		res.Code = 100020
		res.Msg = err.Error()
		return c.JSON(200, res)
	}
	defer rows.Close() // 推迟最后执行并退出

	// 遍历多条数据库
	for rows.Next() {
		mov := &Movie{}
		err = rows.Scan(&mov.Id, &mov.Title, &mov.Description, &mov.Created, &mov.Directo, &mov.Performer)
		if err != nil {
			res.Code = 100010
			res.Msg = err.Error()
			return c.JSON(200, res)
		}
		//
		res.Data.Movies = append(res.Data.Movies, mov)
	}

	// // 遍历Map
	// for _, value := range moviesMap {
	// 	res.Movies = append(res.Movies, value)
	// }

	res.Code = 0
	res.Msg = "ok"
	return c.JSON(200, res)
}
