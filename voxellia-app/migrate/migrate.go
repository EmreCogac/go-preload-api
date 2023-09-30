package main

import (
	"fmt"
	"log"
	"voxellia-app/voxellia-app/initializers"
	"voxellia-app/voxellia-app/models"
)

func init() {
	config, err := initializers.LoadConfig("..")
	if err != nil {
		log.Fatal("? could not load ", err)

	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.Section{})
	initializers.DB.AutoMigrate(&models.Level{})
	fmt.Println("? migration complete")
}
