package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		books := api.Group("/books")
		{
			books.GET("/")
		}
		book := api.Group("/book/:id")
		{
			book.POST("/")
			book.GET("/:book_id")
			book.PUT("/:book_id")
			book.DELETE("/:book_id")
		}

	}

	return router
}
