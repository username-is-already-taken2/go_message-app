package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.New()

	// Manually add Logger and Recovery middleware
	r.Use(gin.Logger(), gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"}, // Add your frontend URL
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Accept", "Content-Type"},
		AllowCredentials: false,
	}))

	r.GET("/", s.MessageHandler)
	r.GET("/_/healthz", s.HealthzHandler)

	return r
}

func (*Server) MessageHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello from Mondoo Engineer!"

	c.JSON(http.StatusOK, resp)
}

func (*Server) HealthzHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["status"] = "ok"

	c.JSON(http.StatusOK, resp)
}
