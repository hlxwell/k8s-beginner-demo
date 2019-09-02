package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		total := 0
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 1000; i++ {
			total += rand.Intn(1000)
		}
		log.Println("/ been visited at:", time.Now(), "-", total)
		return c.String(http.StatusOK, "Hello, World! version 1")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
