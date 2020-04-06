package client

import (
	"context"
)

// OnFunc represents the callback who is executed
// when new data arrives from server
type OnFunc func(data interface{})

// HubConnectionState represents the state of the hub connection
type HubConnectionState int

const (
	// Disconnected is a hub connection state
	Disconnected HubConnectionState = iota

	// Connected is a hub connection state
	Connected

	// Connecting is a hub connection state
	Connecting

	// Reconnecting is a hub connection state
	Reconnecting
)

// HubConnection represents a hub connection behavior
type HubConnection interface {
	Start(ctx context.Context) error
	Stop() error
	Send(methodName string, args ...interface{}) error
	Channel(methodName string) <-chan interface{}
	On(methodName string, onFunc OnFunc) context.CancelFunc
	GetConnectionId() string
	GetHubConnectionState() HubConnectionState
}

type hubConnection struct {
}

func newHubConnection() HubConnection {
	return nil
}
