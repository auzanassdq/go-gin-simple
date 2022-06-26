package models

import "github.com/gin-gonic/gin"

type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

func SetDataMovie() []Movie {
	movie1 := Movie{1, "Spiderman", 2021}
	movie2 := Movie{2, "Venom", 2022}
	movie3 := Movie{3, "Morbius", 2022}

	var listMovie []Movie

	listMovie = append(listMovie, movie1)
	listMovie = append(listMovie, movie2)
	listMovie = append(listMovie, movie3)

	return listMovie
}

func AllMovie(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success get data",
		"data":    SetDataMovie(),
	})
}
