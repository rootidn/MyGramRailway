package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func PhotoCreate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        Photo.ID,
		"caption":   Photo.Caption,
		"title":     Photo.Title,
		"photo_url": Photo.PhotoUrl,
		"user_id":   Photo.UserID,
	})
}

func PhotoGetAll(c *gin.Context) {
	db := database.GetDB()

	var Photo []models.Photo

	err := db.Select("id", "caption", "title", "photo_url", "user_id", "user").Preload("User").Find(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	for _, p := range Photo {
		if p.User != nil {
			p.User.CreatedAt = nil
			p.User.UpdatedAt = nil
			p.User.Age = 0
			p.User.Password = ""
			p.User.ProfileImageURL = ""
		}
	}

	c.JSON(http.StatusOK, Photo)
}

func PhotoGetById(c *gin.Context) {
	db := database.GetDB()
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	Photo := models.Photo{}
	Photo.ID = uint(photoId)

	err := db.Select("id", "caption", "title", "photo_url", "user_id", "user").Preload("User").First(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if Photo.User != nil {
		Photo.User.CreatedAt = nil
		Photo.User.UpdatedAt = nil
		Photo.User.Age = 0
		Photo.User.Password = ""
		Photo.User.ProfileImageURL = ""
	}

	c.JSON(http.StatusOK, Photo)
}

func PhotoUpdate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        Photo.ID,
		"caption":   Photo.Caption,
		"title":     Photo.Title,
		"photo_url": Photo.PhotoUrl,
		"user_id":   Photo.UserID,
	})
}

func PhotoDelete(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Select("Comment").Delete(Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully delete photo",
	})
}
