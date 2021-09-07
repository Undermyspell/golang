package albumspkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "Alber Vaillant", Artist: "The Streetboy", Price: 39.99},
	{ID: "5", Title: "Terry Hills", Artist: "US5", Price: 39.99},
	{ID: "6", Title: "Christopher Wonder", Artist: "Westlife", Price: 39.99},
	{ID: "7", Title: "Chuck Bass", Artist: "Shaggy", Price: 39.99},
	{ID: "8", Title: "Dexter Morgan", Artist: "Tobi van Buren", Price: 39.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func SetAlbumTitle(c *gin.Context) {
	var req AlbumRequest

	err := json.NewDecoder(c.Request.Body).Decode(&req)

	if err != nil {
		println("bad request for set album title")
		c.Status(http.StatusBadRequest)
		return
	}

	albumId := c.Param("id")

	album := findById(albums, albumId)

	println(album)

	album.setTitle(req.Title)

	c.IndentedJSON(http.StatusOK, albums)
}

func StringifyAlbums(c *gin.Context) {
	strings := []string{}

	messages := make(chan string)

	for _, album := range albums {
		go stringify(album, messages)
	}

	for i := 0; i < len(albums); i++ {
		strings = append(strings, <-messages)
	}

	c.IndentedJSON(http.StatusOK, strings)
}

func findById(a []Album, id string) *Album {
	for i, album := range a {
		if album.ID == id {
			return &albums[i]
		}
	}
	return nil
}

func stringify(a Album, c chan string) {
	elements := []string{a.ID, a.Title, a.Artist, fmt.Sprintf("%f", a.Price)}
	s := strings.Join(elements, "|")

	c <- s
}
