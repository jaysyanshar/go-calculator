package http

import (
	"github.com/jaysyanshar/go-calculator/entities"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Add(c echo.Context) error {
	return doCalculation(c, mathService.Add)

}

func Sub(c echo.Context) error {
	return doCalculation(c, mathService.Sub)

}

func Mul(c echo.Context) error {
	return doCalculation(c, mathService.Mul)

}

func Div(c echo.Context) error {
	return doCalculation(c, mathService.Div)
}

func doCalculation(c echo.Context, handler func(a, b float64) float64) error {
	math := &entities.Math{}
	if err := c.Bind(math); err != nil {
		return err
	}
	res := handler(math.A, math.B)
	return c.JSON(http.StatusOK, res)
}
