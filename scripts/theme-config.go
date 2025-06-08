package main

// Dashboard theme config

import (
	"fmt"
	"log"
	"os"
)

type ThemeConfig struct {
	Primary string
	Accent  string
	Success string
	Warning string
	Dark    bool
}

func main() {
	config := ThemeConfig{
		Primary: "262 83% 58%", // Purple
		Accent:  "178 60% 48%", // Teal
		Success: "142 76% 36%", // Green
		Warning: "43 96% 56%",  // Golden
		Dark:    true,
	}

	generateThemeCSS(config)
}

func generateThemeCSS(config ThemeConfig) {
	css := fmt.Sprintf(`
/* Auto-generated Dashboard Theme */
:root {
    --dashboard-primary: hsl(%s);
    --dashboard-accent: hsl(%s);
    --dashboard-success: hsl(%s);
    --dashboard-warning: hsl(%s);
}

.dark {
    --dashboard-primary: hsl(%s);
    --dashboard-accent: hsl(%s);
    --dashboard-success: hsl(%s);
    --dashboard-warning: hsl(%s);
}
`,
		config.Primary, config.Accent, config.Success, config.Warning,
		config.Primary, config.Accent, config.Success, config.Warning,
	)

	if err := os.WriteFile("static/css/dashboard-theme.css", []byte(css), 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\033[32m[+] Dashboard theme generated\033[0m")
}
