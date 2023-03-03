package api

import (
	"github.com/julyusmanurung/Kredit/controller/angsuran"
	"github.com/julyusmanurung/Kredit/controller/pencairan"
	"github.com/robfig/cron/v3"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type server struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func MakeServer(db *gorm.DB) *server {
	s := &server{
		Router: gin.Default(),
		DB:     db,
	}

	pencairanService := pencairan.NewRepository(s.DB)
	angsuranService := angsuran.NewRepository(s.DB)

	c := cron.New()
	c.AddFunc("@every 30m", func() { pencairanService.GetRecentCreditApplicant() })
	c.AddFunc("@every 30m", func() { angsuranService.GetInstallmentScale() })
	c.Start()

	return s
}

func (s *server) RunServer() {
	s.SetupRouter()
	port := os.Getenv("PORT")

	if err := s.Router.Run(":" + port); err != nil {
		log.Panicln(err.Error())
	}
}
