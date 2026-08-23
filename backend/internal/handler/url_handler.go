package handler

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"shortavee/backend/internal/metrics"
	"shortavee/backend/internal/service"
	"shortavee/backend/pkg/redis"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(service *service.URLService) *URLHandler {
	return &URLHandler{service: service}
}

func (h *URLHandler) CreateShortURL(c *gin.Context) {

	var req struct {
		URL string `json:"url"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	parsedURL, err := url.ParseRequestURI(req.URL)
	if err != nil ||
		(parsedURL.Scheme != "http" && parsedURL.Scheme != "https") ||
		parsedURL.Host == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "URL must start with http:// or https://",
		})
		return
	}

	userIDValue, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user ID missing in token"})
		return
	}

	userIDStr, ok := userIDValue.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user ID in token"})
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user ID format"})
		return
	}

	shortURL, err := h.service.CreateShortURL(req.URL, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	metrics.URLsCreated.Inc()

	c.JSON(http.StatusCreated, gin.H{
		"short_url": "http://localhost:8080/" + shortURL.ShortCode,
	})
}

func (h *URLHandler) RedirectURL(c *gin.Context) {

	code := c.Param("code")

	ctx := context.Background()

	cachedURL, err := redis.Client.Get(ctx, code).Result()

	if err == nil {
		c.Redirect(http.StatusMovedPermanently, cachedURL)
		return
	}

	url, err := h.service.GetOriginalURL(code)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "URL not found",
		})
		return
	}

	err = redis.Client.Set(
		ctx,
		code,
		url.OriginalURL,
		24*time.Hour,
	).Err()

	if err != nil {
		// tidak perlu gagal request
		// cukup log saja
		println("Redis set failed:", err.Error())
	}

	h.service.IncrementClickCount(url.ID)
	metrics.Redirects.Inc()

	c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}

func (h *URLHandler) GetAllURLs(c *gin.Context) {

	userIDValue, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user ID missing in token"})
		return
	}

	userIDStr, ok := userIDValue.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user ID in token"})
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user ID format"})
		return
	}

	urls, err := h.service.GetURLsByUser(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, urls)
}

func (h *URLHandler) DeleteURL(c *gin.Context) {

	id := c.Param("id")

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	err := h.service.DeleteURL(
		id,
		userID.(string),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "URL deleted successfully",
	})
}
