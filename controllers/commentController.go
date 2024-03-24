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

func CommentCreate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	err := db.Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       Comment.ID,
		"message":  Comment.Message,
		"photo_id": Comment.PhotoID,
		"user_id":  Comment.UserID,
	})
}

func CommentGetAll(c *gin.Context) {
	db := database.GetDB()

	var Comment []models.Comment

	err := db.Select("id", "message", "photo_id", "user_id").Preload("User").Preload("Photo").Find(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	for _, comm := range Comment {
		if comm.User != nil {
			comm.User.CreatedAt = nil
			comm.User.UpdatedAt = nil
			comm.User.Age = 0
			comm.User.Password = ""
			comm.User.ProfileImageURL = ""
		}
		if comm.Photo != nil {
			comm.Photo.CreatedAt = nil
			comm.Photo.UpdatedAt = nil
		}
	}

	c.JSON(http.StatusOK, Comment)
}

func CommentGetById(c *gin.Context) {
	db := database.GetDB()
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	Comment := models.Comment{}
	Comment.ID = uint(commentId)

	err := db.Select("id", "message", "photo_id", "user_id").Preload("User").Preload("Photo").First(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if Comment.User != nil {
		Comment.User.CreatedAt = nil
		Comment.User.UpdatedAt = nil
		Comment.User.Age = 0
		Comment.User.Password = ""
		Comment.User.ProfileImageURL = ""
	}
	if Comment.Photo != nil {
		Comment.Photo.CreatedAt = nil
		Comment.Photo.UpdatedAt = nil
	}

	c.JSON(http.StatusOK, Comment)
}

func CommentUpdate(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.ID = uint(commentId)

	err := db.Select("photo_id", "user_id").First(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err = db.Model(&Comment).Updates(models.Comment{Message: Comment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       Comment.ID,
		"message":  Comment.Message,
		"photo_id": Comment.PhotoID,
		"user_id":  Comment.UserID,
	})
}

func CommentDelete(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.ID = uint(commentId)

	err := db.Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully delete comment",
	})
}
