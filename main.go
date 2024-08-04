package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func master(c echo.Context) error {
	master := c.Param("master")
	return c.String(http.StatusOK, master)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "http://localhost:8080/ping")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/pong", func(c echo.Context) error {
		return c.String(http.StatusOK, "ping")
	})

	e.GET("/users/:id", master)
	e.Logger.Fatal(e.Start(":8080"))
}
