package controllers

import (
	"fmt"
	"net/http"
	"time"

	"mygram/helpers"
	"mygram/models"

	// "mygram/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

// UserRegister godoc
// @Summary Register a new user
// @Description	register a new user and return user data
// @Tags users
// @Accept json
// @Produce json
// @Param json body utils.RegisterUser true "Register User"
// @Success	201	{object} utils.ResponseDataRegisteredUser
// @Failure	400 {object} utils.ResponseMessageUser
// @Failure	409 {object} utils.ResponseMessageUser
// @Router /users/register [post]
func (u *UserRepo) UserRegister(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	_, _ = u.DB, contentType

	User := models.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.CreatedAt = time.Now()
	User.UpdatedAt = time.Now()

	if err := u.DB.Debug().Create(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to create user data",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":       User.GormModel.ID,
		"email":    User.Email,
		"username": User.Username,
		"age":      User.Age,
	})
	/*
		c.JSON(http.StatusCreated, helpers.ResponseData{
			Status: "success",
			Data: utils.RegisteredUser{
				Age:      user.Age,
				Email:    user.Email,
				ID:       user.ID,
				Username: user.Username,
			},
		})
	*/
}

// UserLogin godoc
// @Summary Login a user
// @Description	Authentication a user and retrieve a token
// @Tags users
// @Accept json
// @Produce json
// @Param json body utils.LoginUser true "Login User"
// @Success	201	{object} utils.ResponseDataLoggedinUser
// @Failure	400 {object} utils.ResponseMessageUser
// @Failure	409 {object} utils.ResponseMessageUser
// @Router /users/login [post]
func (u *UserRepo) UserLogin(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	_, _ = u.DB, contentType

	User := models.User{}

	if contentType == "application/json" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	password := User.Password

	if err := u.DB.Debug().Where("email=?", User.Email).Take(&User).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "Invalid email/password",
		})
		return
	}
	fmt.Println((User.Password), (password))
	if comparePass := helpers.ComparePass([]byte(User.Password), []byte(password)); !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "Invalid email/password",
		})
		return
	}
	token := helpers.GenerateToken(uint(User.ID), User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
