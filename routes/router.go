package routes

import (
	"petapi/internal/pet"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/pets", pet.GetPets)
		api.POST("/pets", pet.CreatePet)
	}
}
