package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/DeliusPit/golang-gorm-postgres/controllers"
	"github.com/DeliusPit/golang-gorm-postgres/middleware"
)

type BucketRouteController struct {
	bucketController controllers.BucketController
}

func NewRouteBucketController(bucketController controllers.BucketController) BucketRouteController {
	return BucketRouteController{bucketController}
}

func (pc *BucketRouteController) BucketRoute(rg *gin.RouterGroup) {

	router := rg.Group("buckets")
	router.Use(middleware.DeserializeUser())
	router.POST("/", pc.bucketController.CreateBucket)
	router.GET("/", pc.bucketController.FindBuckets)
	router.PUT("/:bucketId", pc.bucketController.UpdateBucket)
	router.GET("/:bucketId", pc.bucketController.FindBucketById)
	router.DELETE("/:bucketId", pc.bucketController.DeleteBucket)
}
