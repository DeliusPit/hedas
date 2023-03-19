package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/DeliusPit/golang-gorm-postgres/controllers"
	"github.com/DeliusPit/golang-gorm-postgres/middleware"
)

type FileRouteController struct {
	fileController controllers.FileController
}

func NewRouteFileController(fileController controllers.FileController) FileRouteController {
	return FileRouteController{fileController}
}

func (pc *FileRouteController) FileRoute(rg *gin.RouterGroup) {

	router := rg.Group("files")
	router.Use(middleware.DeserializeUser())
	router.POST("/", pc.fileController.CreateFile)
	router.GET("/", pc.fileController.FindFiles)
	router.PUT("/:fileId", pc.fileController.UpdateFile)
	router.GET("/:fileId", pc.fileController.FindFileById)
	router.DELETE("/:fileId", pc.fileController.DeleteFile)
}
