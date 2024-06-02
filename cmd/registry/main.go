package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/portanic/registry/internal/handlers"
)

func main() {
	e := echo.New()

	e.POST("/plugins", handlers.RegisterPlugin)
	e.GET("/plugins", handlers.ListPlugins)
	e.GET("/plugins/:id", handlers.GetPlugin)

	fmt.Println("Server is running on http://localhost:8081")
	e.Logger.Fatal(e.Start(":8081"))
}
