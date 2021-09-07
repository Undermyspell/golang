package main

import (
	albumspkg "restapi/albums"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", albumspkg.GetAlbums)
	router.GET("/albums/stringify", albumspkg.StringifyAlbums)
	router.PUT("/album/:id/settitle", albumspkg.SetAlbumTitle)

	router.Run("localhost:4711")
}
