package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pynezz/wasmdash/pkg/server"
)

//go:generate tailwindcss -i assets/css/base.css -o static/css/styles.css -m

//go:embed static/*
var static embed.FS
var commit string = ""
var buildTime string = ""

// Wasmdash represents the main application structure
type Wasmdash struct {
	Build   string
	Version string
	Config  *WConfig
}

// WConfig holds application configuration
type WConfig struct {
	Port string
	Host string
	Env  string
}

func main() {
	// Initialize application
	app := &Wasmdash{
		Build:   commit,
		Version: buildTime,
		Config:  &WConfig{},
	}

	// Set up command line flags
	helpFlag := flag.Bool("help", false, "Show this help message")
	portFlag := flag.String("port", "8080", "Port to listen on")
	hostFlag := flag.String("host", "localhost", "Host to listen on")
	envFlag := flag.String("env", "development", "Environment to run in (development, production)")

	// Parse the command line flags
	flag.Parse()

	// Check if help was requested
	if *helpFlag {
		showHelp()
		return
	}

	// Apply parsed flags to configuration
	app.Config.Port = *portFlag
	app.Config.Host = *hostFlag
	app.Config.Env = *envFlag

	fmt.Printf("\033[34mDatdash v%s - %s running on %s\n:\033[36m%s/%s\n\033[0m", app.Version, app.Build, app.Config.Host, app.Config.Port, app.Config.Env)

	// Configure logging
	os.Setenv("WLOGPATH", "wasmdash.log")
	log.Printf("Starting WasmDash (Version: %s, Build: %s)", app.Version, app.Build)
	log.Printf("Configuration: Host=%s, Port=%s, Environment=%s", app.Config.Host, app.Config.Port, app.Config.Env)

	// Run the server
	if err := runServer(app.Config); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func runServer(config *WConfig) error {
	// Create server configuration
	serverConfig := &server.Config{
		Port:        config.Port,
		Host:        config.Host,
		Environment: config.Env,
		ServerName:  "wasmdash",
	}

	// Create new server instance
	srv := server.New(serverConfig)

	// Setup middleware and routes
	srv.SetupMiddleware()
	srv.SetupRoutes()

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a goroutine
	go func() {
		log.Printf("Server listening on %s:%s (environment: %s)", config.Host, config.Port, config.Env)
		if err := srv.Start(); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-quit
	log.Println("Shutting down server...")

	if err := srv.Shutdown(); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	log.Println("Server exited")
	return nil
}

func showHelp() {
	helpMsg := `WasmDash - A modern dashboard built with Go and WASM

Usage:
    wasmdash [flags]

Flags:
    --port PORT      Port to serve on (default: 8080)
    --host HOST      Host to bind to (default: localhost)
    --env ENV        Environment: development, production (default: development)
    --help           Show this help message

Examples:
    wasmdash
    wasmdash --port 3000
    wasmdash --host 192.168.1.193 --port 8081
    wasmdash --env production --port 80
    wasmdash --help`

	fmt.Println(helpMsg)
}
