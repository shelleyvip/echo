package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main()  {
	e :=  echo.New()   //开启一个服务
	e.Use(middleware.Logger()) //日志

	e.Use(middleware.Recover()) //回复


	//路由
	e.GET("/h1",hello01)
	e.GET("/h2",hello02)


	//开启服务和监听端口
	e.Logger.Fatal(e.Start(":1314"))

}



func hello01(c echo.Context)  error {
		return  c.String(http.StatusOK ,"hello01")

}
func hello02(c echo.Context)  error {
	return  c.String(http.StatusOK ,"hello02")

}