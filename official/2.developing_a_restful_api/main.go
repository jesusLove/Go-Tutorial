package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
// 1. 创建数据
// 专辑结构体
type album struct {
	ID string 	`json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
// 2. 编写返回全部专辑的处理方法
// 获取全部专辑 GET  请求
func getAlbums(c *gin.Context) {
	//c.IndentedJSON(http.StatusOK, albums)
	c.JSON(http.StatusOK, albums)
}

// 3. 编写新增专辑的处理方法
func postAlbums(c *gin.Context) {
	var newAlbum album
	// 将 请求体绑定到 newAlbum 上
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "album error"})
		return
	}
	albums = append(albums, newAlbum)
	c.JSON(http.StatusOK, newAlbum)
}

// 4. 返回指定的 Item
func getAlbumById(c *gin.Context) {
	// 获取 id
	id := c.Param("id")
	// 遍历查找
	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message":"album not found"})
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.Run("localhost:8080")
}