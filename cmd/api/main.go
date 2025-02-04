package main

import (
	_ "seatPlanner/docs"
	"seatPlanner/internal/app"
)

// @title           Actualizer
// @version         1.0
// @description     Service for storing Coworking Plan Data

// @contact.name   Koreshkov Daniil
// @contact.email  danielkoreshkov@gmail.com

// @host      localhost:8081
// @BasePath /
func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
