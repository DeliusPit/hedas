package main

import (
	"fmt"
	"log"

	"github.com/DeliusPit/golang-hedas/initializers"
	"github.com/DeliusPit/golang-hedas/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Bucket{}, &models.File{})
	fmt.Println("👍 Migration complete")
}
