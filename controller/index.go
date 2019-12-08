package controller

import (
	"net/http"

	"github.com/b3kt/account-srv/config"
	"github.com/gin-gonic/gin"
)

// IndexController is the default controller
type IndexController struct{}

// GetIndex home page
func (ctrl *IndexController) GetIndex(c *gin.Context) {
	// c.HTML(http.StatusOK, "index.html", gin.H{
	// 	"title":   "Gin Skeleton",
	// 	"content": "This is a skeleton based on gin framework",
	// })
	c.JSON(http.StatusOK, gin.H{
		"hi": config.Server.Version,
	})
}

// GetVersion version json
func (ctrl *IndexController) GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": config.Server.Version,
	})
}
