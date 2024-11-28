package controllers

import (
	"net/http"
	"terraform-manager/services"

	"github.com/gin-gonic/gin"
)

// 初始化 Terraform
func InitTerraform(c *gin.Context) {
	err := services.InitTerraform()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Terraform initialized successfully"})
}

// 应用 Terraform 配置
func ApplyTerraform(c *gin.Context) {
	err := services.ApplyTerraform()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Terraform applied successfully"})
}

// 销毁 Terraform 配置
func DestroyTerraform(c *gin.Context) {
	err := services.DestroyTerraform()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Terraform destroyed successfully"})
}
