package server

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type FrontendServer struct {
	port   int
	assets embed.FS
	server *http.Server
}

func NewFrontendServer(port int, assets embed.FS) *FrontendServer {
	return &FrontendServer{
		port:   port,
		assets: assets,
	}
}

func (s *FrontendServer) Start() error {
	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)
	
	router := gin.New()
	router.Use(gin.Recovery())

	// Serve static assets
	s.setupStaticRoutes(router)
	
	// Setup SPA fallback for client-side routing
	s.setupSPARoutes(router)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: router,
	}

	return s.server.ListenAndServe()
}

func (s *FrontendServer) Shutdown(ctx context.Context) error {
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}
	return nil
}

func (s *FrontendServer) setupStaticRoutes(router *gin.Engine) {
	// Create a sub-filesystem for the frontend assets
	frontendFS, err := fs.Sub(s.assets, "assets/frontend")
	if err != nil {
		logrus.Fatalf("Failed to create frontend filesystem: %v", err)
	}

	// Serve static files (CSS, JS, images, etc.)
	router.StaticFS("/static", http.FS(frontendFS))
	
	// Serve other assets (favicon, etc.)
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.FileFromFS("favicon.ico", http.FS(frontendFS))
	})
	
	// Serve manifest.json if it exists
	router.GET("/manifest.json", func(c *gin.Context) {
		c.FileFromFS("manifest.json", http.FS(frontendFS))
	})
}

func (s *FrontendServer) setupSPARoutes(router *gin.Engine) {
	// Handle all other routes by serving index.html for SPA routing
	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		
		// Don't serve index.html for API routes or static assets
		if strings.HasPrefix(path, "/api/") || strings.HasPrefix(path, "/static/") {
			c.Status(http.StatusNotFound)
			return
		}

		// Serve index.html for all other routes
		c.FileFromFS("index.html", http.FS(s.assets))
	})
} 