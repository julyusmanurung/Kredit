package laporan

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

func (h *Handler) GetLaporanData(c *gin.Context) {
	branch := c.Query("branch")
	channelingCompany := c.Query("channeling_company")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	req := DataRequest{Branch: branch, ChannelingCompany: channelingCompany, StartDate: startDate, EndDate: endDate}
	laporan, status, err := h.Service.GetLaporanData(req)

	if err != nil {
		c.JSON(status, gin.H{
			"message": "get laporan data error",
			"data":    laporan,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "get laporan data success",
		"data":    laporan,
	})
}
