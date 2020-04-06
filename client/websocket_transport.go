package client

import (
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type websocketTransport struct {
	conn         *websocket.Conn
	httpResponse *http.Response
	url          string
	onreceive    func(data interface{})
}

func (w *websocketTransport) Connect() error {
	var err error

	w.url = strings.Replace(w.url, "http", "ws", 1)
	w.conn, w.httpResponse, err = websocket.DefaultDialer.Dial(w.url, nil)

	if err != nil {
		return err
	}

	go w.read()

	return nil
}

func (w *websocketTransport) Send(data interface{}) error {
	return w.conn.WriteJSON(data)
}

func (w *websocketTransport) Close() error {
	return w.conn.Close()
}

func (w *websocketTransport) OnReceive(fn func(data interface{})) {
	w.onreceive = fn
}

func (w *websocketTransport) OnClose(fn func(code int, text string) error) {
	w.conn.SetCloseHandler(fn)
}

func (w *websocketTransport) read() {
	for {
		var data interface{}
		err := w.conn.ReadJSON(data)
		if err != nil {
			return
		}

		w.onreceive(data)
	}
}

// newWebsocketTransport creates a new transport based on websocket
func newWebsocketTransport(url string) Transport {
	return &websocketTransport{url: url}
}
