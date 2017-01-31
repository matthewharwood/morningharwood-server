package main

import (
	"fmt"
	"flag"
	"net/http"
	"github.com/labstack/echo"
)

func landing(c echo.Context) error {
	return c.String(http.StatusOK, "yallo from echo bro")
}

func admin(c echo.Context) error {
	adminName := c.QueryParam("name")
	adminType := c.QueryParam("type")

	return c.JSON(http.StatusOK, map[string]string{
			"name": adminName,
			"type": adminType,
	})
}

func main() {
	port := flag.String("port", ":8000", "server port")
	fmt.Println("start echo")
	e := echo.New()
	e.Static("/", "morningharwood-client/assets/images")
	e.File("/favicon.ico", "morningharwood-client/assets/images/favicon/favicon.ico")
	e.File("/favicon-16x16.png", "morningharwood-client/assets/images/favicon/favicon-16x16.png")
	e.File("/favicon-32x32.png", "morningharwood-client/assets/images/favicon/favicon-32x32.png")
	e.File("/", "morningharwood-client/index.html")
	//e.GET("/", landing)
	//e.GET("/admin", admin)

	e.Start(*port)
}