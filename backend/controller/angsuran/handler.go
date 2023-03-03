package angsuran

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetInstallmentScale(c *gin.Context) {
	installmentScale, err := h.Service.GetInstallmentScale()
	if err != nil {
		c.JSON(1, gin.H{
			"message": "get recent credit applicant error",
			"data":    installmentScale,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "get recent credit applicant success",
		"data":    installmentScale,
	})
}
