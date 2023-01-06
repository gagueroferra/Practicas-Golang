package main

import "github.com/gin-gonic/gin"

func main() {
	server()
	saludo()
}

func server() {
	//crear un router
	router := gin.Default()
	//captura la solicitud GET
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})
	router.Run(":8080") //corremos el server
}

func saludo() {
	router := gin.Default()
	router.POST("/saludo", func(s *gin.Context) {
		s.JSON(200, gin.H{})
	})
}
