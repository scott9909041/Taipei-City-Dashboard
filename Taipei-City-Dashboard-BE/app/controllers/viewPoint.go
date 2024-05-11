package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"TaipeiCityDashboardBE/app/models"

	"github.com/gin-gonic/gin"
)
type ViewPoint struct {
	ID      int       `json:"id"`
	UserID  int       `json:"user_id"`
	Center  []float32 `json:"center"`
	Zoom    float32   `json:"zoom"`
	Pitch   float32   `json:"pitch"`
	Bearing float32   `json:"bearing"`
	Name    string    `json:"name"`
}
func CreateViewPoint(c *gin.Context) {
	var viewPoint ViewPoint

	if err := c.ShouldBindJSON(&viewPoint); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err:= models.CreateViewPoint(viewPoint.UserID, convertFloat32ArrayToString(viewPoint.Center), viewPoint.Zoom, viewPoint.Pitch, viewPoint.Bearing, viewPoint.Name)
	if  err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Viewpoint created"})
}

func GetViewPointByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	viewPoint, err := models.GetViewPointByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, viewPoint)
}

func DeleteViewPoint(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid viewpoint ID"})
		return
	}

	err = models.DeleteViewPoint(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Viewpoint deleted"})
}

func convertFloat32ArrayToString(arr []float32) string {
	var strArr []string
	for _, num := range arr {
		strArr = append(strArr, strconv.FormatFloat(float64(num), 'f', -1, 32))
	}
	return "{" + strings.Join(strArr, ",") + "}"
}