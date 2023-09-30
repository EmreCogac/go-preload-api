package services

import (
	"voxellia-app/voxellia-app/initializers"
	"voxellia-app/voxellia-app/models"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {

	db := initializers.DB
	levels := []models.Level{}
	section := []models.Section{}
	db.Find(&section)
	db.Preload("Section").Find(&levels)

	c.JSON(200, gin.H{
		"Sections": section,
		"Levels":   levels,
	})
}
