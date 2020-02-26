package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

var (
	result string
)

func main() {
	e := echo.New()

	e.GET("/echo", func(c echo.Context) error {
		t1 := time.Now().Year()       //年
		t2 := time.Now().Month()      //月
		t3 := time.Now().Day()        //日
		t4 := time.Now().Hour()       //小时
		t5 := time.Now().Minute()     //分钟
		t6 := time.Now().Second()     //秒
		t7 := time.Now().Nanosecond() //纳秒
		fmt.Print("请求时间；")
		currentTimeData := time.Date(t1, t2, t3, t4, t5, t6, t7, time.Local) //获取当前时间，返回当前时间Time
		fmt.Println(currentTimeData)                                         //打印结果：2017-04-11 12:52:52.794351777 +0800 CST
		return c.String(http.StatusOK, "Echo Server - Hello, World!")
	})

	/* 1.路由*/
	/*127.0.0.1:1323/msg/123*/
	e.GET("/msg/:id", getParam)

	/* 2.请求参数*/
	/*127.0.0.1:1323/show?team=qztc&member=1182*/
	e.GET("/show", show)

	/*3.表单 application/x-www-form-urlencoded*/
	/*curl -F "name=1128" -F "email=1128@qztc.com" http://localhost:1323/save*/
	e.POST("/save", save)

	e.Logger.Fatal(e.Start(":1323"))
}

//e.GET("/param/:id", getParam)
func getParam(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// e.GET("/show", show)
func show(c echo.Context) error {
	// 从请求参数里获取 team 和 member 的值
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// e.POST("/save", save)
func save(c echo.Context) error {
	// 获取 name 和 email 的值
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
