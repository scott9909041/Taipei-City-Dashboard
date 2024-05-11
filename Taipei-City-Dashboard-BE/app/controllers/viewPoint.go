package controllers

import (
	"net/http"
	"strconv"

	"TaipeiCityDashboardBE/app/models"

	"github.com/gin-gonic/gin"
)

func CreateViewPoint(c *gin.Context) {
	var viewPoint models.ViewPoint

	if err := c.ShouldBindJSON(&viewPoint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	viewPoint,err := models.CreateViewPoint(viewPoint.UserID, viewPoint.CenterX, viewPoint.CenterY, viewPoint.Zoom, viewPoint.Pitch, viewPoint.Bearing, viewPoint.Name, viewPoint.PointType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Viewpoint created", "data": viewPoint})
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

// func convertFloat32ArrayToString(arr []float32) string {
// 	var strArr []string
// 	for _, num := range arr {
// 		strArr = append(strArr, strconv.FormatFloat(float64(num), 'f', -1, 32))
// 	}
// 	return "{" + strings.Join(strArr, ",") + "}"
// }
