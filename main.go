package main

import (
  "github.com/gin-gonic/gin"
  "./models"
  "./controllers"

)

func main(){
  r := gin.Default()

  models.ConnectDatabase()

  r.GET("/api/books", controllers.FindBooks)
  r.POST("/api/books", controllers.CreateBook)
  r.GET("/api/books/:id", controllers.FindBook)
  r.PATCH("/api/books/:id", controllers.UpdateBook)




  r.Run()
}
