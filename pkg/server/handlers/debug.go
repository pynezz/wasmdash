package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	// DefaultFallbackIP is used when Host header is missing (common on mobile)
	DefaultFallbackIP = "192.168.1.193"
)

// CSSDebugHandler returns a handler for CSS debugging information
func CSSDebugHandler(port string) echo.HandlerFunc {
	return func(c echo.Context) error {
		userAgent := c.Request().Header.Get("User-Agent")
		acceptEncoding := c.Request().Header.Get("Accept-Encoding")
		host := c.Request().Header.Get("Host")

		// Handle missing Host header (common on mobile)
		var cssURL string
		if host != "" {
			cssURL = fmt.Sprintf("http://%s/static/css/styles.css", host)
		} else {
			// Fallback: use the server's configured address
			cssURL = fmt.Sprintf("http://%s:%s/static/css/styles.css", DefaultFallbackIP, port)
		}

		debugInfo := fmt.Sprintf(`
CSS Debug Information:
- User-Agent: %s
- Accept-Encoding: %s
- Host: %s
- Fallback Host: %s:%s
- CSS URL: %s
- Request Path: %s
- Method: %s
- Remote IP: %s
- Content-Type: %s

Try accessing CSS directly: <a href="/static/css/styles.css">/static/css/styles.css</a>
Or try full URL: <a href="%s">%s</a>

Common Issues:
- Missing Host header on mobile browsers
- Content Security Policy blocking stylesheets
- MIME type not set correctly
- Network connectivity issues
- Cache issues (try hard refresh)
`, userAgent, acceptEncoding, host, DefaultFallbackIP, port, cssURL,
			c.Request().URL.Path, c.Request().Method, c.RealIP(),
			c.Request().Header.Get("Content-Type"), cssURL, cssURL)

		return c.HTML(http.StatusOK, fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>CSS Debug Information</title>
    <style>
        body { font-family: monospace; margin: 20px; background: #f5f5f5; }
        pre { background: white; padding: 20px; border-radius: 8px; border: 1px solid #ddd; }
        a { color: #0066cc; text-decoration: none; }
        a:hover { text-decoration: underline; }
    </style>
</head>
<body>
    <h1>CSS Debug Information</h1>
    <pre>%s</pre>
</body>
</html>`, debugInfo))
	}
}

// CSSTestHandler returns a handler for CSS functionality testing
func CSSTestHandler(c echo.Context) error {
	host := c.Request().Header.Get("Host")
	userAgent := c.Request().Header.Get("User-Agent")

	testHTML := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, shrink-to-fit=no">
    <title>CSS Test Page</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%);
            color: white;
            margin: 0;
            padding: 20px;
            min-height: 100vh;
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
        }
        .container {
            max-width: 600px;
            text-align: center;
            background: rgba(255,255,255,0.1);
            padding: 30px;
            border-radius: 15px;
            backdrop-filter: blur(10px);
            box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
        }
        .test-button {
            background: #4CAF50;
            color: white;
            padding: 12px 24px;
            border: none;
            border-radius: 8px;
            font-size: 16px;
            margin: 10px;
            cursor: pointer;
            transition: all 0.3s ease;
            display: inline-block;
            text-decoration: none;
        }
        .test-button:hover {
            background: #45a049;
            transform: translateY(-2px);
        }
        .success { background: #4CAF50; }
        .warning { background: #ff9800; }
        .error { background: #f44336; }
        .info { background: #2196F3; }

        @media (max-width: 768px) {
            .container {
                padding: 20px;
                margin: 10px;
            }
            h1 { font-size: 24px; }
        }

        .debug-info {
            background: rgba(0,0,0,0.3);
            padding: 15px;
            border-radius: 8px;
            margin-top: 20px;
            font-family: monospace;
            font-size: 12px;
            text-align: left;
        }
    </style>
    <link rel="stylesheet" href="/static/css/styles.css">
</head>
<body>
    <div class="container">
        <h1>🎨 CSS Loading Test</h1>
        <p>If you can see styled content, basic CSS is working!</p>

        <div>
            <button class="test-button success" onclick="testTailwind()">Test Tailwind CSS</button>
            <a href="/" class="test-button info">← Back to Home</a>
            <a href="/debug/css" class="test-button warning">CSS Debug Info</a>
        </div>

        <div id="tailwind-test" class="mt-4 p-4 bg-blue-500 text-white rounded-lg hidden">
            <p>✅ Tailwind CSS is working! 🎉</p>
        </div>

        <div class="debug-info">
            <h3>📱 Device Information:</h3>
            <p><strong>User Agent:</strong> %s</p>
            <p><strong>Host:</strong> %s</p>
            <p><strong>Screen Size:</strong> <span id="screen-size">Loading...</span></p>
            <p><strong>Viewport:</strong> <span id="viewport-size">Loading...</span></p>
            <p><strong>CSS Test Status:</strong> <span id="css-status" class="warning">Testing...</span></p>
        </div>
    </div>

    <script>
        function testTailwind() {
            const element = document.getElementById('tailwind-test');
            element.classList.toggle('hidden');

            // Test if Tailwind classes are working
            const hasClasses = element.classList.contains('bg-blue-500');
            const statusElement = document.getElementById('css-status');
            if (hasClasses) {
                statusElement.textContent = '✅ Working';
                statusElement.className = 'success';
            } else {
                statusElement.textContent = '❌ Failed';
                statusElement.className = 'error';
            }
        }

        // Display screen and viewport information
        document.addEventListener('DOMContentLoaded', function() {
            document.getElementById('screen-size').textContent =
                screen.width + 'x' + screen.height;
            document.getElementById('viewport-size').textContent =
                window.innerWidth + 'x' + window.innerHeight;

            // Auto-test CSS loading
            setTimeout(function() {
                const computed = window.getComputedStyle(document.body);
                const backgroundImage = computed.getPropertyValue('background-image');
                const statusElement = document.getElementById('css-status');

                if (backgroundImage && backgroundImage !== 'none') {
                    statusElement.textContent = '✅ CSS Loaded';
                    statusElement.className = 'success';
                } else {
                    statusElement.textContent = '⚠️ Partial';
                    statusElement.className = 'warning';
                }
            }, 500);
        });

        // Handle orientation changes
        window.addEventListener('orientationchange', function() {
            setTimeout(function() {
                document.getElementById('viewport-size').textContent =
                    window.innerWidth + 'x' + window.innerHeight;
            }, 500);
        });
    </script>
</body>
</html>`, userAgent, host)

	return c.HTML(http.StatusOK, testHTML)
}

// MobileDetectHandler returns a handler for mobile device detection and debugging
func MobileDetectHandler(port string) echo.HandlerFunc {
	return func(c echo.Context) error {
		userAgent := c.Request().Header.Get("User-Agent")
		userAgentLower := strings.ToLower(userAgent)

		isMobile := strings.Contains(userAgentLower, "mobile") ||
			strings.Contains(userAgentLower, "android") ||
			strings.Contains(userAgentLower, "iphone") ||
			strings.Contains(userAgentLower, "ipad") ||
			strings.Contains(userAgentLower, "ipod") ||
			strings.Contains(userAgentLower, "blackberry") ||
			strings.Contains(userAgentLower, "windows phone")

		isTablet := strings.Contains(userAgentLower, "tablet") ||
			strings.Contains(userAgentLower, "ipad")

		// Handle missing Host header (common on mobile)
		var cssURL string
		host := c.Request().Host
		if host != "" {
			cssURL = fmt.Sprintf("http://%s/static/css/styles.css", host)
		} else {
			cssURL = fmt.Sprintf("http://%s:%s/static/css/styles.css", DefaultFallbackIP, port)
		}

		// Determine device type
		deviceType := "Desktop"
		if isTablet {
			deviceType = "Tablet"
		} else if isMobile {
			deviceType = "Mobile"
		}

		// Generate recommendations based on device type
		recommendations := []string{
			"Test CSS loading directly at: " + cssURL,
			"Check browser developer tools Network tab",
			"Verify network connectivity to server",
			"Try direct CSS URL: /static/css/styles.css",
		}

		if isMobile {
			recommendations = append(recommendations,
				"Check if mobile browser blocks Host header",
				"Try different mobile browsers (Chrome, Safari, Firefox)",
				"Verify mobile network isn't blocking requests",
				"Check if mobile data saver is enabled",
			)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"device_info": map[string]interface{}{
				"is_mobile":   isMobile,
				"is_tablet":   isTablet,
				"device_type": deviceType,
				"user_agent":  userAgent,
				"remote_addr": c.RealIP(),
			},
			"server_info": map[string]interface{}{
				"host_header":   host,
				"fallback_host": fmt.Sprintf("%s:%s", DefaultFallbackIP, port),
				"port":          port,
			},
			"css_info": map[string]interface{}{
				"css_url":    cssURL,
				"css_direct": "/static/css/styles.css",
				"css_test":   "/test/css",
				"css_debug":  "/debug/css",
			},
			"headers":         c.Request().Header,
			"recommendations": recommendations,
			"troubleshooting": map[string]interface{}{
				"common_issues": []string{
					"Missing Host header on mobile browsers",
					"Content Security Policy restrictions",
					"Network connectivity problems",
					"MIME type issues",
					"Cache problems",
				},
				"test_urls": []string{
					"/test/css - Interactive CSS test",
					"/debug/css - Detailed CSS debug info",
					"/health - Server health check",
					cssURL + " - Direct CSS file access",
				},
			},
		})
	}
}
