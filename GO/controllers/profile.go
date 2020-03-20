package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/models"
)

func FindUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var profile []models.profile
	db.Find(&profile)
	c.JSON(http.StatusOK, gin.H{
		"profile":profile,
	})

}

// POST Profile
// Criando um novo Profile

type CreateProfileInput struct{
	Avatar 	 string  `json:"avatar"`
	Telefone int8  	 `json:"telefone"`
	Data 	 int8    `json:"data_nascimento"`
}

func CreateProfile( c *gin.Context){

	db:= c.MustGet("db").(*gorm.DB)
	//Validating input
	var input CreateProfileInput
	if err:= c.ShouldBindJSON(&input);err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	//Creating Profile
	profile:=models.Profile{
		Avatar:  	input.Avatar,
		Telefone:	input.Telefone,
		Data: 	 	input.Data
	}
	db.Create(&profile)
	c.JSON(http.StatusOK,gin.H{"data":profile})
	
}
