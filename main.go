package main

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vaeho/go-todo/templates"
	"github.com/vaeho/go-todo/templates/components"
)

func main() {
	e := echo.New()

	component := templates.Index()
	e.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.GET("/components", func(c echo.Context) error {
		t := c.QueryParam("type")
		switch t {
		case "add-todo":
			component := components.AddTodoInput()
			return component.Render(context.Background(), c.Response().Writer)
		case "add-todo-btn":
			component := components.AddTodoButton()
			return component.Render(context.Background(), c.Response().Writer)

		}
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Element")
	})
	e.Static("/static", "static")
	e.Static("/css", "css")
	e.Static("/fonts", "fonts")
	e.Logger.Fatal(e.Start(":1323"))
}
