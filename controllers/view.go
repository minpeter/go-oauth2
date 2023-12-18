package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/minpeter/go-oauth2/database"
)

func Home(c *gin.Context) {
	RenderTemplates(c, nil)
}

func User(c *gin.Context) {

	// 쿠키 확인
	token, err := c.Cookie("token")
	if err != nil {
		c.Redirect(302, "/")
		return
	}

	user, err := database.GetUserByToken(token)

	if err != nil {
		c.Redirect(302, "/")
		return
	}

	RenderTemplates(c, gin.H{
		"name":  user.Name,
		"email": user.Email,
	})
}
