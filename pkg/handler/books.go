package handler

import (
	"booksApi"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getBooks(c *gin.Context) {
	books, err := h.services.Books.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) getBook(c *gin.Context) {
	// Вывод всех переданных параметров в запросе в print
	fmt.Println(c.Request.URL.Query())
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid book id param")
		return
	}

	book, err := h.services.Book.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, book)

}

func (h *Handler) createBook(c *gin.Context) {

	// Создание переменной типа Book
	var book booksApi.Book

	// Привязка данных из запроса к переменной book
	if err := c.BindJSON(&book); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Создание книги
	id, err := h.services.Book.Create(book)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Вывод id созданной книги
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid book id param")
		return
	}

	var input booksApi.UpdateBookInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Book.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})

}

func (h *Handler) deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid book id param")
		return
	}

	err = h.services.Book.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
