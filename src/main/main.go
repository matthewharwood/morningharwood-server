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
/**
 * <directory_to_binary> -port=:8080 -domain=104.198.5.254 -dir=go/src/github.com/matthewharwood/
 */
func main() {
	// flags.
	port := flag.String("port", ":8080", "server port")
	domain := flag.String("domain", "localhost", "server domain")
	dir := flag.String("static", "go/src/github.com/matthewharwood/morningharwood-client/client/dist", "static directory");
	hosts := make(map[string]*Host)
	flag.Parse()

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
	site.Static("/", fmt.Sprintf("%v", *dir))
	site.File("/", "go/src/github.com/matthewharwood/morningharwood-client/client/dist/index.html")
	site.File("/menu", "go/src/github.com/matthewharwood/morningharwood-client/client/dist/index.html")
	site.File("/work", "go/src/github.com/matthewharwood/morningharwood-client/client/dist/index.html")
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