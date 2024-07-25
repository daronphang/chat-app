package ws

import (
	"bytes"
	"chat-service/internal"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 30 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
	logger, _ = internal.WireLogger()
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// CORS.
		CheckOrigin: func(r *http.Request) bool { 
			return true
		},
	}
)

type WebSocketError struct {
	Message string
}

func (e WebSocketError) Error() string {
	return e.Message
}

// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (d *Device) readPump() {
	defer func() {
		d.conn.Close()
		hub.unregisterDevice <- d
	}()

	d.conn.SetReadLimit(maxMessageSize)
	d.conn.SetReadDeadline(time.Now().Add(pongWait))
	d.conn.SetPongHandler(func(string) error { d.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	
	for {
		_, msg, err := d.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure){
				logger.Error("unexpected websocket close error", zap.String("trace", err.Error()))
			}
			return
		} 
		
		msg = bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
		buffer := new(bytes.Buffer)
		if err := json.Compact(buffer, msg); err != nil {
			logger.Error(
				"unable to compact inbound message in JSON",
				zap.String("payload", string(msg)),
			)
			continue
		}
		msg = buffer.Bytes()

		d.hub.receive <- msg
	}
}

func (d *Device) writePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		d.conn.Close()
	}()

	for {
		select {
		case data, ok := <- d.send:
			if !ok {
				// Channel is closed.
				return 
			}
			d.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := d.conn.WriteMessage(websocket.TextMessage, data); err != nil {
				return
			}
		case data, ok := <- d.presence:
			if !ok {
				return 
			}
			d.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := d.conn.WriteMessage(websocket.TextMessage, data); err != nil {
				return
			}
		case <- ticker.C:
			// To keep connection alive by sending ping to client,
			// and preventing readTimeout.
			d.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := d.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        logger.Error(
			"unable to establish websocket connection",
			zap.String("trace", err.Error()),
		)
        return
    }

	params := r.URL.Query()
	clientID := params.Get("client")
	if clientID == "" {
		logger.Error("missing client in query params of websocket url")
		conn.Close()
		return 
	}

	deviceID := params.Get("device")
	if deviceID == "" {
		logger.Error("missing device in query params of websocket url")
		conn.Close()
		return
	}

	device := &Device{
		hub: hub,
		clientID: clientID,
		deviceID: deviceID,
		conn: conn, 
		send: make(chan []byte),
		presence: make(chan []byte),
	}
	hub.registerDevice <- device

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go device.readPump()
	go device.writePump()
}