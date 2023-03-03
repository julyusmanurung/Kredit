package dashboard

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

func (h *Handler) GetCustomerData(c *gin.Context) {
	data, status, err := h.Service.GetCustomerData()
	if err != nil {
		c.JSON(status, gin.H{
			"message": "get all customer data error",
			"data":    data,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "get all customer data success",
		"data":    data,
	})
}
