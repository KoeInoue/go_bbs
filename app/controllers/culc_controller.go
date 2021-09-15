package controllers

import (
	"go_bbs/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CalcController struct{}

func (CalcController) CalcConcurrently(c *gin.Context) {
	queryDate := c.Query("date")

	var calcService services.CalcService
	if result, err := calcService.CalcConcurrently(queryDate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": "ok",
			"result": result,
		})
	}
}

func (CalcController) CalcInSeries(c *gin.Context) {
	queryDate := c.Query("date")

	var calcService services.CalcService
	if result, err := calcService.CalcInSeries(queryDate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": "ok",
			"result": result,
		})
	}
}
