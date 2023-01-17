package main

import (
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()

    e.GET("/hello", Hello())
    e.GET("/api/hello", ApiHelloGet())

    e.Start(":8080")
}

func Hello() echo.HandlerFunc {
    return func(c echo.Context) error {     
        return c.String(http.StatusOK, "4408: world is main")
    }
}
func ApiHelloGet() echo.HandlerFunc {
    return func(c echo.Context) error {     
        return c.String(http.StatusOK, "4408: Hello ver.2")
    }
}
