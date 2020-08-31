package main

import (
	"errors"
	"jpw547/byulib-techinical-challenge/handlers"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.POST("/posts", handlers.AddNewPost)
	e.GET("/posts", handlers.GetAllPosts)
	e.GET("/posts/:title", handlers.GetPostByTitle)

	err := e.Start(":1600")
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("failed to start server", err.Error())
	}
}
