package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserProfile(c *gin.Context) {
	c.JSON(http.StatusOK, "usercreated")
}

func UpdateUserProfile(c *gin.Context) {
	c.JSON(http.StatusOK, "userUpdated")
}

func GetUserProfile(c *gin.Context) {
	c.JSON(http.StatusOK, "userDetails")
}

func GetUserProfileGraph(c *gin.Context) {
	c.JSON(http.StatusOK, "userProfile")
}
