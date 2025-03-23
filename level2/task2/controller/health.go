package controller

import "github.com/gin-gonic/gin"

// HealthCheck godoc
// @Summary      闂佺顑冮崕閬嶅箖瀹ュ憘娑㈠焵椤掑嫬钃?// @Description  濠碘槅鍋€閸嬫捇鏌＄仦璇插姕婵犫偓閸ヮ剙绀夐柍鈺佸暞绗戦梺鍛婄啲闂勫嫰顢楀鍛殰婵繂鐬肩粻銉╂偠?// @Tags         缂備緡鍨靛畷鐢靛垝?// @Accept       application/json
// @Produce      application/json
// @Success      200  {object}  map[string]string
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
