package controllers

import (
	"net/http"
	"strconv"
	"time"

	"mygram/helpers"
	"mygram/models"

	// "mygram/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoRepo struct {
	DB *gorm.DB
}

// GetPhoto godoc
// @Summary    	Get all photos
// @Description	Get all photos with authentication user
// @Tags        photos
// @Accept      json
// @Produce     json
// @Success     200			{object}	utils.ResponseDataGetPhoto
// @Failure     400			{object}	utils.ResponseMessagePhoto
// @Failure     401			{object}	utils.ResponseMessagePhoto
// @Security    Bearer
// @Router      /photos	[get]
func (p *PhotoRepo) GetPhoto(c *gin.Context) {
	Photos := []models.Photo{}

	if err := p.DB.Debug().Preload("Comments").Preload("User").Find(&Photos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}

	/*
		getPhotos := []*utils.GetPhoto{}

		for _, photo := range Photos {
			getPhotos = append(getPhotos, &utils.GetPhoto{
				ID:        photo.ID,
				Title:     photo.Title,
				Caption:   photo.Caption,
				PhotoUrl:  photo.PhotoUrl,
				UserID:    photo.UserID,
				CreatedAt: photo.CreatedAt,
				UpdatedAt: photo.UpdatedAt,
				User: &utils.User{
					Email:    photo.User.Email,
					Username: photo.User.Username,
				},
			})
		}

		c.JSON(http.StatusOK, helpers.ResponseData{
			Status: "success",
			Data:   getPhotos,
		})
	*/

	c.JSON(http.StatusOK, gin.H{
		"photos": Photos,
	})
}

// GetPhotoByIdPhoto godoc
// @Summary     Get a photo by Id
// @Description	Get a photo by id with authentication user
// @Tags        photos
// @Accept      json
// @Produce     json
// @Param       id	path			string	true	"Photo ID"
// @Success     200			{object}	utils.ResponseDataGetPhoto
// @Failure     400		{object}	utils.ResponseMessagePhoto
// @Failure     401		{object}	utils.ResponseMessagePhoto
// @Failure     404		{object}	utils.ResponseMessagePhoto
// @Security    Bearer
// @Router      /photos/{id}	[get]
func (p *PhotoRepo) GetPhotoById(c *gin.Context) {
	GetId, _ := strconv.Atoi(c.Param("photoId"))
	Photo := models.Photo{}

	if err := p.DB.Debug().Preload("Comments").Preload("User").Where("id = ?", GetId).First(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"photo": Photo,
	})
}

// CreatePhoto godoc
// @Summary    	Create a photo
// @Description	Create a photo with authentication user
// @Tags        photos
// @Accept      json
// @Produce     json
// @Param       json		body			utils.CreatePhoto	true	"Create Photo"
// @Success     201			{object}  utils.ResponseDataCreatedPhoto
// @Failure     400			{object}	utils.ResponseMessagePhoto
// @Failure     401			{object}	utils.ResponseMessagePhoto
// @Security    Bearer
// @Router      /photos/	[post]
func (p *PhotoRepo) CreatePhoto(c *gin.Context) {
	Photo := models.Photo{}
	contextType := helpers.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	if contextType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = uint(userId)
	Photo.CreatedAt = time.Now()
	Photo.UpdatedAt = time.Now()

	if err := p.DB.Debug().Create(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to updload photo",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoURL,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
}

// UpdatePhoto godoc
// @Summary     Update a photo
// @Description	Update a photo by id with authentication user
// @Tags        photos
// @Accept      json
// @Produce     json
// @Param       id		path      string	true	"Photo ID"
// @Param       json	body			utils.UpdatePhoto true  "Photo"
// @Success     200		{object}  utils.ResponseDataUpdatedPhoto
// @Failure     400		{object}	utils.ResponseMessagePhoto
// @Failure     401		{object}	utils.ResponseMessagePhoto
// @Failure     404		{object}	utils.ResponseMessagePhoto
// @Security    Bearer
// @Router      /photos/{id}		[put]
func (p *PhotoRepo) UpdatePhoto(c *gin.Context) {
	GetId, _ := strconv.Atoi(c.Param("photoId"))

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"]

	contextType := helpers.GetContentType(c)
	Photo := models.Photo{}
	oldPhoto := models.Photo{}

	if contextType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UpdatedAt = time.Now()
	Photo.UserID = uint(userId.(float64))

	if err := p.DB.Debug().First(&oldPhoto, GetId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}

	if err := p.DB.Debug().Model(&oldPhoto).Updates(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to update photo",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         oldPhoto.ID,
		"title":      oldPhoto.Title,
		"caption":    oldPhoto.Caption,
		"photo_url":  oldPhoto.PhotoURL,
		"user_id":    oldPhoto.UserID,
		"updated_at": oldPhoto.UpdatedAt,
	})
}

// DeletePhoto godoc
// @Summary     Delete a photo
// @Description	Delete a photo by id with authentication user
// @Tags        photos
// @Accept      json
// @Produce     json
// @Param       id	path			string	true	"Photo ID"
// @Success     200	{object}	utils.ResponseMessageDeletedPhoto
// @Failure     400	{object}	utils.ResponseMessagePhoto
// @Failure     401	{object}	utils.ResponseMessagePhoto
// @Failure     404	{object}	utils.ResponseMessagePhoto
// @Security    Bearer
// @Router      /photos/{id}	[delete]
func (p *PhotoRepo) DeletePhoto(c *gin.Context) {
	GetId, _ := strconv.Atoi(c.Param("photoId"))
	Photo := models.Photo{}

	if err := p.DB.Debug().First(&Photo, GetId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}

	if err := p.DB.Debug().Delete(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to delete photo",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "your photo has been succesfully deleted",
	})
}
