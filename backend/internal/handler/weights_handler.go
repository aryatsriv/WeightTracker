package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWeight(c *gin.Context) {
	c.JSON(http.StatusOK, "created")
}

func GetWeight(c *gin.Context) {
	c.JSON(http.StatusOK, "sent")
}

func DeleteWeightByID(c *gin.Context) {
	c.JSON(http.StatusOK, "deketed")
}
