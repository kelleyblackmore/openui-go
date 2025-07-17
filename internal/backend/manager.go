package backend

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

type Manager struct {
	port     int
	cmd      *exec.Cmd
	ctx      context.Context
	cancel   context.CancelFunc
}

func NewManager(port int) *Manager {
	return &Manager{
		port: port,
	}
}

func (m *Manager) Start(ctx context.Context) error {
	m.ctx, m.cancel = context.WithCancel(ctx)

	// For now, we'll assume Python is available
	// In Phase 3, this will be replaced with a bundled binary
	backendPath := m.getBackendPath()
	
	if backendPath == "" {
		return fmt.Errorf("backend not found. Please ensure the backend is properly bundled")
	}

	// Set environment variables for the backend
	env := os.Environ()
	env = append(env, fmt.Sprintf("PORT=%d", m.port))
	env = append(env, "HOST=127.0.0.1")

	// Start the backend process
	m.cmd = exec.CommandContext(m.ctx, "python3", backendPath)
	m.cmd.Env = env
	m.cmd.Stdout = os.Stdout
	m.cmd.Stderr = os.Stderr

	logrus.Infof("Starting backend on port %d", m.port)
	
	if err := m.cmd.Start(); err != nil {
		return fmt.Errorf("failed to start backend: %v", err)
	}

	// Wait a bit for the backend to start
	time.Sleep(2 * time.Second)

	// Check if process is still running
	if m.cmd.ProcessState != nil && m.cmd.ProcessState.Exited() {
		return fmt.Errorf("backend process exited unexpectedly")
	}

	logrus.Info("Backend started successfully")
	return nil
}

func (m *Manager) Stop() {
	if m.cancel != nil {
		m.cancel()
	}
	
	if m.cmd != nil && m.cmd.Process != nil {
		logrus.Info("Stopping backend...")
		m.cmd.Process.Kill()
		m.cmd.Wait()
	}
}

func (m *Manager) getBackendPath() string {
	// For Phase 1, we'll look for a Python backend
	// In Phase 3, this will return the path to the bundled binary
	
	// Check if we have a backend directory
	backendDir := "backend"
	if _, err := os.Stat(backendDir); err == nil {
		mainPy := filepath.Join(backendDir, "main.py")
		if _, err := os.Stat(mainPy); err == nil {
			return mainPy
		}
	}

	// Check if we have a bundled backend binary
	binaryName := "backend"
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}
	
	binaryPath := filepath.Join("assets", "backend", binaryName)
	if _, err := os.Stat(binaryPath); err == nil {
		return binaryPath
	}

	return ""
}

func (m *Manager) IsHealthy() bool {
	// TODO: Implement health check by calling backend health endpoint
	// For now, just check if the process is running
	if m.cmd == nil || m.cmd.Process == nil {
		return false
	}
	
	// Check if process is still running
	if m.cmd.ProcessState != nil && m.cmd.ProcessState.Exited() {
		return false
	}
	
	return true
} 