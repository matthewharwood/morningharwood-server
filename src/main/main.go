package main

import (
	"fmt"
	"flag"
	"net/http"
	"github.com/labstack/echo"
	"github.com/matthewharwood/morningharwood-server/src/api"
	"github.com/labstack/echo/middleware"
)

func todoistEndpoint(c echo.Context) error {
	return c.JSON(http.StatusOK, api.Todoist())
}


func main() {
	port := flag.String("port", ":8000", "server port")
	fmt.Println("start echo")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "morningharwood-client/assets/images")
	e.File("/favicon.ico", "morningharwood-client/assets/images/favicon/favicon.ico")
	e.File("/favicon-16x16.png", "morningharwood-client/assets/images/favicon/favicon-16x16.png")
	e.File("/favicon-32x32.png", "morningharwood-client/assets/images/favicon/favicon-32x32.png")
	e.File("/", "morningharwood-client/index.html")

	g := e.Group("/api/v1")
	g.Use(middleware.CORS())
	g.GET("/todoist", todoistEndpoint)
	e.Logger.Fatal(e.Start(*port))
}