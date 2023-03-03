package user

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

func (h *Handler) GetUserDetails(c *gin.Context) {
	idUser := c.Param("user_id")

	res, err := h.Service.GetUserDetails(idUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    res,
	})
}

func (h *Handler) UpdatePassword(c *gin.Context) {
	idUser := c.Param("user_id")
	oldPassword := c.Query("old_password")
	newPassword := c.Query("new_password")

	req := DataRequestUpdatePassword{OldPassword: oldPassword, NewPassword: newPassword}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
		return
	}

	res, status, err := h.Service.UpdatePassword(idUser, req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": "password gagal update",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}
