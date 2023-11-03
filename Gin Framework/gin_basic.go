package main

import "github.com/gin-gonic/gin"

func main() {
	/*a new Gin router is created by calling gin.Default(), This router will handle incoming 
	HTTP requests and route them to the appropriate handlers.*/
    r := gin.Default()
	
    // Defining a Route for "/hello"
    r.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello, Gin!"})
    })

    // Define a route with a route parameter
    r.GET("/users/:id", func(c *gin.Context) {
        userID := c.Param("id") // Access the "id" parameter from the URL
        c.JSON(200, gin.H{"user_id": userID})
    })

    r.Run(":8080")
}
