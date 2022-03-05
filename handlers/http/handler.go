package http

import (
	"github.com/jaysyanshar/go-calculator/entities"
	"github.com/jaysyanshar/go-calculator/services"
)

var (
	mathService services.MathService
	users       map[int]*entities.User
	seq         int
)

func Init() {
	mathService = &services.MathServiceImpl{}
	users = map[int]*entities.User{}
	seq = 1
}
