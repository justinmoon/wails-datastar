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
	_ = views.Index("World", a.count).Render(context.Background(), &buf)
	return buf.String()
}

// Inc increments the counter and returns the updated HTML fragment
func (a *App) Inc() string {
	a.count++
	
	// Render Count component to string
	var buf bytes.Buffer
	_ = views.Count(a.count).Render(context.Background(), &buf)
	
	return buf.String()
}

// GetCount returns the current count
func (a *App) GetCount() string {
	return strconv.Itoa(a.count)
}

// IncHTML increments the counter and returns the rendered HTML fragment
func (a *App) IncHTML() (string, error) {
	runtime.LogInfo(a.ctx, "IncHTML method called")
	a.count++
	var buf bytes.Buffer
	_ = views.Count(a.count).Render(context.Background(), &buf)
	result := buf.String()
	runtime.LogInfo(a.ctx, "IncHTML returning: " + result)
	return result, nil
}