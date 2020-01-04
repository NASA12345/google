package main

import ("github.com/gin-gonic/gin"
        "net/http"
)
func main() {
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
  router.GET("/greet/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello " + name + " !")
	})
	router.Run(":3000")
}
