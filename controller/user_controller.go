package controller

import (
	"context"
	"encoding/json"
	"go-es/config"
	"go-es/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var indexName string = "user"

func FindUserById(c *gin.Context) {
	id := c.Param("id")
	var user model.User

	res, err := config.Client.Get().Index(indexName).Id(id).Do(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := json.Unmarshal(res.Source, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func SaveUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = uuid.New().String()
	_, err := config.Client.Index().Index(indexName).Id(user.ID).BodyJson(&user).Do(context.Background())
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.Client.Update().Index("user").Id(user.ID).Doc(&user).Do(context.Background())
	if err != nil {
		c.JSON(http.StatusNotModified, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	_, err := config.Client.Delete().Index(indexName).Id(id).Do(context.Background())
	if err != nil {
		c.JSON(http.StatusNotModified, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
