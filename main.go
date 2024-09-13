package main

import (
	"context"
	"fmt"
	"morethancoder/goat/components"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var (
	// Define App Routes & Public Directory
	routes    = []string{"/", "/examples", "/contact"}
	publicDir = "public"

	// Logger & Styles
	gopherStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00FFFF"))
	goatStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFF00"))

	logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          goatStyle.Render("Maa 🐐 "),
	})
	styles = log.DefaultStyles()
)

func main() {
	// Fancy logging
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().Bold(true).
		SetString("INFO").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("#00FFFF")).
		Foreground(lipgloss.Color("0"))
	logger.SetStyles(styles)

	// Checking public directory existence
	if err := os.MkdirAll(publicDir, 0755); err != nil {
		logger.Fatal("Failed to create public directory", "error", err)
	}

	// Clear the terminal
	fmt.Print("\033[H\033[2J")

	// Generate HTML for each route
	for _, route := range routes {
		if err := generateHTML(route); err != nil {
			logger.Fatal("Failed to generate HTML", "route", route, "error", err)
		}
	}

	// Clean up obsolete files
	if err := cleanUp(); err != nil {
		logger.Fatal("Failed to clean up obsolete files", "error", err)
	}
}

func generateHTML(route string) error {
	filename := "index.html"
	if route != "/" {
		filename = route[1:] + ".html"
	}
	filePath := filepath.Join(publicDir, filename)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer file.Close()

	if err := components.App(route).Render(context.Background(), file); err != nil {
		return fmt.Errorf("failed to render HTML for route %s: %w", route, err)
	}

	logger.Infof("Successfully generated file %s", gopherStyle.Render(filePath))
	return nil
}

func routeExists(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func cleanUp() error {
	files, err := os.ReadDir(publicDir)
	if err != nil {
		return fmt.Errorf("error reading directory: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}
		filename := file.Name()
		route := "/" + strings.TrimSuffix(filename, ".html")
		if filename == "index.html" {
			route = "/"
		}
		if !routeExists(routes, route) {
			filePath := filepath.Join(publicDir, filename)
			if err := os.Remove(filePath); err != nil {
				logger.Error("Failed to remove file", "file", filePath, "error", err)
			} else {
				logger.Info("Removed obsolete file", "file", filePath)
			}
		}
	}
	return nil
}
