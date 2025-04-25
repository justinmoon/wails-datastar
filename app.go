package main

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"ssr-poc/internal/views"
	"ssr-poc/pkg/datastar"
)

// App struct
type App struct {
	ctx           context.Context
	fragmentCount int // existing counter using fragment updates
	signalCount   int // new counter using signal updates
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		fragmentCount: 0,
		signalCount:   0,
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
	_ = views.Index(a.fragmentCount, a.signalCount).Render(context.Background(), &buf)
	return buf.String()
}

// Inc increments the counter and returns the updated value
func (a *App) Increment() int {
	a.fragmentCount++
	return a.fragmentCount
}

// GetCount returns the current count
func (a *App) GetCount() string {
	return strconv.Itoa(a.fragmentCount)
}

// increments the count and returns HTML fragment
func (a *App) IncrementFragments() ([]byte, error) {
	a.fragmentCount++

	// Render the updated count HTML
	var buf bytes.Buffer
	_ = views.Count(a.fragmentCount).Render(context.Background(), &buf)

	// Create IPC builder
	ipc := datastar.NewIpc()

	// Add fragment update
	ipc.MergeFragments(buf.String())

	// Return JSON-encoded IPC envelope
	jsonData, err := ipc.JSON()
	if err != nil {
		runtime.LogErrorf(a.ctx, "Error marshaling IPC JSON: %v", err)
		return nil, err
	}

	runtime.LogInfof(a.ctx, "IPC JSON response: %s", string(jsonData))
	return jsonData, nil
}

// decrements the count and returns HTML fragment
func (a *App) DecrementFragments() ([]byte, error) {
	a.fragmentCount--

	// Render the updated count HTML
	var buf bytes.Buffer
	_ = views.Count(a.fragmentCount).Render(context.Background(), &buf)

	// Create IPC builder
	ipc := datastar.NewIpc()

	// Add fragment update
	ipc.MergeFragments(buf.String())

	// Return JSON-encoded IPC envelope
	jsonData, err := ipc.JSON()
	if err != nil {
		runtime.LogErrorf(a.ctx, "Error marshaling IPC JSON: %v", err)
		return nil, err
	}

	runtime.LogInfof(a.ctx, "IPC JSON response: %s", string(jsonData))
	return jsonData, nil
}

// removes the #deleteMe div via Datastar remove-fragments
func (a *App) RemoveFragmentsUI() ([]byte, error) {
	ipc := datastar.NewIpc()
	// Remove both the demo div and the button itself with a single selector
	ipc.RemoveFragments("#fragmentsContainer") // CSS selector can target multiple elements

	// Return JSON-encoded IPC envelope
	jsonData, err := ipc.JSON()
	if err != nil {
		runtime.LogErrorf(a.ctx, "Error marshaling IPC JSON: %v", err)
		return nil, err
	}

	runtime.LogInfof(a.ctx, "IPC JSON response: %s", string(jsonData))
	return jsonData, nil
}

// Helper function that returns an IPC envelope containing only merge-signals
func (a *App) buildSignalsIPC() ([]byte, error) {
	ipc := datastar.NewIpc()

	// Create the signal data with the current count2 value
	// Note: Using root level "count2" to match the $count2 in the template
	sigJSON, _ := json.Marshal(map[string]any{
		"count2": a.signalCount,
	})

	// Add merge-signals event to the IPC envelope
	ipc.MergeSignals(sigJSON)

	// Return JSON-encoded IPC envelope
	jsonData, err := ipc.JSON()
	if err != nil {
		runtime.LogErrorf(a.ctx, "Error marshaling signals IPC JSON: %v", err)
		return nil, err
	}

	runtime.LogInfof(a.ctx, "Signals IPC JSON response: %s", string(jsonData))
	return jsonData, nil
}

// IncrementSignals increments the signal counter and returns a merge-signals event
func (a *App) IncrementSignals() ([]byte, error) {
	a.signalCount++
	return a.buildSignalsIPC()
}

// DecrementSignals decrements the signal counter and returns a merge-signals event
func (a *App) DecrementSignals() ([]byte, error) {
	a.signalCount--
	return a.buildSignalsIPC()
}

// RemoveSignalsUI removes the signals UI container via remove-fragments
func (a *App) RemoveSignalsUI() ([]byte, error) {
	ipc := datastar.NewIpc()
	ipc.RemoveFragments("#signalsContainer")

	// Return JSON-encoded IPC envelope
	jsonData, err := ipc.JSON()
	if err != nil {
		runtime.LogErrorf(a.ctx, "Error marshaling remove IPC JSON: %v", err)
		return nil, err
	}

	runtime.LogInfof(a.ctx, "Remove IPC JSON response: %s", string(jsonData))
	return jsonData, nil
}
