package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemHandler struct {
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}

// @Summary	Check api server is still responsed
// @Description	for monitor to check service is activated
// @Tags system
// @Produce	plain
// @Success	200	string string "should be 'ok'"
// @Router /livez [get]
func (s *SystemHandler) Health(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
