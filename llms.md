# WasmDash: A Modern Web Application Framework

## Overview

WasmDash is a modern web application framework built with Go, WebAssembly, and cutting-edge frontend technologies. It provides a sleek, minimalistic dashboard with robust security features, mobile-friendly design, and excellent performance characteristics. The architecture has been recently refactored to support better maintainability, testability, and a cleaner separation of concerns.

## Technical Architecture

### Core Technologies

- **Backend**: Go with Echo web framework
- **Frontend**: Templ templating, TailwindCSS, WebAssembly
- **CSS**: TailwindCSS with modern features and responsive design
- **Rendering**: Server-side rendering with progressive enhancement
- **Security**: Content Security Policy (CSP) with nonce-based script protection
- **Server**: Modular architecture with configuration-driven design
- **Mobile**: Enhanced mobile compatibility with Host header fallback mechanisms

### Directory Structure

```
wasmdash/
├── main.go                     # Entry point with CLI interface
├── handlers.go                 # Legacy compatibility layer
├── Makefile                    # Build and development tasks
├── pkg/
│   ├── core/                   # Core functionality
│   ├── ui/                     # UI components
│   │   ├── components/         # Reusable UI components
│   │   │   ├── button/         # Button component
│   │   │   ├── card/           # Card component
│   │   │   ├── icon/           # Icon system
│   │   │   └── ...             # Other components
│   │   ├── pages/              # Page templates
│   │   │   ├── home.templ      # Homepage template
│   │   │   └── ...             # Other pages
│   │   ├── head.templ          # HTML head section
│   │   └── layout.templ        # Main layout template
│   └── server/                 # Server implementation
│       ├── server.go           # Server configuration
│       ├── handlers/           # HTTP handlers
│       │   ├── app.go          # Application handlers
│       │   └── debug.go        # Debug/testing handlers
│       └── middleware/         # Middleware components
│           └── static.go       # Static file handling
├── static/                     # Static assets
│   ├── css/                    # Compiled CSS
│   └── ...                     # Other static assets
└── assets/                     # Source assets
    └── css/                    # CSS source files
```

### Key Components

### Server Configuration (pkg/server/server.go)

The server is built using a modular, configuration-driven approach with a clean separation of concerns:

```go
// Config holds server configuration
type Config struct {
    Port        string
    Host        string
    Environment string
    ServerName  string
}

// Server wraps the Echo instance with configuration
type Server struct {
    echo   *echo.Echo
    config *Config
}
```

This design separates concerns and allows for environment-specific configurations (development vs. production). The refactored architecture provides clear methods for server lifecycle management:

```go
// Server methods
func New(config *Config) *Server
func (s *Server) SetupMiddleware()
func (s *Server) SetupRoutes()
func (s *Server) Start() error
func (s *Server) Shutdown() error
```

#### UI System (pkg/ui)

The UI system uses the Templ templating engine, which combines Go's type safety with reactive components:

```go
// Layout defines the page structure
templ Layout(content templ.Component, nonce string, path ...string) {
    <!DOCTYPE html>
    <html lang="en">
        @Head("wDash", path...)
        <body>
            <main class="flex-grow">
                @content
            </main>
        </body>
    </html>
}
```

#### Component Architecture

Components follow a consistent props-based pattern:

```go
// Props for a Button component
type Props struct {
    ID           string
    Class        string
    Variant      Variant
    Size         Size
    // ...other properties
}

// Button component with props
templ Button(props ...Props) {
    // Component implementation
}
```

### Middleware System (pkg/server/middleware)

Middleware is organized into specialized modules with enhanced functionality:

```go
// StaticFileHeaders middleware sets proper MIME types and cache headers
func StaticFileHeaders() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            if strings.HasPrefix(c.Request().URL.Path, "/static/") {
                // Set proper MIME types based on file extension
                ext := filepath.Ext(c.Request().URL.Path)
                switch ext {
                case ".css":
                    c.Response().Header().Set("Content-Type", "text/css; charset=utf-8")
                // ... other MIME types
                }

                // Set cache headers for static assets
                c.Response().Header().Set("Cache-Control", "public, max-age=31536000")
                c.Response().Header().Set("Vary", "Accept-Encoding")
            }
            return next(c)
        }
    }
}
```

The middleware system now includes comprehensive type handling for various file formats and security-focused headers.

### Security Implementation

WasmDash implements strong security measures:

1. **Content Security Policy**: Strict CSP with nonce-based script execution
2. **MIME Type Enforcement**: Proper MIME types for all static resources
3. **Header Security**: Security headers for XSS protection
4. **Secure Defaults**: Production-ready security defaults

### Mobile Compatibility

The framework addresses mobile-specific concerns with enhanced debugging capabilities:

1. **Responsive Design**: TailwindCSS-based responsive layouts
2. **Host Header Handling**: Fallback mechanisms for missing Host headers on mobile devices
3. **Debug Tools**: Comprehensive mobile-specific debugging endpoints
4. **Meta Tags**: Mobile-optimized viewport and format detection
5. **CSS Troubleshooting**: Special endpoints for diagnosing mobile CSS issues
6. **Device Detection**: Intelligent mobile device type detection with tailored recommendations

### Mobile Optimization

The framework includes special handling for mobile devices:

```go
// Handle missing Host header (common on mobile)
var cssURL string
if host != "" {
    cssURL = fmt.Sprintf("http://%s/static/css/styles.css", host)
} else {
    // Fallback to configured IP and port
    cssURL = fmt.Sprintf("http://%s:%s/static/css/styles.css", DefaultFallbackIP, port)
}
// Where DefaultFallbackIP = "192.168.1.193"
```

**Note**: The `DefaultFallbackIP` constant is defined in `pkg/server/handlers/debug.go` and should be updated to match your server's actual IP address when deploying to a different environment. This is particularly important for mobile device testing as some mobile browsers may not send the Host header.

## API Reference

### CLI Interface

The server provides a clean CLI interface with sensible defaults:

```bash
wasmdash [flags] [command]

Commands:
  serve    Start the web server (default)

Flags:
  --port   Port to serve on (default: 8080)
  --host   Host to bind to (default: all interfaces)
  --env    Environment: development, production (default: development)
  --help   Show this help message

Examples:
  wasmdash                                  # Default serve
  wasmdash serve                            # Explicit serve command
  wasmdash --port 3000                      # Custom port
  wasmdash --host 192.168.1.193 --port 8081 # Specific host and port (use your server's IP)
  wasmdash --env production                 # Production environment
```

### HTTP Endpoints

- **Main Application**:
  - **`/`**: Main application homepage
  - **`/about`**: About page
  - **`/health`**: Health check endpoint
  - **`/404`**: Custom not found page
  - **`/robots.txt`**: Robots instructions
  - **`/service-worker.js`**: PWA service worker

- **Development Tools** (only available in development environment):
  - **`/debug/css`**: Comprehensive CSS debugging information
  - **`/test/css`**: Interactive CSS testing page
  - **`/mobile/detect`**: Mobile device detection and troubleshooting

### Component API

Components follow a consistent props-based pattern:

```go
// Using a Button component
@button.Button(button.Props{
    Variant: button.VariantPrimary,
    Size:    button.SizeMedium,
    Class:   "custom-class",
}) {
    Button Text
}
```

## Implementation Details

### Rendering Pipeline

1. HTTP request received by Echo server
2. Route matched to handler function
3. Handler prepares data and selects template
4. Templ renders component tree
5. CSP nonce generated and injected
6. HTML response sent to client

### CSS Processing

1. TailwindCSS processes `assets/css/base.css`
2. Output minified to `static/css/styles.css`
3. CSS loaded with preload hints for performance
4. Mobile-specific optimizations applied

### Mobile Optimization

The framework includes special handling for mobile devices:

```go
// Handle missing Host header (common on mobile)
var cssURL string
if host != "" {
    cssURL = fmt.Sprintf("http://%s/static/css/styles.css", host)
} else {
    // Fallback: use the server's configured address
    cssURL = fmt.Sprintf("http://%s:%s/static/css/styles.css", DefaultFallbackIP, port)
}
```

### Graceful Shutdown

The server now implements proper lifecycle management with graceful shutdown:

```go
// In main.go
func runServer(port, host, environment string) error {
    // Create and configure server
    config := &server.Config{
        Port:        port,
        Host:        host,
        Environment: environment,
        ServerName:  "wasmdash",
    }
    srv := server.New(config)
    srv.SetupMiddleware()
    srv.SetupRoutes()

    // Setup graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    // Start server in a goroutine
    go func() {
        log.Printf("Starting WasmDash server on %s:%s (environment: %s)",
                   host, port, environment)
        if err := srv.Start(); err != nil {
            log.Printf("Server error: %v", err)
        }
    }()

    // Wait for interrupt signal
    <-quit
    log.Println("Shutting down server...")

    // Perform graceful shutdown
    if err := srv.Shutdown(); err != nil {
        return fmt.Errorf("server forced to shutdown: %w", err)
    }

    log.Println("Server exited")
    return nil
}
```

## Development Workflow

### Building and Running

```bash
# Build CSS and Go binary
make build

# Run development server with hot reload
make dev

# Generate templ files
make gen

# Clean build artifacts
make clean
```

### Adding New Features

1. Create component in `pkg/ui/components/` directory
2. Add handler in `pkg/server/handlers/` directory
3. Register route in `pkg/server/server.go`
4. Update CSS in `assets/css/base.css` if needed
5. Generate and build with `make build`

## Troubleshooting

Common issues and solutions:

- **CSS not loading on mobile**:
   - Check Host header handling and MIME types
   - Use `/debug/css` endpoint to diagnose
   - Check if mobile browser is missing Host header
   - Try accessing CSS directly: `/static/css/styles.css`
   - Verify fallback IP (currently `192.168.1.193`) matches your network setup
   - Modify `DefaultFallbackIP` in `pkg/server/handlers/debug.go` if needed

2. **MIME type errors**:
   - Verify proper Content-Type headers in static file middleware
   - Ensure MIME types are set correctly for all file types

3. **CSP violations**:
   - Ensure scripts have proper nonce attributes
   - Check CSP headers in browser developer tools
   - Verify style-src includes 'unsafe-inline' for Tailwind compatibility

4. **Mobile-specific issues**:
   - Test with the `/mobile/detect` endpoint for detailed diagnostics
   - Use `/test/css` for interactive CSS testing on mobile devices
   - Check if Host header is missing on your mobile device
   - Verify network connectivity between mobile device and server

5. **Build errors**:
   - Run `make clean` and then `make build`
   - Check for any middleware or handler errors

## References

- Go: https://golang.org
- Echo: https://echo.labstack.com
- Templ: https://github.com/a-h/templ
- TailwindCSS: https://tailwindcss.com
- WebAssembly: https://webassembly.org
