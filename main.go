package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"TodoApp/api/controllers"

	"github.com/kamva/mgm/v3"
	"github.com/gin-contrib/cors"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct{
	Name string `json:"name"`
	Body string `json:"body"`
}

func main() {	
	godotenv.Load()
	app := gin.Default()
	err := mgm.SetDefaultConfig(nil, os.Getenv("DB_NAME"), options.Client().ApplyURI(os.Getenv("DB_STRING")))
	mountController := controllers.MountController{}

	if err != nil{
		log.Fatal(err)
	}

	app.Use(cors.Default())
	mountController.Init(app.Group("/api/v1"))
	
	fmt.Println("Connected to DB!")


	// app.GET("views/:id/end", func(c *gin.Context) {
	// 	fmt.Println("id:", c.Param("id"), "and context value", c.Value("user"))
	// 	c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	// })

	// app.GET("views/:id/start", func(c *gin.Context) {
	// 	fmt.Println("id:", c.Param("id"))
	// 	c.JSON(http.StatusOK, gin.H{"message": "hello world from: " + c.Request.URL.String()})
	// })

	// app.NoRoute(func(c *gin.Context) {
	// 	c.JSON(404, gin.H{"error": "route " + c.FullPath() + " does not exist"})
	// })

	app.Run()
}

func myMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("user", gin.H{"hello": "context"})
		c.Next()
	}
}
