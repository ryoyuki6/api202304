package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// var IpAddress string = "localhost" // 頭文字を大文字にして公開オブジェクト
var IpAddress string = "18.183.1.167" // 頭文字を大文字にして公開オブジェクト

type HealthData struct {
	ID		int		`json:"id"`
	Month	int		`json:"month"`
	Day		int		`json:"day"`
	Weight	float64	`json:"weight"`
}

var health_data = []HealthData{
	{ID:1, Month:1,	Day:1, Weight:55.5},
	{ID:2, Month:1,	Day:2, Weight:66.6},
	{ID:3, Month:1,	Day:3, Weight:77.7},
}

func getHealthdata(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, health_data)
}

func getHealthdataById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	for _, t := range health_data {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "the ID data does not exit!!!"})
}

func postHealthdata(c *gin.Context) {
	var newHealthdata HealthData
	if err := c.BindJSON(&newHealthdata); err != nil {
		return
	}
	health_data = append(health_data, newHealthdata)
}

func patchHealthdata(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	var health_data_patch HealthData
	health_data_patch.ID = id
	if err = c.BindJSON(&health_data_patch); err != nil {
		return
	}

	for i, t := range health_data {
		if t.ID == id {
			health_data[i] = health_data_patch
			c.IndentedJSON(http.StatusOK, health_data_patch)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "the id's health data not fount"})
}

func deleteHealthdata(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	for i, t := range health_data {
		if t.ID == id {
			health_data = append(health_data[:i],health_data[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "health_data(" + strconv.Itoa(id) + ") is deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "health_data not found"})
}

func main() {
	router := gin.Default()

	// CORS 対応
	config := cors.DefaultConfig()
	ip := IpAddress
	endpoint := "http://" + ip + ":3000"
	config.AllowOrigins = []string{endpoint}
	router.Use(cors.New(config))

	router.GET("/health_data", getHealthdata)
	router.GET("/health_data/:id", getHealthdataById)
	router.POST("/health_data", postHealthdata)
	router.PATCH("/health_data/:id", patchHealthdata)
	router.DELETE("/health_data/:id", deleteHealthdata)

	// router.Run("localhost:8080")
	router.Run(":8081")
}
