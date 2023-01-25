package handler

import (
	"go-gin-ent-rest/ent"
	"go-gin-ent-rest/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService    *service.UserService
	profileService *service.ProfileService
}

func CreateUserHandler(client *ent.Client) *UserHandler {
	return &UserHandler{service.CreateUserService(client), service.CreateProfileService(client)}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	user := ent.User{}
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	profile := ent.Profile{
		Age:    user.Edges.Profile.Age,
		Gender: user.Edges.Profile.Gender,
	}
	if res, err := h.userService.Create(&user).Save(c); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		user = *res
	}

	if res, err := h.profileService.Create(&profile).SetUser(&user).Save(c); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		user.Edges.Profile = res
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	user := ent.User{}
	var err error

	if user.ID, err = strconv.Atoi(c.Param("id")); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if res, err := h.userService.Get(&user).Only(c); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (h *UserHandler) GetUserProfile(c *gin.Context) {
	user := ent.User{}

	if res, err := strconv.Atoi(c.Param("id")); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		user.ID = res
	}

	profile := ent.Profile{Edges: ent.ProfileEdges{User: &user}}

	if res, err := h.profileService.Get(&profile).Only(c); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	user := ent.User{}
	var err error

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.ID, err = strconv.Atoi(c.Param("id")); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	profile := ent.Profile{
		Age:    user.Edges.Profile.Age,
		Gender: user.Edges.Profile.Gender,
	}
	if _, err := h.userService.Update(&user).Save(c); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		profile.ID = user.Edges.Profile.ID
	}
	if _, err := h.profileService.Update(&profile).Save(c); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.Status(http.StatusNoContent)
	}
}

func (h *UserHandler) UpdateUserProfile(c *gin.Context) {
	user := ent.User{}

	profile := ent.Profile{}

	if err := c.BindJSON(&profile); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if res, err := strconv.Atoi(c.Param("id")); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		user.ID = res
	}
	profile.Edges = ent.ProfileEdges{User: &user}

	if _, err := h.profileService.Update(&profile).Save(c); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	user := ent.User{}
	var err error

	if user.ID, err = strconv.Atoi(c.Param("id")); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userService.Delete(&user).Exec(c); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.Status(http.StatusNoContent)
	}
}
