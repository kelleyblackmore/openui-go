package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"openwebui-go/internal/server"
	"openwebui-go/internal/backend"

	"github.com/sirupsen/logrus"
)

//go:embed all:assets/frontend
var frontendAssets embed.FS

var (
	port = flag.Int("port", 8080, "Port to serve the frontend on")
	backendPort = flag.Int("backend-port", 11434, "Port for the backend API")
	debug = flag.Bool("debug", false, "Enable debug logging")
)

func main() {
	flag.Parse()

	// Configure logging
	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.Info("Starting OpenWebUI Go binary...")

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start backend process
	backendManager := backend.NewManager(*backendPort)
	if err := backendManager.Start(ctx); err != nil {
		logrus.Fatalf("Failed to start backend: %v", err)
	}
	defer backendManager.Stop()

	// Create and start frontend server
	frontendServer := server.NewFrontendServer(*port, frontendAssets)
	
	// Start server in goroutine
	go func() {
		logrus.Infof("Frontend server starting on port %d", *port)
		if err := frontendServer.Start(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Frontend server failed: %v", err)
		}
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	logrus.Info("Shutting down...")
	
	// Graceful shutdown
	if err := frontendServer.Shutdown(ctx); err != nil {
		logrus.Errorf("Error during frontend server shutdown: %v", err)
	}
} 