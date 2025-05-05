package router

import (
	"backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		weights := v1.Group("/weights")
		{
			weights.POST("", handler.CreateWeight)
			weights.GET("", handler.GetWeight)
			weights.DELETE("/:id", handler.DeleteWeightByID)
		}

		userProfiles := v1.Group("/users")
		{
			userProfiles.GET("", handler.GetUserProfile)
			userProfiles.POST("", handler.CreateUserProfile)
			userProfiles.PUT("", handler.UpdateUserProfile)
			userProfiles.GET("/graph", handler.GetUserProfileGraph)
		}
	}

	return r
}
