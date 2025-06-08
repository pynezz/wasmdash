package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/pynezz/wasmdash/pkg/server/handlers"
	"github.com/pynezz/wasmdash/pkg/server/middleware"
)

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

// New creates a new server instance
func New(config *Config) *Server {
	if config == nil {
		config = &Config{
			Port:        "8080",
			Host:        "localhost",
			Environment: "development",
			ServerName:  "wasmdash",
		}
	}

	e := echo.New()
	e.HideBanner = true

	return &Server{
		echo:   e,
		config: config,
	}
}

// SetupMiddleware configures all middleware
func (s *Server) SetupMiddleware() {
	// Server header middleware
	s.echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderServer, s.config.ServerName+":"+s.config.Port)
			return next(c)
		}
	})

	// Static files middleware
	s.echo.Static("/static", "static")

	// MIME type and cache headers middleware
	s.echo.Use(middleware.StaticFileHeaders())
}

// SetupRoutes configures all application routes
func (s *Server) SetupRoutes() {
	// Main application routes
	s.echo.GET("/", handlers.HomeHandler)
	s.echo.GET("/about", handlers.AboutHandler)
	s.echo.GET("/service-worker.js", handlers.ServiceWorkerHandler)

	// Utility routes
	s.echo.GET("/robots.txt", handlers.RobotsHandler)
	s.echo.GET("/404", handlers.NotFoundHandler)
	s.echo.GET("/health", handlers.HealthHandler)

	// Debug routes (only in development)
	if s.config.Environment == "development" {
		s.setupDebugRoutes()
	}
}

// setupDebugRoutes configures debug and testing routes
func (s *Server) setupDebugRoutes() {
	debugGroup := s.echo.Group("/debug")
	debugGroup.GET("/css", handlers.CSSDebugHandler(s.config.Port))

	testGroup := s.echo.Group("/test")
	testGroup.GET("/css", handlers.CSSTestHandler)

	mobileGroup := s.echo.Group("/mobile")
	mobileGroup.GET("/detect", handlers.MobileDetectHandler(s.config.Port))
}

func (s *Server) ContentSecurityPolicy(next echo.HandlerFunc) string {
	csp := "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data:; font-src 'self' data:; connect-src 'self'; frame-src 'self'; object-src 'self'; media-src 'self'; child-src 'self';"
	return csp
}

// Start starts the HTTP server
func (s *Server) Start() error {
	log.Printf("Starting server on %s:%s", s.config.Host, s.config.Port)

	address := ":" + s.config.Port
	if s.config.Host != "" && s.config.Host != "localhost" {
		address = s.config.Host + ":" + s.config.Port
	}

	return s.echo.Start(address)
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown() error {
	log.Println("Shutting down server...")
	return s.echo.Close()
}

// Echo returns the underlying Echo instance for advanced configuration
func (s *Server) Echo() *echo.Echo {
	return s.echo
}

// GetConfig returns the server configuration
func (s *Server) GetConfig() *Config {
	return s.config
}
