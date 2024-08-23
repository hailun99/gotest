package handler

import (
	"log"
	"time"

	"github.com/hailun99/gotest/dbutil"
	"github.com/labstack/echo/v4"
)

type addRagiserReq struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type addRagiserRes struct {
	Success bool `json:"success"`
}

func HanbleAddRagiser(c echo.Context) error {
	req := &addRagiserReq{}

	c.Bind(req)

	r, err := dbutil.DB.Exec("INSERT INTO users(username, password) VALUES (?, ?)",
		req.Username, req.Password, time.Now().Unix())
	if err != nil {
		return err
	} else {
		log.Println(r)
	}

	res := &addRagiserRes{
		Success: true,
	}
	return c.JSON(200, res)

}
