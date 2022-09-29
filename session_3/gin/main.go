package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// let's add a middleware to count request
	r.Use(counter())
	// if you open "localhost:8080/ping"
	// it would reply with JSON '{"message":"pong"}'
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"counter": c.GetInt("counter"),
		}) // gin.H is just a map[string]interface{}
	})
	// will listen at 0.0.0.0:8080 and block the main goroutine
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func counter() func(c *gin.Context) {
	i := 0

	return func(c *gin.Context) {
		i++

		c.Set("counter", i)
	}
}
