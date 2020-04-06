package client

// Transport decribes the behavior any transport will have
type Transport interface {
	Connect() error
	Send(data interface{}) error
	OnReceive(fn func(data interface{}))
	OnClose(fn func(code int, err string) error)
	Close() error
}

// TransportType defines a transport type
type TransportType int

const (
	// WebSocket is a type of transport
	WebSocket TransportType = iota

	// LongPolling is a type of transport
	LongPolling
)
