package controllers

import (
	bookModel "LearnAPI/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var counter int
var books []bookModel.BookDetails

func MapBookData(book bookModel.BookDetails) bookModel.BookBody {
	return bookModel.BookBody{book.Id, book.Name, book.Publisher}
}

func AddBook(c *gin.Context) {
	var newBook bookModel.BookDetails
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Internal server error"})
		return
	}
	if newBook.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Gagal menambahkan buku. Mohon isi nama buku"})
		return
	}
	if newBook.ReadPage > newBook.PageCount {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Gagal menambahkan buku. readPage tidak boleh lebih besar dari pageCount"})
		return
	}
	newBook.Finished = false
	if newBook.ReadPage == newBook.PageCount {
		newBook.Finished = true
	}
	currentTimeStr := time.Now().String()
	newBook.InsertedAt = currentTimeStr
	newBook.UpdatedAt = currentTimeStr
	newBook.Id = counter
	counter++
	books = append(books, newBook)
	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Buku berhasil ditambahkan", "data": gin.H{"bookId": newBook.Id}})
}

func GetBooks(c *gin.Context) {
	name := c.Query("name")
	finished := c.Query("finished")
	reading := c.Query("reading")
	b2s := map[bool]string{false: "0", true: "1"}

	var booksBody []bookModel.BookBody = []bookModel.BookBody{}
	for _, book := range books {
		if (name == "" || strings.Contains(strings.ToLower(book.Name), strings.ToLower(name))) && (finished == "" || finished == b2s[book.Finished]) && (reading == "" || reading == b2s[book.Reading]) {
			booksBody = append(booksBody, MapBookData(book))
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"books": booksBody}})
}

func GetBookById(c *gin.Context) {
	for _, book := range books {
		if id, err := strconv.Atoi(c.Param("id")); id == book.Id {
			if err == nil {
				c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"book": book}})
				return
			} else {
				break
			}
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Buku tidak ditemukan"})
}

func UpdateBookById(c *gin.Context) {
	var newBook bookModel.BookDetails
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Internal server error"})
		return
	}
	if newBook.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Gagal memperbarui buku. Mohon isi nama buku"})
		return
	}
	if newBook.ReadPage > newBook.PageCount {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Gagal memperbarui buku. readPage tidak boleh lebih besar dari pageCount"})
		return
	}
	newBook.Finished = false
	if newBook.ReadPage == newBook.PageCount {
		newBook.Finished = true
	}
	currentTimeStr := time.Now().String()
	newBook.InsertedAt = currentTimeStr
	newBook.UpdatedAt = currentTimeStr

	for index, book := range books {
		if id, err := strconv.Atoi(c.Param("id")); id == book.Id {
			if err == nil {
				books[index] = newBook
				c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Buku berhasil diperbarui"})
				return
			} else {

				break
			}
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Gagal memperbarui buku. Id tidak ditemukan"})
}

func DeleteBookById(c *gin.Context) {
	for index, book := range books {
		if id, err := strconv.Atoi(c.Param("id")); id == book.Id {
			if err == nil {
				books = append(books[:index], books[index+1:]...)
				c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Buku berhasil dihapus"})
				return
			} else {
				break
			}
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Buku gagal dihapus. Id tidak ditemukan"})
}
