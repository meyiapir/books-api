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
		book := api.Group("/book")
		{
			book.POST("/create", h.createBook)
			book.GET("/:id", h.getBook)
			book.PUT("/:id", h.updateBook)
			book.DELETE("/:id", h.deleteBook)
		}
	}

	return router
}
