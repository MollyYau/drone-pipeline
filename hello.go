package main

import (
        "net/http"

        "github.com/labstack/echo"
        "github.com/labstack/echo/middleware"
)

func main() {
        e := echo.New()
        e.GET("/", func(c echo.Context) error {
                return c.String(http.StatusOK, "Hello, World!")
        })

        e.GET("/ping", func(c echo.Context) error {
                return c.String(http.StatusOK, "pong")
        })

        e.Use(middleware.Logger())

        e.Logger.Debug(e.Start(":1323"))
}