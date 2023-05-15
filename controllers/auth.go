package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlShortener/models"
)

type LoginInput struct {
	login    string
	password string
}

func Register(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{Login: input.login, Password: input.password}

	err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := models.LoginCheck(input.login, input.password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

type ShortenInput struct {
	url string
}

func Shorten(c *gin.Context) {
	var input ShortenInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortenUrl, err := models.ShortenUrl(input.url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to shorten url"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": shortenUrl})
}

func Redirect(c *gin.Context) {
	shortCode := c.Param("shortCode")
	if shortCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid url"})
	}

	fullUrl, err := models.GetFullUrl(shortCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to find full url"})
	}

	c.Redirect(http.StatusOK, fullUrl)
}
