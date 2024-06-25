package handler

import (
	"api/internal/domain"
	"api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type responseError struct {
	Message string
}

type handler struct {
	itemService service.ItemService
}

func NewHandler(itemService service.ItemService) *handler {
	return &handler{
		itemService: itemService,
	}
}

func (h *handler) CreateItem(c *gin.Context) {
	var newItem domain.Item

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	items := h.itemService.AddItem(newItem)

	c.JSON(http.StatusCreated, items)
}

func (h *handler) ReadItemId(c *gin.Context) {
	id := c.Param("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError{Message: "id not found"})
		return
	}

	item := h.itemService.ReadItem(convertId)
	if item != nil {
		c.JSON(http.StatusOK, item)
		return
	}

	c.JSON(http.StatusNotFound, responseError{Message: "id not found"})
}

func (h *handler) ReadItem(c *gin.Context) {
	c.JSON(http.StatusOK, h.itemService.GetAllItems())
}

func (h *handler) UpdateItem(c *gin.Context) {
	var existItem domain.Item

	if err := c.ShouldBindJSON(&existItem); err != nil {
		c.JSON(http.StatusBadRequest, responseError{Message: "decoding error"})
	}

	id := c.Param("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError{Message: "id not found"})
	}

	result := h.itemService.UpdateItem(convertId, existItem)

	if result != nil {
		c.JSON(http.StatusOK, result)
		return
	}

	c.JSON(http.StatusNotFound, responseError{Message: "error 404"})
}

func (h *handler) DeleteItem(c *gin.Context) {
	id := c.Param("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	result := h.itemService.DeleteItem(convertId)
	if result != nil {
		c.JSON(http.StatusBadRequest, responseError{Message: "error delete item"})
	}

	c.JSON(http.StatusOK, "Item deleted.")
}

func (h handler) HelloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
