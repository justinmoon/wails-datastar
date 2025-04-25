package datastar

import (
	"encoding/json"
)

// EventType represents the type of an IPC event
type EventType string

// Common EventTypes that match Datastar's SSE event types
const (
	TypeStarted         EventType = "datastar-started"
	TypeFinished        EventType = "datastar-finished"
	TypeError           EventType = "datastar-error"
	TypeMergeFragments  EventType = "datastar-merge-fragments"
	TypeMergeSignals    EventType = "datastar-merge-signals"
	TypeRemoveFragments EventType = "datastar-remove-fragments"
	TypeRemoveSignals   EventType = "datastar-remove-signals"
	TypeExecuteScript   EventType = "datastar-execute-script"
)

// IpcEvent represents a single event in the IPC envelope
type IpcEvent struct {
	Type EventType         `json:"type"`
	Args map[string]string `json:"args"`
}

// IpcBuilder collects events to be sent as an IPC response
type IpcBuilder struct {
	events []IpcEvent
}

// NewIpc creates a new IpcBuilder
func NewIpc() *IpcBuilder {
	return &IpcBuilder{}
}

// MergeFragmentOption represents options for merging fragments
type MergeFragmentOption func(map[string]string)

// WithSelectorID sets the selector ID to merge fragments into
func WithSelectorID(id string) MergeFragmentOption {
	return func(args map[string]string) {
		args["selector"] = id
	}
}

// WithBeforeSelector sets the position to 'before' for fragment insertion
func WithBeforeSelector() MergeFragmentOption {
	return func(args map[string]string) {
		args["position"] = "before"
	}
}

// WithAfterSelector sets the position to 'after' for fragment insertion
func WithAfterSelector() MergeFragmentOption {
	return func(args map[string]string) {
		args["position"] = "after"
	}
}

// WithAppendSelector sets the position to 'append' for fragment insertion
func WithAppendSelector() MergeFragmentOption {
	return func(args map[string]string) {
		args["position"] = "append"
	}
}

// WithPrependSelector sets the position to 'prepend' for fragment insertion
func WithPrependSelector() MergeFragmentOption {
	return func(args map[string]string) {
		args["position"] = "prepend"
	}
}

// MergeFragments adds a merge-fragments event to the IPC envelope
func (b *IpcBuilder) MergeFragments(frag string, opts ...MergeFragmentOption) {
	args := map[string]string{
		"fragments": frag,
	}

	// Apply all options
	for _, opt := range opts {
		opt(args)
	}

	b.events = append(b.events, IpcEvent{
		Type: TypeMergeFragments,
		Args: args,
	})
}

// MergeSignalsOption represents options for merging signals
type MergeSignalsOption func(map[string]string)

// MergeSignals adds a merge-signals event to the IPC envelope
func (b *IpcBuilder) MergeSignals(jsonData []byte, opts ...MergeSignalsOption) {
	args := map[string]string{
		"signals": string(jsonData),
	}

	// Apply all options
	for _, opt := range opts {
		opt(args)
	}

	b.events = append(b.events, IpcEvent{
		Type: TypeMergeSignals,
		Args: args,
	})
}

// RemoveFragmentsOption represents options for removing fragments
type RemoveFragmentsOption func(map[string]string)

// RemoveFragments adds a remove-fragments event to the IPC envelope
func (b *IpcBuilder) RemoveFragments(selector string, opts ...RemoveFragmentsOption) {
	args := map[string]string{
		"selector": selector,
	}

	// Apply all options
	for _, opt := range opts {
		opt(args)
	}

	b.events = append(b.events, IpcEvent{
		Type: TypeRemoveFragments,
		Args: args,
	})
}

// RemoveSignalsOption represents options for removing signals
type RemoveSignalsOption func(map[string]string)

// RemoveSignals adds a remove-signals event to the IPC envelope
func (b *IpcBuilder) RemoveSignals(path string, opts ...RemoveSignalsOption) {
	args := map[string]string{
		"path": path,
	}

	// Apply all options
	for _, opt := range opts {
		opt(args)
	}

	b.events = append(b.events, IpcEvent{
		Type: TypeRemoveSignals,
		Args: args,
	})
}

// ExecuteScriptOption represents options for executing scripts
type ExecuteScriptOption func(map[string]string)

// ExecuteScript adds an execute-script event to the IPC envelope
func (b *IpcBuilder) ExecuteScript(code string, opts ...ExecuteScriptOption) {
	args := map[string]string{
		"code": code,
	}

	// Apply all options
	for _, opt := range opts {
		opt(args)
	}

	b.events = append(b.events, IpcEvent{
		Type: TypeExecuteScript,
		Args: args,
	})
}

// JSON marshals the IPC envelope to JSON
func (b *IpcBuilder) JSON() ([]byte, error) {
	return json.Marshal(b.events)
}

