package middleware

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

// StaticFileHeaders returns middleware that sets proper MIME types and cache headers for static files
func StaticFileHeaders() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if strings.HasPrefix(c.Request().URL.Path, "/static/") {
				// Set proper MIME types
				ext := filepath.Ext(c.Request().URL.Path)
				switch ext {
				case ".css":
					c.Response().Header().Set("Content-Type", "text/css; charset=utf-8")
				case ".js":
					c.Response().Header().Set("Content-Type", "application/javascript; charset=utf-8")
				case ".woff2":
					c.Response().Header().Set("Content-Type", "font/woff2")
				case ".woff":
					c.Response().Header().Set("Content-Type", "font/woff")
				case ".ttf":
					c.Response().Header().Set("Content-Type", "font/ttf")
				case ".otf":
					c.Response().Header().Set("Content-Type", "font/otf")
				case ".ico":
					c.Response().Header().Set("Content-Type", "image/x-icon")
				case ".png":
					c.Response().Header().Set("Content-Type", "image/png")
				case ".jpg", ".jpeg":
					c.Response().Header().Set("Content-Type", "image/jpeg")
				case ".gif":
					c.Response().Header().Set("Content-Type", "image/gif")
				case ".svg":
					c.Response().Header().Set("Content-Type", "image/svg+xml")
				case ".webp":
					c.Response().Header().Set("Content-Type", "image/webp")
				case ".json":
					c.Response().Header().Set("Content-Type", "application/json")
				case ".xml":
					c.Response().Header().Set("Content-Type", "application/xml")
				case ".pdf":
					c.Response().Header().Set("Content-Type", "application/pdf")
				case ".txt":
					c.Response().Header().Set("Content-Type", "text/plain; charset=utf-8")
				case ".html":
					c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
				}

				// Set cache headers for static assets
				c.Response().Header().Set("Cache-Control", "public, max-age=31536000")
				c.Response().Header().Set("Vary", "Accept-Encoding")

				// Add security headers for static files
				c.Response().Header().Set("X-Content-Type-Options", "nosniff")
				
				// Add CORS headers if needed
				if isFont(ext) {
					c.Response().Header().Set("Access-Control-Allow-Origin", "*")
					c.Response().Header().Set("Access-Control-Allow-Methods", "GET")
				}
			}
			return next(c)
		}
	}
}

// isFont checks if the file extension is a font file
func isFont(ext string) bool {
	fontExtensions := []string{".woff", ".woff2", ".ttf", ".otf", ".eot"}
	for _, fontExt := range fontExtensions {
		if ext == fontExt {
			return true
		}
	}
	return false
}

// StaticFileConfig holds configuration for static file middleware
type StaticFileConfig struct {
	CacheMaxAge int    // Cache max age in seconds
	Prefix      string // URL prefix for static files
}

// StaticFileHeadersWithConfig returns middleware with custom configuration
func StaticFileHeadersWithConfig(config StaticFileConfig) echo.MiddlewareFunc {
	if config.CacheMaxAge == 0 {
		config.CacheMaxAge = 31536000 // 1 year default
	}
	if config.Prefix == "" {
		config.Prefix = "/static/"
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if strings.HasPrefix(c.Request().URL.Path, config.Prefix) {
				// Set proper MIME types
				ext := filepath.Ext(c.Request().URL.Path)
				setMimeType(c, ext)

				// Set cache headers for static assets
				cacheControl := fmt.Sprintf("public, max-age=%d", config.CacheMaxAge)
				c.Response().Header().Set("Cache-Control", cacheControl)
				c.Response().Header().Set("Vary", "Accept-Encoding")

				// Add security headers for static files
				c.Response().Header().Set("X-Content-Type-Options", "nosniff")
				
				// Add CORS headers for fonts
				if isFont(ext) {
					c.Response().Header().Set("Access-Control-Allow-Origin", "*")
					c.Response().Header().Set("Access-Control-Allow-Methods", "GET")
				}
			}
			return next(c)
		}
	}
}

// setMimeType sets the appropriate MIME type based on file extension
func setMimeType(c echo.Context, ext string) {
	switch ext {
	case ".css":
		c.Response().Header().Set("Content-Type", "text/css; charset=utf-8")
	case ".js":
		c.Response().Header().Set("Content-Type", "application/javascript; charset=utf-8")
	case ".woff2":
		c.Response().Header().Set("Content-Type", "font/woff2")
	case ".woff":
		c.Response().Header().Set("Content-Type", "font/woff")
	case ".ttf":
		c.Response().Header().Set("Content-Type", "font/ttf")
	case ".otf":
		c.Response().Header().Set("Content-Type", "font/otf")
	case ".ico":
		c.Response().Header().Set("Content-Type", "image/x-icon")
	case ".png":
		c.Response().Header().Set("Content-Type", "image/png")
	case ".jpg", ".jpeg":
		c.Response().Header().Set("Content-Type", "image/jpeg")
	case ".gif":
		c.Response().Header().Set("Content-Type", "image/gif")
	case ".svg":
		c.Response().Header().Set("Content-Type", "image/svg+xml")
	case ".webp":
		c.Response().Header().Set("Content-Type", "image/webp")
	case ".json":
		c.Response().Header().Set("Content-Type", "application/json")
	case ".xml":
		c.Response().Header().Set("Content-Type", "application/xml")
	case ".pdf":
		c.Response().Header().Set("Content-Type", "application/pdf")
	case ".txt":
		c.Response().Header().Set("Content-Type", "text/plain; charset=utf-8")
	case ".html":
		c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	}
}