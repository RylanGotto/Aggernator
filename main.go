package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	path_dir, err := os.Getwd()
	_ = godotenv.Load(filepath.Join(path_dir, ".env"))

	if err != nil {
		log.Fatal()
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

	e.Logger.Fatal(e.Start(":8000"))
}
