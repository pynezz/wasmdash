package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/pynezz/wasmdash/pkg/server"
)

//go:embed static/*
var static embed.FS
var commit string = ""
var buildTime string = ""

var wasmdash Wasmdash

type Wasmdash struct {
	Build   string
	Version string
	Flags   *flag.FlagSet

	Config *WConfig
}

type WConfig struct {
	Port string
	Host string
	Env  string
}

func NewWConfig() *WConfig {
	return &WConfig{
		Port: "8080",
		Host: "0.0.0.0",
		Env:  "development",
	}
}

func (w *Wasmdash) Run() error {
	log.Println("initializing wasmdash application")
	os.Setenv("WLOGPATH", "wasmdash.log")

	wasmdash = Wasmdash{
		Build:   commit,
		Version: "June 2025",
		Flags:   flag.NewFlagSet("wasmdash", flag.ExitOnError),
	}

	return nil
}

//go:generate tailwindcss -i assets/css/base.css -o static/css/styles.css -m

func init() {
	log.Println("initializing wasmdash application")
	os.Setenv("WLOGPATH", "wasmdash.log")

	wasmdash = Wasmdash{
		Build:   commit,
		Version: time.Now().Format("January 2006"),
		Flags:   flag.NewFlagSet("wasmdash", flag.ExitOnError),
	}
}

func args() []string {
	return flag.Args()
}

func main() {

	if *help {
		showHelp()
		return
	}

	args := flag.Args()
	command := "serve" // Default command
	if len(args) > 0 {
		command = args[0]
	}

	switch command {
	case "serve":
		if err := runServer(*port, *host, *env); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
		showHelp()
		os.Exit(1)
	}
}

func showHelp() {
	helpMsg := `#WasmDash - A modern dashboard built with Go and WASM

		#Usage:
			wasmdash [flags] [command]

		#Commands:
			serve    Start the web server (default)

		#Flags:
			--port   Port to serve on (default: 8080)
			--host   Host to bind to (default: all interfaces)
			--env    Environment: development, production (default: development)
			--help   Show this help message

		#Examples:
			wasmdash
			wasmdash serve
			wasmdash --port 3000 serve
			wasmdash --host 192.168.1.193 --port 8081 serve
			wasmdash --env production --port 80 serve
			wasmdash --help`
	for _, line := range strings.Split(helpMsg, "\n") {
		line = strings.Trim(line, "	")
		if strings.HasPrefix(line, "#") {
			line = strings.TrimPrefix(line, "#")
			fmt.Println(line)
			continue
		}

		fmt.Println("\t", line)
	}
}

func runServer(port, host, environment string) error {
	// Create server configuration
	config := &server.Config{
		Port:        port,
		Host:        host,
		Environment: environment,
		ServerName:  "wasmdash",
	}

	// Create new server instance
	srv := server.New(config)

	// Setup middleware and routes
	srv.SetupMiddleware()
	srv.SetupRoutes()

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Starting WasmDash server on %s:%s (environment: %s)", host, port, environment)
		if err := srv.Start(); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()
	<-quit
	log.Println("Shutting down server...")

	if err := srv.Shutdown(); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	log.Println("Server exited")
	return nil
}
