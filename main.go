package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/DeliusPit/golang-hedas/controllers"
	"github.com/DeliusPit/golang-hedas/initializers"
	"github.com/DeliusPit/golang-hedas/routes"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	BucketController      controllers.BucketController
	BucketRouteController routes.BucketRouteController

	FileController      controllers.FileController
	FileRouteController routes.FileRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	BucketController = controllers.NewBucketController(initializers.DB)
	BucketRouteController = routes.NewRouteBucketController(BucketController)

	FileController = controllers.NewFileController(initializers.DB)
	FileRouteController = routes.NewRouteFileController(FileController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
//	corsConfig.AllowOrigins = []string{"http://192.168.178.31:8800", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Hedas - Headless Document Archiving System"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	BucketRouteController.BucketRoute(router)
	FileRouteController.FileRoute(router)	
	log.Fatal(server.Run(":" + config.ServerPort))
}
