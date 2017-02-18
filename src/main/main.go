package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/matthewharwood/morningharwood-server/src/api"
	"github.com/labstack/echo/middleware"
	"flag"
	"fmt"
)

func todoistEndpoint(c echo.Context) error {
	return c.JSON(http.StatusOK, api.Todoist())
}

type (
	Host struct {
		Echo *echo.Echo
	}
)

func main() {
	// POINTERS
	port := flag.String("port", ":8080", "server port")
	domain := flag.String("domain", "localhost", "server domain")
	hosts := make(map[string]*Host)


	//------
	// store
	//------

	store := echo.New()
	store.Use(middleware.Logger())
	store.Use(middleware.Recover())

	hosts[fmt.Sprintf("store.%v%s",*domain, *port)] = &Host{store}

	store.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "store")
	})

	//---------
	// Website
	//---------

	site := echo.New()
	site.Static("/", "go/src/github.com/matthewharwood/morningharwood-client/client/dist/")
	site.File("/favicon.ico", "go/src/github.com/matthewharwood/morningharwood-client/assets/images/favicon/favicon.ico")
	site.File("/favicon-16x16.png", "go/src/github.com/matthewharwood/morningharwood-client/assets/images/favicon/favicon-16x16.png")
	site.File("/favicon-32x32.png", "go/src/github.com/matthewharwood/morningharwood-client/assets/images/favicon/favicon-32x32.png")
	site.File("/", "go/src/github.com/matthewharwood/morningharwood-client/client/dist/index.html")
	site.Use(middleware.Logger())
	site.Use(middleware.Recover())

	hosts[fmt.Sprintf("%v%s", *domain, *port)] = &Host{site}


	//-----
	// API
	//-----
	g := site.Group("/api/v1")
	g.Use(middleware.CORS())
	g.GET("/todoist", todoistEndpoint)

	// Server
	e := echo.New()
	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		host := hosts[req.Host]

		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}

		return
	})
	e.Logger.Fatal(e.Start(*port))
}