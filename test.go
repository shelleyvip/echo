package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main()  {
	e := echo.New()
	e.GET("/" , func( c  echo.Context) error {
		return c.String(http.StatusOK, "Hello, 周雪莉你好")
	})
	//在1323号端口开启服务
	e.Logger.Fatal(e.Start(":9999"))
}

