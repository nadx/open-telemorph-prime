package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"open-telemorph-prime/internal/config"
	"open-telemorph-prime/internal/dogfood"
	"open-telemorph-prime/internal/ingestion"
	"open-telemorph-prime/internal/storage"
	"open-telemorph-prime/internal/web"

	"github.com/gin-gonic/gin"
)

var (
	configPath = flag.String("config", "config.yaml", "Path to configuration file")
	version    = "0.2.1"
)

func main() {
	flag.Parse()

	// Load configuration
	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize storage
	storage, err := storage.NewSQLiteStorage(cfg.Storage)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}
	defer storage.Close()

	// Initialize ingestion service
	ingestionService := ingestion.NewService(storage, cfg.Ingestion)

	// Initialize web service
	webService := web.NewService(storage, cfg.Web, version)

	// Initialize dogfood service
	dogfoodService := dogfood.NewService(cfg.Web, storage, cfg.Server.Port)

	// Set up Gin router
	if cfg.Server.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())

	// Load HTML templates
	router.LoadHTMLGlob("web/*.html")

	// Register routes
	registerRoutes(router, ingestionService, webService, dogfoodService)

	// Create HTTP server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	// Start ingestion service
	go func() {
		if err := ingestionService.Start(); err != nil {
			log.Fatalf("Failed to start ingestion service: %v", err)
		}
	}()

	// Start dogfood service
	go func() {
		ctx := context.Background()
		dogfoodService.Start(ctx)
	}()

	// Start HTTP server
	go func() {
		log.Printf("Starting Open-Telemorph-Prime server on port %d", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down Open-Telemorph-Prime...")

	// Shutdown ingestion service
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := ingestionService.Stop(ctx); err != nil {
		log.Printf("Error stopping ingestion service: %v", err)
	}

	// Shutdown HTTP server
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down server: %v", err)
	}

	log.Println("Open-Telemorph-Prime stopped")
}

func registerRoutes(router *gin.Engine, ingestionService *ingestion.Service, webService *web.Service, dogfoodService *dogfood.Service) {
	// Health endpoints
	router.GET("/health", healthCheck)
	router.GET("/ready", readinessCheck)

	// API routes
	api := router.Group("/api/v1")
	{
		api.GET("/metrics", webService.GetMetrics)
		api.GET("/traces", webService.GetTraces)
		api.GET("/logs", webService.GetLogs)
		api.GET("/services", webService.GetServices)
		api.POST("/query", webService.Query)
	}

	// Admin API routes
	admin := router.Group("/api/v1/admin")
	{
		admin.GET("/config", webService.GetConfig)
		admin.POST("/config", webService.SaveConfig)
		admin.GET("/status", webService.GetSystemStatus)
		admin.GET("/dogfood", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"enabled": dogfoodService.IsEnabled()})
		})
		admin.POST("/dogfood", func(c *gin.Context) {
			var req struct {
				Enabled bool `json:"enabled"`
			}
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			dogfoodService.SetEnabled(req.Enabled)
			c.JSON(http.StatusOK, gin.H{"message": "Dogfood mode updated", "enabled": req.Enabled})
		})
	}

	// OTLP endpoints are now served on dedicated ingestion ports (4317/4318)
	// These are handled by the ingestion service directly

	// Web UI
	router.Static("/static", "./web/static")
	router.StaticFile("/open-telemorph-logo.svg", "./web/open-telemorph-logo.svg")
	router.GET("/", webService.Index)
	router.GET("/dashboard", webService.Dashboard)
	router.GET("/metrics", webService.MetricsPage)
	router.GET("/traces", webService.TracesPage)
	router.GET("/logs", webService.LogsPage)
	router.GET("/services", webService.ServicesPage)
	router.GET("/alerts", webService.AlertsPage)
	router.GET("/query", webService.QueryPage)
	router.GET("/admin", webService.AdminPage)
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().Unix(),
		"version":   version,
	})
}

func readinessCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ready",
		"timestamp": time.Now().Unix(),
		"version":   version,
	})
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
