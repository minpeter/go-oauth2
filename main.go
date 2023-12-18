package main

import (
	"github.com/gin-gonic/gin"
	"github.com/minpeter/go-oauth2/config"
	"github.com/minpeter/go-oauth2/controllers"
	"github.com/minpeter/go-oauth2/database"
)

func main() {

	database.InitDB()

	app := gin.Default()

	config.GithubConfig()
	{
		app.GET("/github_login", controllers.GithubLogin)
		app.GET("/github_callback", controllers.GithubCallback)
	}

	{
		app.GET("/", controllers.Home)
		app.GET("/user", controllers.User)
	}

	app.Run(":8080")
}
