package handlers

import (
	"fmt"
	"jpw547/byulib-techinical-challenge/database"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllPosts(ctx echo.Context) error {
	results, err := database.GetAllPosts()
	if err != nil {
		fmt.Printf("Failed to get all posts from the database: %s", err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, results)
}

func GetPostByTitle(ctx echo.Context) error {
	postTitle := ctx.Param("title")

	post, err := database.GetPostByTitle(postTitle)
	if err != nil {
		fmt.Printf("failed to get the post with the specified title - %s: %s", postTitle, err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, post)
}

func AddNewPost(ctx echo.Context) error {
	var blogPost database.BlogPost
	err := ctx.Bind(&blogPost)
	if err != nil {
		fmt.Printf("Failed to bind blog post data from the body: %s", err.Error())
		return ctx.JSON(http.StatusBadRequest, err)
	}

	err = database.AddNewPost(blogPost)
	if err != nil {
		fmt.Printf("Failed to add the new blog post: %s", err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.String(http.StatusOK, "OK")
}
