package main

import (
	_ "embed"
	"github.com/joho/godotenv"
	"pendaftaran-sidang/internal/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app.StartApp()
}
