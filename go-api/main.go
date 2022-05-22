package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var todos []string

func Lists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"list": todos})
}

func ListItem(c *gin.Context) {
	errormessage := "Index out of range"
	indexstring := c.Param("index")
	index, err := strconv.Atoi(indexstring)
	if err == nil && index < len(todos) {
		c.JSON(http.StatusOK, gin.H{"item": todos[index]})
	} else {
		if err != nil {
			errormessage = "Number expected: " + indexstring
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errormessage})
	}
}

func AddListItem(c *gin.Context) {
	item := c.PostForm("item")
	todos = append(todos, item)
	c.JSON(http.StatusCreated, c.FullPath()+"/"+strconv.Itoa(len(todos)-1))
}

func DeleteListItem(c *gin.Context) {
	errormessage := "Index out of range"
	indexstring := c.Param("index")
	index, err := strconv.Atoi(indexstring)
	if err == nil && index < len(todos) {
		todos = append(todos[:index], todos[index+1:]...)
		c.JSON(http.StatusOK, gin.H{"item": todos})
	} else {
		if err != nil {
			errormessage = "Number expected: " + indexstring
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errormessage})
	}
}

func UpdateListItem(c *gin.Context) {
	errormessage := "Index out of range"
	indexstring := c.Param("index")
	index, err := strconv.Atoi(indexstring)
	if err == nil && index < len(todos) {
		item := c.PostForm("item")
		todos[index] = item
		c.JSON(http.StatusOK, gin.H{"item": todos})
	} else {
		if err != nil {
			errormessage = "Number expected: " + indexstring
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errormessage})
	}
}

func main() {
	todos = append(todos, "Write the application")
	r := gin.Default()
	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	r.Use(cors.Default())
	r.GET("/api/lists", Lists)
	r.GET("/api/lists/:index", ListItem)
	r.POST("/api/lists", AddListItem)
	r.DELETE("/api/lists/:index", DeleteListItem)
	r.PUT("/api/lists/:index", UpdateListItem)
	r.Run(":6060")
}
