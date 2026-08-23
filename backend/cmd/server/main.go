package main

import (
	"fmt"
	"os"
	"shortavee/backend/internal/handler"
	"shortavee/backend/internal/metrics"
	"shortavee/backend/internal/middleware"
	"shortavee/backend/internal/model"
	"shortavee/backend/internal/repository"
	"shortavee/backend/internal/service"
	"shortavee/backend/pkg/database"
	"shortavee/backend/pkg/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"shortavee/backend/pkg/redis"
)

func main() {

	db, err := database.Connect()
	redis.Connect()
	if err != nil {
		panic(err)
	}

	// --- PERBAIKAN DI SINI ---
	// Ambil sql.DB dari gorm.DB untuk bisa menggunakan Close()
	sqlDB, err := db.DB()
	if err != nil {
		panic("Gagal mendapatkan objek sql.DB: " + err.Error())
	}
	defer sqlDB.Close()
	// -------------------------

	db.AutoMigrate(
		&model.User{},
		&model.URL{},
	)

	fmt.Println(utils.GenerateShortCode(6))

	repo := repository.NewURLRepository(db)
	svc := service.NewURLService(repo)
	h := handler.NewURLHandler(svc)

	userRepo := repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepo)

	authHandler := handler.NewAuthHandler(authService)

	router := gin.Default()

	// 1. Ambil URL frontend dari environment variable
	frontendURL := os.Getenv("FRONTEND_URL")

	// 2. Jika kosong (saat di komputer lokal), gunakan localhost sebagai cadangan
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}

	router.Use(cors.New(cors.Config{
		// 3. Masukkan variabel frontendURL ke dalam array AllowOrigins
		AllowOrigins: []string{
			frontendURL,
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS", // Tambahkan OPTIONS untuk menangani preflight request dari browser
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
	}))

	authorized := router.Group("/api")
	authorized.Use(middleware.JWTAuth())

	authorized.POST("/shorten", h.CreateShortURL)
	authorized.GET("/urls", h.GetAllURLs)
	authorized.GET("/me", authHandler.GetMe)
	authorized.DELETE("/urls/:id", h.DeleteURL)

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.GET("/:code", h.RedirectURL)
	router.POST("/api/register", authHandler.Register)
	router.POST("/api/login", authHandler.Login)

	metrics.Register()
	router.Run(":8080")

}
