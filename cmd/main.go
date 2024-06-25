package main

import (
	"api/internal/handler"
	"api/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	h := handler.NewHandler(service.ItemService{})

	router := gin.Default()

	router.GET("/", h.HelloWorld)
	router.GET("/items", h.ReadItem)
	router.POST("/items", h.CreateItem)
	router.GET("/items/:id", h.ReadItemId)
	router.PUT("/items/:id", h.UpdateItem)
	// router.DELETE("/items/:id", )

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
