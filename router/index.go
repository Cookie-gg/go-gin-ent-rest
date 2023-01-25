package router

import (
	"go-gin-ent-rest/ent"
	"go-gin-ent-rest/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(client *ent.Client) *gin.Engine {
	r := gin.Default()

	h := handler.CreateUserHandler(client)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World!"})
	})

	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetUser)
	r.PUT("/user/:id", h.UpdateUser)
	r.DELETE("/user/:id", h.DeleteUser)
	r.GET("/user/:id/profile", h.GetUserProfile)
	r.PUT("/user/:id/profile", h.UpdateUserProfile)

	return r
}
