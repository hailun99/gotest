package handler

import (
	"strconv"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type selectCommentsRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Comments []*CommentsReq `json:"comments"`
	} `json:"data"`

	// Username string `josn:"username"`
	// // Movieid  int64  `josn:"movieid"`
	// Score   int64  `josn:"score"`
	// Comment string `josn:"comment"`
}

func HandlMovieCommentsRes(c echo.Context) error {
	//用一个匿名函数来接收id, 一个长度为10容量为64的变量
	movieid, _ := strconv.ParseInt(c.QueryParam("movieid"), 10, 64)

	// 初始化Movie
	res := &selectCommentsRes{}

	rows, err := dbutil.DB.Query("select username, movieid, score, comment from comments where movieid = ?", movieid)
	if err != nil {
		res.Code = 100010
		res.Msg = err.Error()
		return c.JSON(200, res)
	}
	defer rows.Close() // 推迟最后执行并退出

	// 遍历多条数据库
	for rows.Next() {
		com := &CommentsReq{}
		err = rows.Scan(&com.Username, &com.Movieid, &com.Score, &com.Comment)
		if err != nil {
			res.Code = 100010
			res.Msg = err.Error()
			return c.JSON(200, res)
		}
		res.Data.Comments = append(res.Data.Comments, com)
	}

	res.Code = 0
	res.Msg = "ok"
	return c.JSON(200, res)
	// 查询数据库里的信息并赋值
	// err := dbutil.DB.QueryRow("select username, score, comment from comments where movieid = ?", movieid).
	// 	Scan(&res.Username, &res.Score, &res.Comment)
	// if err != nil {
	// 	return err
	// }
	// return c.JSON(200, res)
}
