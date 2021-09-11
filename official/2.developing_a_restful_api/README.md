# 使用 Go 和 Gin 开发 RESTful API 

介绍使用 Go 和 Gin Web 框架编写 RESTful Web 服务 API 的基础知识

> [Gin Web Framework](https://gin-gonic.com/docs/)

## API 设计

创建一个唱片管理仓库，客户端可以通过接口添加和获取数据。接口定义：

`/albums`
  * GET - 获取专辑列表，以 JSON 形式返回。
  * POST - 添加新专辑，以 JSON 形式发送请求数据

`/albums/:id`
  * GET - 根据 ID 获取专辑信息，以 JSON 形式返回。

## 代码