package main

import (
	"encoding/json"
	"fmt"
	"imdb-db/ent"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type name struct {
	nconst            string   `json:"nconst"`
	primaryName       string   `json:"primaryName"`
	birthYear         int      `json:"birthYear"`
	deathYear         int      `json:"deathYear"`
	primaryProfession []string `json:"primaryProfession"`
	knownForTitles    []string `json:"knownForTitles"`
	createdAt         string   `json:"createdAt""`
	updatedAt         string   `json:"updatedAt""`
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	envError := godotenv.Load(".env")

	if envError != nil {
		log.Fatalf("Error loading .env file")
	}

	user := os.Getenv("IMDB_USER")
	pass := os.Getenv("IMDB_PASS")

	connectionString := fmt.Sprintf("host=localhost port=5432 user=%s dbname=imdb password=%s", user, pass)

	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/ratings", func(c *gin.Context) {
		getRatings(client)
	})

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getRatings(client *ent.Client) []*ent.Ratings {
	// ctx := context.Context
	ratings, err := client.Ratings. // UserClient.
					Query(). // User query builder.
					All()
	ratingsJson, jsonError := json.Marshal(ratings)
	fmt.Println(string(ratingsJson))
	return ratings
}
