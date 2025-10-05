package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
} // length = 3, capacity = 3

/*
1- Initialize a Gin router with Default
2- Use the GET function to associate the GET HTTP method and /albums path with getAlbums function handler
3- Use the POST function to associate the POST HTTP method and /albums path with the addAlbum function handler
4- Use the Run function to attach the router to an http.Server and start the server
*/
func main() {
	var router *gin.Engine = gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", addAlbum)
	router.Run("localhost:8080")
}

/*
This function creates a JSON from the slice of albums and returns the JSON as the response.
The function takes a variable c or type *gin.Context
The function uses the gin context to call the IndentedJson  function,
which serializes the slice into JSON and add it to the response
*/
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func addAlbum(context *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the request's payload to newAlbum
	var err error = context.BindJSON(&newAlbum)
	if err != nil {
		return
	}

	// Add the new acceptable album to the in-memory slice
	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}
