package handler

import (
	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type tickeksRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Tickeks []*Tickek `json:"tickeks"` // 电影票
	} `json:"data"`
}

// 定义一个接收变量
func HandleTickeks(c echo.Context) error {
	res := &tickeksRes{}
	// 查询movies里的全部数据
	rows, err := dbutil.DB.Query("select id, ciname, movie, type, seat, price, created from tickek")
	if err != nil {
		res.Code = 100020
		res.Msg = err.Error()
		return c.JSON(200, res)
	}
	defer rows.Close() // 推迟最后执行并退出

	// 遍历多条数据库
	for rows.Next() {
		tic := &Tickek{}
		err = rows.Scan(&tic.Id, &tic.Ciname, &tic.Movie, &tic.Type, &tic.Seat, &tic.Price, &tic.Created)
		if err != nil {
			res.Code = 100010
			res.Msg = err.Error()
			return c.JSON(200, res)
		}
		//
		res.Data.Tickeks = append(res.Data.Tickeks, tic)
	}

	// // 遍历Map
	// for _, value := range moviesMap {
	// 	res.Movies = append(res.Movies, value)
	// }

	res.Code = 0
	res.Msg = "ok"
	return c.JSON(200, res)
}
