# GitHub Actions Workflow

This repository includes a comprehensive GitHub Actions workflow for building, testing, and publishing the OpenWebUI Go application.

## Workflow Overview

The workflow (`build-and-publish.yml`) consists of four main jobs:

### 1. Test Job
- Runs on every push and pull request
- Sets up Go environment
- Caches Go modules for faster builds
- Runs unit tests with race detection
- Generates code coverage reports
- Uploads coverage to Codecov (optional)

### 2. Build Job
- Builds the application for multiple platforms:
  - Linux (amd64, arm64)
  - Windows (amd64)
  - macOS (amd64, arm64)
- Uses cross-compilation for efficient builds
- Embeds frontend assets using the prepare_frontend.sh script
- Uploads build artifacts for each platform
- Artifacts are retained for 30 days

### 3. Docker Job
- Builds and pushes Docker images to GitHub Container Registry
- Only runs on pushes to main/master branches or tags
- Supports multi-platform builds (linux/amd64, linux/arm64)
- Uses Docker layer caching for efficiency
- Tags images appropriately based on branch/tag

### 4. Release Job
- Only runs when tags are pushed (e.g., `v1.0.0`)
- Downloads all build artifacts
- Creates compressed archives for each platform
- Creates a GitHub release with all binaries
- Auto-generates release notes

## Triggering the Workflow

### Automatic Triggers
- **Push to main/master**: Runs test, build, and docker jobs
- **Pull requests**: Runs test and build jobs
- **Tag push** (v*): Runs all jobs including release creation

### Manual Triggers
The workflow can be manually triggered from the Actions tab in your GitHub repository.

## Artifacts and Outputs

### Build Artifacts
- **Linux**: `openwebui-go-linux-amd64.tar.gz`, `openwebui-go-linux-arm64.tar.gz`
- **Windows**: `openwebui-go-windows-amd64.zip`
- **macOS**: `openwebui-go-darwin-amd64.tar.gz`, `openwebui-go-darwin-arm64.tar.gz`

### Docker Images
- Registry: `ghcr.io/[username]/[repository]`
- Tags: branch names, semantic versions, commit SHAs
- Multi-architecture support (amd64, arm64)

## Setup Requirements

### System Requirements
- **Node.js 20+**: Required for building the OpenWebUI frontend components
- **Go 1.21**: Specified in the workflow for building the backend

### Repository Secrets
No additional secrets are required. The workflow uses the built-in `GITHUB_TOKEN` for:
- Publishing Docker images to GitHub Container Registry
- Creating releases and uploading assets

### Optional: Codecov Integration
Code coverage is generated but not uploaded by default. To enable Codecov:
1. Sign up at https://codecov.io
2. Add your repository
3. Uncomment the codecov upload step in the workflow
4. Add CODECOV_TOKEN secret if your repository is private

## Customization

### Changing Build Targets
Modify the matrix strategy in the build job to add/remove platforms:

```yaml
strategy:
  matrix:
    goos: [linux, windows, darwin]
    goarch: [amd64, arm64]
```

### Docker Registry
To use a different container registry, update the docker job:
- Change the registry URL
- Update login credentials
- Modify image names

### Release Assets
Customize the release job to include additional files or change compression formats.

## Local Development

To test the build process locally:

### Prerequisites
- Go 1.21+
- Node.js 20+ (required for frontend preparation)
- Make

### Commands
```bash
# Prepare frontend assets
make prepare-frontend

# Build for current platform
make build

# Build for all platforms
make build-all

# Test the application
go test ./...

# Build Docker image
docker build -t openwebui-go .
```