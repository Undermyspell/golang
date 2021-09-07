package albumspkg

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
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

func findById(a []Album, id string) *Album {
	for i, album := range a {
		if album.ID == id {
			return &albums[i]
		}
	}
	return nil
}
