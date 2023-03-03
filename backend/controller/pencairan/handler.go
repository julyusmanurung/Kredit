package pencairan

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

func (h *Handler) GetRecentCreditApplicant(c *gin.Context) {
	recentCreditApplicant, status, err := h.Service.GetRecentCreditApplicant()
	if err != nil {
		c.JSON(status, gin.H{
			"message": "get recent credit applicant error",
			"data":    recentCreditApplicant,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "get recent credit applicant success",
		"data":    recentCreditApplicant,
	})
}

func (h *Handler) GetAllBranch(c *gin.Context) {
	branch, status, err := h.Service.GetAllBranch()
	if err != nil {
		c.JSON(status, gin.H{
			"message": "get recent credit applicant error",
			"data":    branch,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "get all branch success",
		"data":    branch,
	})
}

func (h *Handler) GetAllCompany(c *gin.Context) {
	company, status, err := h.Service.GetAllCompany()
	if err != nil {
		c.JSON(status, gin.H{
			"message": "get all company error",
			"data":    company,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "get all company success",
		"data":    company,
	})
}

func (h *Handler) GetAllApprovalStatusNine(c *gin.Context) {
	approval, status, err := h.Service.GetAllApprovalStatusNine()
	if err != nil {
		c.JSON(status, gin.H{
			"message": "get approval status error",
			"data":    approval,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "get all company success",
		"data":    approval,
	})
}

func (h *Handler) GetAllApprovalStatusNineFilter(c *gin.Context) {
	branch := c.Query("branch")
	channelingCompany := c.Query("channeling_company")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	req := DataRequest{Branch: branch, ChannelingCompany: channelingCompany, StartDate: startDate, EndDate: endDate}
	approval, status, err := h.Service.GetAllApprovalStatusNineFilter(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": "get approval status error",
			"data":    approval,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "get all company success",
		"data":    approval,
	})
}

func (h *Handler) UpdateApprovalStatus(c *gin.Context) {
	var req DataPostPPK

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
		return
	}

	result := h.Service.UpdateApprovalStatus(req)

	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Gagal Mengambil Data",
			"error":   "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sukses mengupdate approval status :D",
	})

}
