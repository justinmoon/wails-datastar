package main

import (
	"bytes"
	"context"
	"strconv"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"ssr-poc/internal/views"
	"ssr-poc/pkg/datastar"
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

// IncHTML increments the count and returns HTML fragment and signal updates as IPC events
func (a *App) IncHTML() ([]byte, error) {
	a.count++

	// Render the updated count HTML
	var buf bytes.Buffer
	_ = views.Count(a.count).Render(context.Background(), &buf)

	// Create IPC builder
	ipc := datastar.NewIpc()

	// Add fragment update
	ipc.MergeFragments(buf.String())

	// Add signal update
	// ipc.MergeSignals(
	// 	[]byte(fmt.Sprintf(`{"count":%d}`, a.count)),
	// )

	// Return JSON-encoded IPC envelope
	jsonData, err := ipc.JSON()
	if err != nil {
		runtime.LogErrorf(a.ctx, "Error marshaling IPC JSON: %v", err)
		return nil, err
	}

	runtime.LogInfof(a.ctx, "IPC JSON response: %s", string(jsonData))
	return jsonData, nil
}
