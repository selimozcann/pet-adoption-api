package pet

import (
	"net/http"
	"petapi/internal/db"

	"github.com/gin-gonic/gin"
)

func GetPets(c *gin.Context) {
	var pets []Pet
	db.DB.Find(&pets)
	c.JSON(http.StatusOK, pets)
}

func CreatePet(c *gin.Context) {
	var pet Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&pet)
	c.JSON(http.StatusCreated, pet)
}
