package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/DeliusPit/golang-hedas/models"
	"gorm.io/gorm"
)

type FileController struct {
	DB *gorm.DB
}

func NewFileController(DB *gorm.DB) FileController {
	return FileController{DB}
}

func (pc *FileController) CreateFile(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreateFileRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newFile := models.File{
		Title:          payload.Title,
		Bucket:         payload.Bucket,
		User:           currentUser.ID,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	result := pc.DB.Create(&newFile)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "File with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newFile})
}

func (pc *FileController) UpdateFile(ctx *gin.Context) {
	fileId := ctx.Param("fileId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdateFile
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedFile models.File
	result := pc.DB.First(&updatedFile, "id = ?", fileId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	now := time.Now()
	postToUpdate := models.File{
		Title:          payload.Title,
		Bucket:         payload.Bucket,
		User:           currentUser.ID,
		CreatedAt:      updatedFile.CreatedAt,
		UpdatedAt:      now,
	}

	pc.DB.Model(&updatedFile).Updates(postToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedFile})
}

func (pc *FileController) FindFileById(ctx *gin.Context) {
	fileId := ctx.Param("fileId")

	var post models.File
	result := pc.DB.First(&post, "id = ?", fileId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": post})
}

func (pc *FileController) FindFiles(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var posts []models.File
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&posts)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(posts), "data": posts})
}

func (pc *FileController) DeleteFile(ctx *gin.Context) {
	fileId := ctx.Param("fileId")

	result := pc.DB.Delete(&models.File{}, "id = ?", fileId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
