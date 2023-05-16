package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlShortener/auth"
	"urlShortener/models"
)

func Shorten(c *gin.Context) {
	var input ShortenInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("Authorization")
	user, err := auth.ExtractJWTTokenUser(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to parse user id"})
		return
	}

	shortenUrl, err := models.ShortenUrl(input.url, user)
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

func UrlsList(c *gin.Context) {
	token := c.GetHeader("Authorization")
	user, err := auth.ExtractJWTTokenUser(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to parse user id"})
		return
	}
	urls, err := models.GetUsersUrls(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to find user's urls"})
	}
	//todo check how it is serialised
	c.JSON(http.StatusOK, gin.H{"urls": urls})
}
