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

type BucketController struct {
	DB *gorm.DB
}

func NewBucketController(DB *gorm.DB) BucketController {
	return BucketController{DB}
}

func (pc *BucketController) CreateBucket(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreateBucketRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newBucket := models.Bucket{
		Title:          payload.Title,
		Versioning:     payload.Versioning,
		Locking:        payload.Locking,
		Quota:          payload.Quota,
		User:           currentUser.ID,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	result := pc.DB.Create(&newBucket)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Bucket with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newBucket})
}

func (pc *BucketController) UpdateBucket(ctx *gin.Context) {
	postId := ctx.Param("postId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdateBucket
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedBucket models.Bucket
	result := pc.DB.First(&updatedBucket, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	now := time.Now()
	postToUpdate := models.Bucket{
		Title:          payload.Title,
		Versioning:     payload.Versioning,
		Locking:        payload.Locking,
		Quota:          payload.Quota,
		User:           currentUser.ID,
		CreatedAt:      updatedBucket.CreatedAt,
		UpdatedAt:      now,
	}

	pc.DB.Model(&updatedBucket).Updates(postToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedBucket})
}

func (pc *BucketController) FindBucketById(ctx *gin.Context) {
	postId := ctx.Param("postId")

	var post models.Bucket
	result := pc.DB.First(&post, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": post})
}

func (pc *BucketController) FindBuckets(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var posts []models.Bucket
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&posts)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(posts), "data": posts})
}

func (pc *BucketController) DeleteBucket(ctx *gin.Context) {
	postId := ctx.Param("postId")

	result := pc.DB.Delete(&models.Bucket{}, "id = ?", postId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
