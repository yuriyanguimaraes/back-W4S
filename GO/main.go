package main

import (
	"github.com/gin-gonic/gin"
	"w4s/DB"
	"w4s/controllers"
)

func main() {
	// db:=""
	//creating connection with database
	r := gin.Default()         //starting the gin. //Iniciando o gin
	db := DB.SetupModels() //Connection database //Conexão banco de dados
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.GET("/user", controllers.FindUser)
	r.POST("/user", controllers.CreateUser)
	//r.POST("/login", controllers.Login) //still in progress //ainda a ser feito


	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080") // listando e escutando no localhost:8080
}
