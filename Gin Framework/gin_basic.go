package main

import "github.com/gin-gonic/gin"

func main() {
	/*a new Gin router is created by calling gin.Default(),
	This router will handle incoming HTTP requests and route them to the appropriate handlers.*/
	r := gin.Default()

	// Defining a Route for "/hello"
	/*
		defines a route that responds to a GET request at the path "/hello".
		When a GET request is made to this path, the provided anonymous function is executed.
		Inside this function:
			It sends a JSON response with a status code of 200 (OK).
			The response includes a JSON object with a "message" field, set to "Hello, Gin!"*/

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, Gin!"})
	})

	// Define a route with a route parameter
	/*
		defines a route with a route parameter, denoted by :id. This means that the route
		can match URLs like "/users/123", "/users/456", and so on. The parameter value
		(e.g., "123") will be captured and can be accessed using c.Param("id").
		Inside the handler function:
			It captures the "id" parameter from the URL.
			It sends a JSON response that includes the captured "user_id" parameter.*/

	r.GET("/users/:id", func(c *gin.Context) {
		userID := c.Param("id") // Access the "id" parameter from the URL
		c.JSON(200, gin.H{"user_id": userID})
	})

	/*
		This line starts the Gin web server, and it will listen for incoming HTTP requests 
		on port 8080. You can access the defined routes by making requests to this server*/
	r.Run(":8080")
}
