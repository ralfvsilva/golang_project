package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type responseError struct {
	Message string
}

type handler struct {
}

type Item struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created-at"`
	UpdatedAt   time.Time `json:"updated-at"`
}

var items = []Item{
	{
		ID:          1,
		Code:        "Item001",
		Title:       "Camisa",
		Description: "camisa de algodão",
		Price:       79,
		Stock:       3,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

func main() {

	h := newHandler()
	i := newItem()

	router := gin.Default()

	router.GET("/", h.helloWorld)
	router.GET("/items", i.readItem)
	router.POST("/items", i.createItem)
	router.GET("/items/:id", i.readItemId)
	router.PUT("/items/:id", i.updateItem)
	// router.DELETE("/items/:id", )

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func newHandler() *handler {
	return &handler{}
}

func newItem() *Item {
	return &Item{}
}

func (u *Item) createItem(c *gin.Context) {
	var newItem Item

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	items = append(items, newItem)

	c.JSON(http.StatusCreated, items)
}

func (i *Item) readItemId(c *gin.Context) {

	id := c.Param("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range items {
		if convertId == v.ID {
			c.JSON(http.StatusOK, items)
			break
		} else {
			c.JSON(http.StatusNotFound, responseError{Message: "id not found"})
			break
		}
	}
}

func (i *Item) readItem(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

func (i *Item) updateItem(c *gin.Context) {
	var newItem Item

	if err := c.ShouldBindJSON(&newItem); err != nil {
		log.Fatal(err)
	}

	id := c.Param("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range items {
		if convertId == v.ID {
			if newItem.ID != 0 && newItem.ID > 0 {
				v.ID = newItem.ID
			}
			if newItem.Code != "" {
				v.Code = newItem.Code
			}
			c.JSON(http.StatusOK, items)
			return
		}
	}

	c.JSON(http.StatusNotFound, responseError{Message: "error 404"})

}

func (h handler) helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
