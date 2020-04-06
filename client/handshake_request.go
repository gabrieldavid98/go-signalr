package client

// handshakeRequest is sent by the client to agree on the message format
type handshakeRequest struct {
	Protocol string `json:"protocol"`
	Version  byte   `json:"version"`
}
