# 使用 Go 和 Gin 开发 RESTful API 

介绍使用 Go 和 Gin Web 框架编写 RESTful Web 服务 API 的基础知识。

> [Gin Web Framework](https://gin-gonic.com/docs/)

## Gin 简介

> Gin 是一个用 Go 编写 HTTP Web 的框架。

特点：
* 快速：基于 Radix tree 的路由，内存占用小，没有反射，可预测的 API 性能。
* 中间件：传入的 HTTP 请求经过一系列中间件处理生成最终结果。
* Crash-free: 捕获恐慌并恢复，保证服务器始终可用。
* JSON 验证：解析和验证请求的 JSON。
* 路由分组：更好的组织路由，接口授权、API 版本。不同的组可以嵌套使用。
* 错误管理：提供方便的方法收集 HTTP 请求期间发生的错误。最终中间件可以将他们写入日志、数据库或网络发送。
* 内置渲染：提供 JSON、XML 和 HTML 渲染 API。
* 扩展性：中间件创建简单

## API 设计

创建一个唱片管理仓库，客户端可以通过接口添加和获取数据。接口定义：

`/albums`
  * GET - 获取专辑列表，以 JSON 形式返回。
  * POST - 添加新专辑，以 JSON 形式发送请求数据

`/albums/:id`
  * GET - 根据 ID 获取专辑信息，以 JSON 形式返回。

## 代码

### 1. 创建数据

定义结构体：

```go
type album struct {
    ID string 	`json:"id"`
    Title string `json:"title"`
    Artist string `json:"artist"`
    Price float64 `json:"price"`
}
// 测试数据
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
```
`json:"title"` 用于指定结构在序列号为 JSON 时的字段名称。

### 2. 创建返回所有专辑的处理函数

```go
package main

// ...省略代码

// 处理函数
func getAlbums(c *gin.Context) {
	// 读取 albums 切片数据，写入到 JSON 中
	c.IndentedJSON(http.StatusOK, albums)
}
// 路由注册
func main() {
    router := gin.Default()
    // 注册 GET 请求
    router.GET("/albums", getAlbums)
    router.Run("localhost:8080")
}
```
上面使用 Gin 相关的 API，挨个来看是什么意思？

先看 `getAlbums` 函数接收一个 `gin.Context` 参数。
通过该参数我们可以在中间件之间传递参数，管理数据流、验证请求JSON 和 呈现 JSON 的响应等。
`c.IndentedJSON()` 和 `c.JSON()` 两个函数都可以将给定的结构作为 JSON 序列号到响应体中，
他们还将 `Content-Type` 设置为 `application/json`。 
两者不同点，前者会对生成的 JSON 序列进行格式化缩进等美化操作。

下面看看 main 函数，`gin.Default()` 函数创建 Engine 实例。
内部添加了 `Logger, Recovery` 中间件。前者用于日志写入、后者用于处理恐慌恢复。

`router.GET()` 注册 GET 路由

`router.Run()` 连接 HTTP 协议，开始监听 HTTP 请求。除非发生错误，否则不会停止运行。

### 3. 新增专辑

```go
// 3. 编写新增专辑的处理方法
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "album error"})
		return
	}
	albums = append(albums, newAlbum)
	c.JSON(http.StatusOK, newAlbum)
}

func main()  {
    // ...
	router.POST("/albums", postAlbums)
	// ...
}
```
`err := c.BindJSON(&newAlbum)` 将请求体 JSON 绑定到 newAlbum 变量上。

### 4. 根据 ID 查询专辑

```go
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
	// ...
    router.GET("/albums/:id", getAlbumById)
    // ...
}
```