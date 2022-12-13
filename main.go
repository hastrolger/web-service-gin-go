package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

// home responds with a simple message 
func home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Home page")
}

// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// newAlbum creates a new album 
func postAlbum(c *gin.Context) {
	var newAlbum album

	// call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}	

	// add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
}

func main() {
	// instantiate a router of gin
	router := gin.Default()
	
	// routes 
	router.GET("/", home)
	router.GET("/albums", getAlbums)
	router.POST("/albums/new", postAlbum)

	// execute the server
	router.Run("localhost:8001")

}