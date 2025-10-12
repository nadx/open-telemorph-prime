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
