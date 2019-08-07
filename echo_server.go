package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		log.Println("/ been visited at:", time.Now())
		return c.String(http.StatusOK, "Hello, World! version 3")
	})
	e.GET("/headers", func(c echo.Context) error {
		log.Println("/headers been visited")
		for k, v := range c.Request().Header {
			log.Println(k, "=>", v)
		}
		return c.String(http.StatusOK, "headers")
	})
	e.GET("/proxy", func(c echo.Context) error {
		log.Println("/proxy been visited")
		resp, err := http.Get("http://echo-server-svc:1323/")
		if err != nil {
			log.Println("got error", err)
		}
		return c.String(http.StatusOK, fmt.Sprintln("proxy", resp.Status))
	})

	e.Logger.Fatal(e.Start(":1323"))
}
