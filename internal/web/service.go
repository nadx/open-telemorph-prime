package web

import (
	"net/http"
	"strconv"

	"open-telemorph-prime/internal/config"
	"open-telemorph-prime/internal/storage"

	"github.com/gin-gonic/gin"
)

type Service struct {
	storage storage.Storage
	config  config.WebConfig
}

func NewService(storage storage.Storage, config config.WebConfig) *Service {
	return &Service{
		storage: storage,
		config:  config,
	}
}

// API endpoints
func (s *Service) GetMetrics(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	metrics, err := s.storage.GetMetrics(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   metrics,
		"total":  len(metrics),
		"limit":  limit,
		"offset": offset,
	})
}

func (s *Service) GetTraces(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	traces, err := s.storage.GetTraces(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   traces,
		"total":  len(traces),
		"limit":  limit,
		"offset": offset,
	})
}

func (s *Service) GetLogs(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	logs, err := s.storage.GetLogs(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   logs,
		"total":  len(logs),
		"limit":  limit,
		"offset": offset,
	})
}

func (s *Service) GetServices(c *gin.Context) {
	services, err := s.storage.GetServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"services": services,
	})
}

func (s *Service) Query(c *gin.Context) {
	var queryReq struct {
		Type   string `json:"type" binding:"required"`
		Query  string `json:"query" binding:"required"`
		Limit  int    `json:"limit"`
		Offset int    `json:"offset"`
	}

	if err := c.ShouldBindJSON(&queryReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if queryReq.Limit == 0 {
		queryReq.Limit = 100
	}

	switch queryReq.Type {
	case "metrics":
		metrics, err := s.storage.GetMetrics(queryReq.Limit, queryReq.Offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": metrics})
	case "traces":
		traces, err := s.storage.GetTraces(queryReq.Limit, queryReq.Offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": traces})
	case "logs":
		logs, err := s.storage.GetLogs(queryReq.Limit, queryReq.Offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": logs})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query type"})
	}
}

// Web UI handlers
func (s *Service) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": s.config.Title,
		"theme": s.config.Theme,
	})
}

func (s *Service) Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"title": s.config.Title + " - Dashboard",
		"theme": s.config.Theme,
	})
}

func (s *Service) MetricsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "metrics.html", gin.H{
		"title": s.config.Title + " - Metrics",
		"theme": s.config.Theme,
	})
}

func (s *Service) TracesPage(c *gin.Context) {
	c.HTML(http.StatusOK, "traces.html", gin.H{
		"title": s.config.Title + " - Traces",
		"theme": s.config.Theme,
	})
}

func (s *Service) LogsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "logs.html", gin.H{
		"title": s.config.Title + " - Logs",
		"theme": s.config.Theme,
	})
}

func (s *Service) ServicesPage(c *gin.Context) {
	c.HTML(http.StatusOK, "services.html", gin.H{
		"title": s.config.Title + " - Services",
		"theme": s.config.Theme,
	})
}

func (s *Service) AlertsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "alerts.html", gin.H{
		"title": s.config.Title + " - Alerts",
		"theme": s.config.Theme,
	})
}

func (s *Service) QueryPage(c *gin.Context) {
	c.HTML(http.StatusOK, "query.html", gin.H{
		"title": s.config.Title + " - Query Builder",
		"theme": s.config.Theme,
	})
}

func (s *Service) AdminPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", gin.H{
		"title": s.config.Title + " - Administration",
		"theme": s.config.Theme,
	})
}

// Admin API endpoints
func (s *Service) GetConfig(c *gin.Context) {
	// TODO: Implement config retrieval
	c.JSON(http.StatusOK, gin.H{
		"server": gin.H{
			"port":          8080,
			"read_timeout":  "5s",
			"write_timeout": "10s",
			"idle_timeout":  "120s",
		},
		"storage": gin.H{
			"type":           "sqlite",
			"path":           "./data/telemorph.db",
			"retention_days": 30,
		},
		"ingestion": gin.H{
			"api_endpoint":    "0.0.0.0:9013",
			"health_endpoint": "0.0.0.0:8080",
		},
		"web": gin.H{
			"port":      3000,
			"enable_ui": true,
			"theme":     "auto",
		},
		"logging": gin.H{
			"level":       "info",
			"format":      "console",
			"development": true,
		},
	})
}

func (s *Service) SaveConfig(c *gin.Context) {
	var config map[string]interface{}
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement config saving
	c.JSON(http.StatusOK, gin.H{"message": "Configuration saved successfully"})
}

func (s *Service) GetSystemStatus(c *gin.Context) {
	// TODO: Implement system status retrieval
	c.JSON(http.StatusOK, gin.H{
		"uptime":       "2h 15m 30s",
		"memory_usage": "45.2 MB",
		"storage_used": "128.5 MB",
		"status":       "healthy",
	})
}
