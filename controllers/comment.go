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

type CommentRepo struct {
	DB *gorm.DB
}

// GetComment godoc
// @Summary			Get all comments
// @Description	Get all comments with authentication user
// @Tags        comments
// @Accept      json
// @Produce     json
// @Success     200	{object}	utils.ResponseDataGetComment
// @Failure     400	{object}	utils.ResponseMessageComment
// @Failure     401	{object}	utils.ResponseMessageComment
// @Security    Bearer
// @Router      /comments     [get]
func (o *CommentRepo) GetComment(c *gin.Context) {
	Comments := []models.Comment{}

	if err := o.DB.Debug().Preload("User").Preload("Photo").Find(&Comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "comment not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Comments,
	})
}

// GetCommentByIdPhoto godoc
// @Summary     Get a comment by Id
// @Description	Get a comment by id with authentication user
// @Tags        comments
// @Accept      json
// @Produce     json
// @Param       id	path			string	true	"Photo ID"
// @Success     200			{object}	utils.ResponseDataGetComment
// @Failure     400		{object}	utils.ResponseMessageComment
// @Failure     401		{object}	utils.ResponseMessageComment
// @Failure     404		{object}	utils.ResponseMessageComment
// @Security    Bearer
// @Router      /comments/{id}	[get]
func (p *CommentRepo) GetCommentById(c *gin.Context) {
	GetId, _ := strconv.Atoi(c.Param("commentId"))
	Comment := models.Comment{}

	if err := p.DB.Debug().Preload("User").Preload("Photo").Where("id = ?", GetId).First(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "comment not found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Comment,
	})
}

// CreateComment godoc
// @Summary			create a comment
// @Description	create a comment with authentication user
// @Tags        comments
// @Accept      json
// @Produce     json
// @Param       json	body			utils.CreateComment true  "Create Comment"
// @Success     201		{object}  utils.ResponseDataCreatedComment
// @Failure     400		{object}	utils.ResponseMessageComment
// @Failure     401		{object}	utils.ResponseMessageComment
// @Security    Bearer
// @Router      /comments/	[post]
func (o *CommentRepo) CreateComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)
	Comment := models.Comment{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	if err := o.DB.Debug().Find(&models.Comment{}, Comment.PhotoID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}
	Comment.UserID = uint(userId)
	Comment.CreatedAt = time.Now()
	Comment.UpdatedAt = time.Now()

	if err := o.DB.Debug().Create(&Comment).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error":   "failed to upload comment",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}

// UpdateComment godoc
// @Summary			Update a comment
// @Description	Update a comment by id with authentication user
// @Tags        comments
// @Accept      json
// @Produce     json
// @Param       id		path			string  true  "Comment ID"
// @Param       json	body			utils.UpdateComment	true	"Update Comment"
// @Success     200		{object}  utils.ResponseDataUpdatedComment
// @Failure     400		{object}	utils.ResponseMessageComment
// @Failure     401		{object}	utils.ResponseMessageComment
// @Failure     404		{object}	utils.ResponseMessageComment
// @Security    Bearer
// @Router      /comments/{id}	[put]
func (o *CommentRepo) UpdateComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	oldComment := models.Comment{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	getId, _ := strconv.Atoi(c.Param("commentId"))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = uint(userId)
	Comment.UpdatedAt = time.Now()

	if err := o.DB.Debug().First(&oldComment, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Comment not found",
			"messsage": err.Error(),
		})
		return
	}

	if err := o.DB.Debug().Model(&oldComment).Updates(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to update comment",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         oldComment.ID,
		"message":    oldComment.Message,
		"user_id":    oldComment.UserID,
		"updated_at": oldComment.UpdatedAt,
	})
}

// DeleteComment godoc
// @Summary			Delete a comment
// @Description	Delete a comment by id with authentication user
// @Tags        comments
// @Accept      json
// @Produce     json
// @Param       id  path			string	true	"Comment ID"
// @Success     200 {object}	utils.ResponseMessageDeletedComment
// @Failure     400 {object}	utils.ResponseMessageComment
// @Failure     401	{object}	utils.ResponseMessageComment
// @Failure     404	{object}	utils.ResponseMessageComment
// @Security    Bearer
// @Router      /comments/{id}	[delete]
func (o *CommentRepo) DeleteComment(c *gin.Context) {
	getId, _ := strconv.Atoi(c.Param("commentId"))

	Comment := models.Comment{}

	if err := o.DB.Debug().First(&Comment, getId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "comment not found",
			"message": err.Error(),
		})
		return
	}

	if err := o.DB.Debug().Delete(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to delete comment",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your comment has been successfully deleted",
	})
}
