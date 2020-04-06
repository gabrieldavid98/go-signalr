package client

// handshakeResponse is sent by the server as an acknowledgment of the previous
// HandshakeRequest message. Contains an error if the handshake failed.
type handshakeResponse struct {
	Error string `json:"error"`
}
