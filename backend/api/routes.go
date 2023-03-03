package api

import (
	"github.com/gin-contrib/cors"
	"github.com/julyusmanurung/Kredit/controller/angsuran"
	"github.com/julyusmanurung/Kredit/controller/auth"
	"github.com/julyusmanurung/Kredit/controller/dashboard"
	"github.com/julyusmanurung/Kredit/controller/laporan"
	"github.com/julyusmanurung/Kredit/controller/pencairan"
	"github.com/julyusmanurung/Kredit/controller/user"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"*"},
	}))

	pencairanRepository := pencairan.NewRepository(s.DB)
	pencairanService := pencairan.NewService(pencairanRepository)
	pencairanHandler := pencairan.NewHandler(pencairanService)

	angsuranRepository := angsuran.NewRepository(s.DB)
	angsuranService := angsuran.NewService(angsuranRepository)
	angsuranHandler := angsuran.NewHandler(angsuranService)

	userRepository := user.NewRepository(s.DB)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	dashboardRepository := dashboard.NewRepository(s.DB)
	dashboardService := dashboard.NewService(dashboardRepository)
	dashboardHandler := dashboard.NewHandler(dashboardService)

	laporanRepository := laporan.NewRepository(s.DB)
	laporanService := laporan.NewService(laporanRepository)
	laporanHandler := laporan.NewHandler(laporanService)

	authRepository := auth.NewRepository(s.DB)
	authService := auth.NewService(authRepository)
	authHandler := auth.NewHandler(authService)

	s.Router.GET("/", pencairanHandler.GetRecentCreditApplicant)
	s.Router.GET("/branch", pencairanHandler.GetAllBranch)
	s.Router.GET("/company", pencairanHandler.GetAllCompany)
	s.Router.GET("/statusapprovalnine", pencairanHandler.GetAllApprovalStatusNine)
	s.Router.GET("/statusapprovalninewithfilter", pencairanHandler.GetAllApprovalStatusNineFilter)
	s.Router.PATCH("/updateapprovalstatus", pencairanHandler.UpdateApprovalStatus)

	s.Router.GET("/angsuran", angsuranHandler.GetInstallmentScale)

	s.Router.GET("/profile/:user_id", userHandler.GetUserDetails)
	s.Router.PATCH("/updatepassword/:user_id", userHandler.UpdatePassword)

	s.Router.GET("/customer", dashboardHandler.GetCustomerData)

	s.Router.GET("laporan", laporanHandler.GetLaporanData)

	s.Router.GET("/login", authHandler.Login)
}
