package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html",nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "It's alive, %v", "Hoffmann")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	admin := r.Group("/admin")
	admin.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	r.Static("/public", "./public")
    //r.StaticFS("/public", httpDir("./public")) most verbose

	r.Run(":3000") // listen and serve on localhost:8080
}