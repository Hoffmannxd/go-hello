package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func registerRoutes() *gin.Engine {
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

	////////
	r.GET("/rain-probability/:seed", func(c *gin.Context) {
		seed := c.Param("seed")
		fmt.Println("Reveived seed :", seed)
	})


	////////

	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	admin.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	r.Static("/public", "./public")
	//r.StaticFS("/public", httpDir("./public")) most verbose

	return r
}
