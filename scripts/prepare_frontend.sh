#!/bin/bash

set -e

echo "🚀 Preparing OpenWebUI frontend for Go embedding..."

# Configuration
OPENWEBUI_REPO="https://github.com/open-webui/open-webui.git"
TEMP_DIR="temp_openwebui"
FRONTEND_TARGET_DIR="assets/frontend"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if git is available
if ! command -v git &> /dev/null; then
    print_error "Git is not installed. Please install git first."
    exit 1
fi

# Check if node/npm is available
if ! command -v npm &> /dev/null; then
    print_warning "npm is not installed. You'll need to build the frontend manually."
    print_warning "Please install Node.js 20+ and npm, then run: cd $TEMP_DIR && npm install && npm run build"
elif command -v node &> /dev/null; then
    NODE_VERSION=$(node --version | sed 's/v//')
    NODE_MAJOR=$(echo $NODE_VERSION | cut -d. -f1)
    if [ "$NODE_MAJOR" -lt 20 ]; then
        print_warning "Node.js version $NODE_VERSION detected. OpenWebUI requires Node.js 20+."
        print_warning "Please upgrade Node.js to version 20 or higher."
    fi
fi

# Clean up any existing temp directory
if [ -d "$TEMP_DIR" ]; then
    print_status "Cleaning up existing temp directory..."
    rm -rf "$TEMP_DIR"
fi

# Clone OpenWebUI repository
print_status "Cloning OpenWebUI repository..."
git clone --depth 1 "$OPENWEBUI_REPO" "$TEMP_DIR"

# Check if this is a SvelteKit application
if [ -f "$TEMP_DIR/package.json" ] && [ -d "$TEMP_DIR/src" ]; then
    print_status "Detected SvelteKit application structure"
    
    # Create target directory
    print_status "Creating target directory..."
    mkdir -p "$FRONTEND_TARGET_DIR"
    
    # Check if we can build the frontend
    if command -v npm &> /dev/null; then
        print_status "Building SvelteKit frontend..."
        cd "$TEMP_DIR"
        
        # Install dependencies
        print_status "Installing npm dependencies..."
        npm install
        
        # Build the frontend
        print_status "Building SvelteKit application..."
        NODE_OPTIONS="--max-old-space-size=4096" npm run build
        
        # Check if build was successful
        if [ -d "build" ]; then
            print_status "Moving build output..."
            # Copy build contents to target directory
            cp -r build/* "../$FRONTEND_TARGET_DIR/"
            cd - > /dev/null
        else
            print_error "Build failed - no build directory found"
            exit 1
        fi
    else
        print_warning "npm not available. Using source files as-is."
        print_warning "You'll need to build the frontend manually:"
        print_warning "  cd $TEMP_DIR"
        print_warning "  npm install"
        print_warning "  npm run build"
        print_warning "  cp -r build/* ../$FRONTEND_TARGET_DIR/"
    fi
else
    print_error "Unexpected OpenWebUI structure. Expected SvelteKit application."
    exit 1
fi

# Clean up temp directory
print_status "Cleaning up..."
rm -rf "$TEMP_DIR"

# Verify we have the essential files
if [ ! -f "$FRONTEND_TARGET_DIR/index.html" ]; then
    print_error "index.html not found in build output"
    exit 1
fi

print_status "Frontend preparation completed!"
print_status "You can now build the Go application with: make build" 