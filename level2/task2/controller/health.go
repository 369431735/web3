package controller

import "github.com/gin-gonic/gin"

// HealthCheck godoc
// @Summary      健康检查接口
// @Description  检查系统是否正常运行的API接口
// @Tags         系统
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  map[string]string
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
