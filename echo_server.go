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
		return c.String(http.StatusOK, "Hello, World! version 1")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
