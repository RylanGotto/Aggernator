package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/sports", GameScoreHandler)
	e.GET("/news", NewsHandler)
	e.GET("/reddit", RedditHandler)

	e.File("/", "public/index.html")

	e.Logger.Fatal(e.Start(":8080"))
}
