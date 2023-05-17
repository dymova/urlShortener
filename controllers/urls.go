package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlShortener/auth"
	"urlShortener/models"
)

type ShortenInput struct {
	Url string `json:"url" binding:"required"`
}

func Shorten(c *gin.Context) {
	var input ShortenInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := getUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to parse user id"})
		return
	}

	shortenUrl, err := models.ShortenUrl(input.Url, user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to shorten url"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": shortenUrl})
}

func getUser(c *gin.Context) (models.User, error) {
	token := c.GetHeader("Authorization")
	userId, err := auth.ExtractJWTTokenUser(token)
	if err != nil {
		var user models.User
		return user, err
	}
	return models.GetUser(userId)
}

func Redirect(c *gin.Context) {
	shortCode := c.Param("shortCode")
	if shortCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid url"})
		return
	}

	fullUrl, err := models.GetFullUrl(shortCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to find full url"})
		return
	}

	c.Redirect(http.StatusFound, fullUrl)
}

func UrlsList(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	urls, err := models.GetUsersUrls(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to find user's urls"})
	}
	//todo check how it is serialised
	c.JSON(http.StatusOK, gin.H{"urls": urls})
}
