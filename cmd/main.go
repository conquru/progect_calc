package main

import (
	"github.com/conquru/project_calc/internal/application"
)

func main() {
	app := application.New()
	app.RunServer()
}
