package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/models"
)

func FindUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user []models.User
	db.Find(&user)
	c.JSON(http.StatusOK, gin.H{
		"user":user,
	})

}

//POST /user
//Create a new user

type CreateUserInput struct{
	Nickname string  `json:"nickname" binding:required`
	Email 	 string  `json:"email" binding:required `
	Password string  `json:"password" binding:required`
	Name     string  `json:"name" binding:required`
	Lastname string  `json:"string" binding:required`
}

func CreateUser( c *gin.Context){

	db:= c.MustGet("db").(*gorm.DB)
	//Validating input
	var input CreateUserInput
	if err:= c.ShouldBindJSON(&input);err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	//Creating user
	user:=models.User{
		Nickname: input.Nickname,
		Email:input.Email,
		Password :input.Password,
		Name: input.Name,
		Lastname: input.Lastname,
	}
	db.Create(&user)
	c.JSON(http.StatusOK,gin.H{"data":user})
	
}
