package auth

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

func (h *Handler) Login(c *gin.Context) {
	nik := c.Query("nik")
	password := c.Query("password")

	req := DataRequest{Nik: nik, Password: password}
	data, status, err := h.Service.Login(req)

	if err != nil {
		c.JSON(status, gin.H{
			"message": "login failed",
			"data":    data,
		})
		return
	}

	message := ""
	if data.Nik == "" {
		message = "wrong nik or password"
	} else {
		message = "login succcess"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
	})
}
