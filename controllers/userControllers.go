package controllers

import (
	"challenge294/database"
	"challenge294/helpers"
	"challenge294/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_,_ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": User.ID,
		"username": User.Username,
		"role": User.Role,
	})
}

func UserLogin(c *gin.Context)  {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_,_ = db,contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("username = ?", User.Username).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "invaild username/password",
		})
		return
	}
	
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "invaild username/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Username, User.Role)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}