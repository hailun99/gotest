package handler

import (
	"strconv"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

// 定义一个请求参数
type Commentsres struct {
	// Username string `json:"username"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Comments []*CommentsReq `json:"comments"`
	} `json:"data"`
}

func HandlUsernameCommentsRes(c echo.Context) error {
	//用一个匿名函数来接收id, 一个长度为10容量为64的变量
	username, _ := strconv.ParseInt(c.QueryParam("username"), 10, 64)

	// 初始化Movie
	res := &Commentsres{}
	// 通过username查询comments里的数据
	rows, err := dbutil.DB.Query("select username, movieid, score, comment from comments where username = ?", username)
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
}
