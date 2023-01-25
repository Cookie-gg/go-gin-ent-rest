package handler

import (
	"go-gin-ent-rest/ent"
	"go-gin-ent-rest/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func CreateUserHandler(client *ent.Client) *UserHandler {
	return &UserHandler{service.CreateUserService(client)}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	user := ent.User{}
	if err := c.BindJSON(&user); err != nil {
		log.Fatalf("failed binding json: %v", err)
	}
	if res, err := h.userService.Create(&user).Save(c); err == nil {
		c.JSON(http.StatusCreated, res)
	} else {
		log.Fatalf("failed creating user: %v", err)
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	user := ent.User{}
	var err error

	if user.ID, err = strconv.Atoi(c.Param("id")); err != nil {
		log.Fatalf("failed getting params: %v", err)
	}
	if res, err := h.userService.Get(&user).Only(c); err != nil {
		log.Fatalf("failed creating user: %v", err)
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	user := ent.User{}
	var err error

	if user.ID, err = strconv.Atoi(c.Param("id")); err != nil {
		log.Fatalf("failed getting params: %v", err)
	}
	if err := c.BindJSON(&user); err != nil {
		log.Fatalf("failed binding json: %v", err)
	}
	if _, err := h.userService.Update(&user).Save(c); err != nil {
		log.Fatalf("failed creating user: %v", err)
	} else {
		c.Status(http.StatusNoContent)
	}
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	user := ent.User{}
	var err error

	if user.ID, err = strconv.Atoi(c.Param("id")); err != nil {
		log.Fatalf("failed getting params: %v", err)
	}
	if err := h.userService.Delete(&user).Exec(c); err != nil {
		log.Fatalf("failed creating user: %v", err)
	} else {
		c.Status(http.StatusNoContent)
	}
}
