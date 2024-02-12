package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// penulisan router untuk api
	router := gin.Default()

	router.GET("/", rootHandle)
	router.GET("/hello", HelloHandler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query/", QueryHandler)

	// route defaultnya adalah localhost:8080/
	router.Run()
}

func rootHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":   "Muhammad Rizki Syahputra",
		"bio":    "Sofware Engineer",
		"status": "mahasiswa",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "Hello world",
		"subtitle": "learn go api",
	})
}

func booksHandler(c *gin.Context) {

	// menangkap id dan membuat id di url api, dengan fungsi Param
	id := c.Param("id")

	// menangkap title dan membuat title di url api, dengan fungsi Param
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"title":   title,
		"Pembuat": "James Bond",
	})
}

func QueryHandler(c *gin.Context) {

	// menangkap title dan membuat title dari url api, dengan fungsi Query
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"Judul": "Bicara itu harus pake mulut",
	})
}
