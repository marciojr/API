package go_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marciojr/API/data"
)

func Initialize() {
	r := gin.Default()
	r.GET("/users", getUsers)
	r.POST("/users", postUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Users)
}

func postUser(c *gin.Context) {
	var newUser data.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	data.AddUser(newUser)

	c.IndentedJSON(http.StatusCreated, newUser)
}
