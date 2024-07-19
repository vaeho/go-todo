package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/vaeho/go-todo/templates"
)

func main() {
	e := echo.New()

	component := templates.Index()
	e.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.Static("/static", "static")
	e.Static("/css", "css")
	e.Static("/fonts", "fonts")
	e.Logger.Fatal(e.Start(":1323"))
}
