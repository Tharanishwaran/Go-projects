package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Quantity int    `json:"quantity"`
}

var movies = []movie{
	{ID: "1", Title: "Inception", Director: "Christopher Nolan", Quantity: 3},
	{ID: "2", Title: "The Godfather", Director: "Francis Ford Coppola", Quantity: 5},
	{ID: "3", Title: "The Dark Knight", Director: "Christopher Nolan", Quantity: 4},
}

func getMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}

func movieById(c *gin.Context) {
	id := c.Param("id")
	movie, err := getMovieById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, movie)
}

func checkoutMovie(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	movie, err := getMovieById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie not found."})
		return
	}

	if movie.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Movie not available."})
		return
	}

	movie.Quantity -= 1
	c.IndentedJSON(http.StatusOK, movie)
}

func returnMovie(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	movie, err := getMovieById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie not found."})
		return
	}

	movie.Quantity += 1
	c.IndentedJSON(http.StatusOK, movie)
}

func getMovieById(id string) (*movie, error) {
	for i, m := range movies {
		if m.ID == id {
			return &movies[i], nil
		}
	}

	return nil, errors.New("movie not found")
}

func createMovie(c *gin.Context) {
	var newMovie movie

	if err := c.BindJSON(&newMovie); err != nil {
		return
	}

	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusCreated, newMovie)
}

func main() {
	router := gin.Default()
	router.GET("/movies", getMovies)
	router.GET("/movies/:id", movieById)
	router.POST("/movies", createMovie)
	router.PATCH("/checkout", checkoutMovie)
	router.PATCH("/return", returnMovie)
	router.Run("localhost:8080")
}
