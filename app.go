package main

import (
	"bytes"
	"context"
	"strconv"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"ssr-poc/internal/views"
)

// App struct
type App struct {
	ctx   context.Context
	count int
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		count: 0,
	}
}

// startup is called when the app starts
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// Log status
	runtime.LogInfo(a.ctx, "Application started")
}

// GetHTML returns the rendered HTML from our templ components
func (a *App) GetHTML() string {
	var buf bytes.Buffer
	_ = views.Index(a.count).Render(context.Background(), &buf)
	return buf.String()
}

// Inc increments the counter and returns the updated value
func (a *App) Increment() int {
	a.count++
	return a.count
}

// GetCount returns the current count
func (a *App) GetCount() string {
	return strconv.Itoa(a.count)
}
