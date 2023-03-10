package handler

import (
	"booksApi/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		books := api.Group("/books")
		{
			books.GET("/", h.getBooks)
		}
		book := api.Group("/book/:id")
		{
			book.POST("/", h.createBook)
			book.GET("/:book_id", h.getBook)
			book.PUT("/:book_id", h.updateBook)
			book.DELETE("/:book_id", h.deleteBook)
		}
	}

	return router
}
