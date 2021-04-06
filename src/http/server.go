package http

import (
	"net/http"
	"strconv"

	// "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

type WebServer struct {
}

func NewWebserver() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve(porta int) {
	e := echo.New()
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	// }))

	e.Static("/imgs", "imgs") // funcional
	// e.Static("/imgs1", "./imgs") // errado
	// e.Static("/imgs2", "./../imgs") // errado
	// e.Static("/imgs3", "./../../imgs") // errado
	// e.Static("/imgs4", "../imgs") // errado
	// e.Static("/imgs5", "../../imgs") // errado
	// e.Static("/src", "t1") // errado
	// e.Static("/t2", "src") // funcional

	e.GET("/", w.getHome)
	e.GET("/teste", w.getTeste)
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(porta))) // launch server
}

func (w WebServer) getTeste(c echo.Context) error {
	return c.JSON(http.StatusOK, "Teste")
}

func (w WebServer) getHome(c echo.Context) error {
	var html string

	html =
		"<style>" +
			"body {" +
			"		background-image: url(\"imgs/background.jpg\");" +
			"		height: 100%;" +
			"		background-position: center;" +
			"		background-repeat: no-repeat;" +
			"		background-size: cover;" +
			"  }" +
			".center {" +
			"	margin: 0;" +
			"   position: absolute;" +
			"	top: 50%;" +
			"	left: 50%;" +
			"	-ms-transform: translate(-50%, -50%);" +
			"	transform: translate(-50%, -50%);" +
			"}" +
			"</style>" +
			"<h1 class=\"center\"><center>Imersão Full Cycle</center></h1>"

	return c.HTML(http.StatusOK, html)
}
