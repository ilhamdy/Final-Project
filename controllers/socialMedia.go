package controllers

import (
	"net/http"
	"strconv"
	"time"

	"mygram/helpers"
	"mygram/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SocialMediaRepo struct {
	DB *gorm.DB
}

// CreateSocialMedia godoc
// @Summary    	Create a social media
// @Description	Create a social media with authentication user
// @Tags        socialmedia
// @Accept      json
// @Produce     json
// @Param       json	body			utils.CreateSocialMedia true  "Create Social Media"
// @Success     201		{object}  utils.ResponseDataCreatedSocialMedia
// @Failure     400		{object}	utils.ResponseMessageSocialMedia
// @Failure     401		{object}	utils.ResponseMessageSocialMedia
// @Security    Bearer
// @Router      /socialmedia		[post]
func (m *SocialMediaRepo) CreateSocialMedia(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	socialMedia := models.SocialMedia{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	if contentType == "application/json" {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	socialMedia.UserID = uint(userId)
	socialMedia.CreatedAt = time.Now()
	socialMedia.UpdatedAt = time.Now()

	if err := m.DB.Debug().Create(&socialMedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to upload social media",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaURL,
		"user_id":          socialMedia.UserID,
		"created_at":       socialMedia.CreatedAt,
	})
}

// GetSocialMedia godoc
// @Summary    	Get all social media
// @Description	Get all social media with authentication user
// @Tags        socialmedia
// @Accept      json
// @Produce     json
// @Success     200	{object}	utils.ResponseDataGetSocialMedia
// @Failure     400	{object}	utils.ResponseMessageSocialMedia
// @Failure     401	{object}	utils.ResponseMessageSocialMedia
// @Security    Bearer
// @Router      /socialmedia	[get]
func (m *SocialMediaRepo) GetSocialMedia(c *gin.Context) {
	socialMedia := []models.SocialMedia{}

	if err := m.DB.Debug().Find(&socialMedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "can't find social media",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"social_media": socialMedia,
	})
}

// GetSocialMediaByIdPhoto godoc
// @Summary     Get a social media by Id
// @Description	Get a social media by id with authentication user
// @Tags        socialmedia
// @Accept      json
// @Produce     json
// @Param       id	path			string	true	"SocialMedia ID"
// @Success     200			{object}	utils.ResponseDataGetSocialMedia
// @Failure     400		{object}	utils.ResponseMessageSocialMedia
// @Failure     401		{object}	utils.ResponseMessageSocialMedia
// @Failure     404		{object}	utils.ResponseMessageSocialMedia
// @Security    Bearer
// @Router      /socialmedia/{id}	[get]
func (p *SocialMediaRepo) GetSocialMediaById(c *gin.Context) {
	GetId, _ := strconv.Atoi(c.Param("socialMediaId"))
	socialMedia := models.SocialMedia{}

	if err := p.DB.Debug().Where("id = ?", GetId).First(&socialMedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "social media not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"social_media": socialMedia,
	})
}

// UpdateSocialMedia godoc
// @Summary     Update a social media
// @Description	Update a social media by id with authentication user
// @Tags        socialmedia
// @Accept      json
// @Produce     json
// @Param       id		path      string	true	"SocialMedia ID"
// @Param				json	body			utils.UpdateSocialMedia	true	"Update Social Media"
// @Success     200		{object}	utils.ResponseDataUpdatedSocialMedia
// @Failure     400		{object}	utils.ResponseMessageSocialMedia
// @Failure     401		{object}	utils.ResponseMessageSocialMedia
// @Failure     404		{object}	utils.ResponseMessageSocialMedia
// @Security    Bearer
// @Router      /socialmedia/{id} [put]
func (m *SocialMediaRepo) UpdateSocialMedia(c *gin.Context) {
	contentType := helpers.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	getId, _ := strconv.Atoi(c.Param("socialMediaId"))

	socialMedia := models.SocialMedia{}
	oldSocialMedia := models.SocialMedia{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	socialMedia.UpdatedAt = time.Now()
	socialMedia.UserID = uint(userId)

	if err := m.DB.Debug().First(&oldSocialMedia, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "media not found",
			"message": err.Error(),
		})
		return
	}
	if err := m.DB.Debug().Model(&oldSocialMedia).Updates(&socialMedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to update media",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               oldSocialMedia.ID,
		"name":             oldSocialMedia.Name,
		"social_media_url": oldSocialMedia.SocialMediaURL,
		"user_id":          oldSocialMedia.UserID,
		"updated_at":       oldSocialMedia.UpdatedAt,
	})
}

// DeleteSocialMedia godoc
// @Summary     Delete a social media
// @Description	Delete a social media by id with authentication user
// @Tags        socialmedia
// @Accept      json
// @Produce     json
// @Param       id   path     string  true  "SocialMedia ID"
// @Success     200  {object}	utils.ResponseMessageDeletedSocialMedia
// @Failure     400  {object}	utils.ResponseMessageSocialMedia
// @Failure     401  {object}	utils.ResponseMessageSocialMedia
// @Failure     404  {object}	utils.ResponseMessageSocialMedia
// @Security    Bearer
// @Router      /socialmedia/{id} [delete]
func (m *SocialMediaRepo) DeleteSocialMedia(c *gin.Context) {
	getId, _ := strconv.Atoi(c.Param("socialMediaId"))
	socialMedia := models.SocialMedia{}

	if err := m.DB.Debug().First(&socialMedia, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "media not found",
			"message": err.Error(),
		})
		return
	}
	if err := m.DB.Debug().Delete(&socialMedia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to delete media",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "your social media has been successfully deleted",
	})
}
