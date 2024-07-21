package go_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marciojr/API/data"
)

func Initialize() {
	r := gin.Default()
	r.GET("/users", getUsers)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.GetUsers())
}
