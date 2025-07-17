# OpenWebUI Go

A Go-based wrapper for OpenWebUI that packages both the frontend and backend into a single, distributable binary.

## 🎯 Project Goals

- **Single Binary**: Package OpenWebUI frontend and backend into one Go executable
- **Cross-Platform**: Support Linux, macOS, and Windows
- **Easy Distribution**: No Python/Node.js dependencies required at runtime
- **Simple Deployment**: Just download and run

## 🏗️ Architecture

```
┌─────────────────────────────────────┐
│           Go Binary                 │
│  ┌─────────────────────────────────┐ │
│  │      Frontend Server            │ │
│  │   (Embedded React App)          │ │
│  └─────────────────────────────────┘ │
│  ┌─────────────────────────────────┐ │
│  │      Backend Manager            │ │
│  │   (Python FastAPI Process)      │ │
│  └─────────────────────────────────┘ │
└─────────────────────────────────────┘
```

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or later
- Git
- Node.js and npm (for building frontend)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/your-username/openwebui-go.git
   cd openwebui-go
   ```

2. **Prepare the frontend**
   ```bash
   ./scripts/prepare_frontend.sh
   ```

3. **Build the application**
   ```bash
   make build
   ```

4. **Run the application**
   ```bash
   ./bin/openwebui-go
   ```

5. **Open your browser**
   Navigate to `http://localhost:8080`

## 📋 Development

### Project Structure

```
openwebui-go/
├── cmd/
│   └── main.go                # Application entry point
├── internal/
│   ├── server/                # Frontend server
│   │   └── frontend.go
│   └── backend/               # Backend process manager
│       └── manager.go
├── assets/
│   └── frontend/              # Embedded frontend assets
├── scripts/
│   └── prepare_frontend.sh    # Frontend preparation script
├── backend/                   # Python backend (Phase 2)
├── Makefile                   # Build automation
└── README.md
```

### Available Commands

```bash
# Build the application
make build

# Build for all platforms
make build-all

# Run the application
make run

# Run with debug mode
make run-debug

# Development mode (with hot reload)
make dev

# Run tests
make test

# Clean build artifacts
make clean

# Install globally
make install

# Show help
make help
```

### Development Workflow

1. **Frontend Development**
   ```bash
   # Prepare frontend (first time only)
   ./scripts/prepare_frontend.sh
   
   # Make changes to assets/frontend/
   # Rebuild and test
   make build && make run
   ```

2. **Backend Development**
   ```bash
   # Add Python backend files to backend/
   # Modify internal/backend/manager.go as needed
   make build && make run
   ```

## 🔧 Configuration

### Command Line Options

```bash
./openwebui-go [options]

Options:
  --port int         Port to serve the frontend on (default 8080)
  --backend-port int Port for the backend API (default 11434)
  --debug           Enable debug logging
  --help            Show help
```

### Environment Variables

- `PORT`: Frontend server port (overrides --port)
- `BACKEND_PORT`: Backend API port (overrides --backend-port)
- `DEBUG`: Enable debug mode (overrides --debug)

## 📦 Distribution

### Building for Different Platforms

```bash
# Build for current platform
make build

# Build for all platforms
make build-all
```

This creates binaries for:
- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)

### Creating Installers

*Coming in Phase 4*

## 🧪 Testing

### Manual Testing

1. **Frontend Test**
   ```bash
   make run
   # Open http://localhost:8080
   # Verify frontend loads correctly
   ```

2. **Backend Test**
   ```bash
   make run-debug
   # Check logs for backend startup
   # Verify backend health endpoint
   ```

3. **Integration Test**
   ```bash
   make run
   # Test frontend-backend communication
   # Verify API endpoints work
   ```

### Automated Testing

```bash
make test
```

## 🔄 Project Phases

### ✅ Phase 1: Frontend Integration
- [x] Basic Go project structure
- [x] Frontend server with embedded assets
- [x] Placeholder frontend
- [ ] Clone and embed OpenWebUI frontend
- [ ] Test frontend serving

### 🔄 Phase 2: Backend Integration
- [ ] Extract OpenWebUI backend
- [ ] Python subprocess management
- [ ] Backend health checks
- [ ] Graceful shutdown

### ⏳ Phase 3: Backend Packaging
- [ ] PyInstaller integration
- [ ] Standalone backend binary
- [ ] Remove Python dependency

### ⏳ Phase 4: CLI & Installer
- [ ] Enhanced CLI options
- [ ] Installer scripts
- [ ] Auto-update functionality

### ⏳ Phase 5: Cross-platform
- [ ] Platform-specific builds
- [ ] Comprehensive testing
- [ ] Release automation

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [OpenWebUI](https://github.com/open-webui/open-webui) - The original project
- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [Logrus](https://github.com/sirupsen/logrus) - Structured logging

## 📞 Support

- **Issues**: [GitHub Issues](https://github.com/your-username/openwebui-go/issues)
- **Discussions**: [GitHub Discussions](https://github.com/your-username/openwebui-go/discussions)
- **Documentation**: [Wiki](https://github.com/your-username/openwebui-go/wiki)