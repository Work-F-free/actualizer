package main

import (
	_ "actualizer/docs"
	"actualizer/internal/app"
)

// @title           Actualizer
// @version         1.0
// @description     Service for storing Coworking Plan Data

// @contact.name   Koreshkov Daniil
// @contact.email  danielkoreshkov@gmail.com

// @BasePath /
func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
