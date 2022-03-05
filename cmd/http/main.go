package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	handlers "github.com/jaysyanshar/go-calculator/handlers/http"
)

func main() {
	e := echo.New()
	handlers.Init()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	// users
	e.POST("/users", handlers.CreateUser)
	e.GET("/users/:id", handlers.GetUser)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.DELETE("/users/:id", handlers.DeleteUser)
	// math
	e.POST("/math/add", handlers.Add)
	e.POST("/math/sub", handlers.Sub)
	e.POST("/math/mul", handlers.Mul)
	e.POST("/math/div", handlers.Div)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
