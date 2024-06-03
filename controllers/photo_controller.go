package controllers

import (
    "net/http"
    "BTPN/database"
    "BTPN/models"
    "github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
    var photo models.Photo
    if err := c.ShouldBindJSON(&photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, _ := c.Get("userID")
    photo.UserID = userID.(uint)

    result := database.DB.Create(&photo)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": photo})
}

func GetPhotos(c *gin.Context) {
    var photos []models.Photo
    database.DB.Preload("User").Find(&photos)
    c.JSON(http.StatusOK, gin.H{"data": photos})
}

func UpdatePhoto(c *gin.Context) {
    var photo models.Photo
    if err := database.DB.Where("id = ?", c.Param("photoId")).First(&photo).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
        return
    }

    userID, _ := c.Get("userID")
    if photo.UserID != userID {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    if err := c.ShouldBindJSON(&photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Save(&photo)
    c.JSON(http.StatusOK, gin.H{"data": photo})
}

func DeletePhoto(c *gin.Context) {
    var photo models.Photo
    if err := database.DB.Where("id = ?", c.Param("photoId")).First(&photo).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
        return
    }

    userID, _ := c.Get("userID")
    if photo.UserID != userID {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    database.DB.Delete(&photo)
    c.JSON(http.StatusOK, gin.H{"data": true})
}
