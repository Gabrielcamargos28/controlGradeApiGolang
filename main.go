package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gradeControleApi/configuration/database"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading .env")
	}
	err = database.Connect()

	if err != nil {
		log.Fatal("Error ao conectar no banco:", err)
	}
	defer database.Close()

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "1323"
	}
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello Gabis")
	})
	apiAdress := fmt.Sprintf("localhost:%s", port)

	e.Logger.Fatal(e.Start(apiAdress))
}
