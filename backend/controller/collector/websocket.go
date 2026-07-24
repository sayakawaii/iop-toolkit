package collector

import (
	"encoding/json"
	"omciAnalyzer/utils"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocketConnection represents a client WebSocket connection
type WebSocketConnection struct {
	Conn      *websocket.Conn
	RequestID string
	mu        sync.Mutex
}

// WebSocketManager manages all active WebSocket connections
type WebSocketManager struct {
	connections map[string][]*WebSocketConnection // requestID -> list of connections
	mu          sync.RWMutex
}

var wsManager = &WebSocketManager{
	connections: make(map[string][]*WebSocketConnection),
}

// GetWebSocketManager returns the singleton WebSocket manager
func GetWebSocketManager() *WebSocketManager {
	return wsManager
}

// AddConnection adds a new WebSocket connection for a request ID
func (m *WebSocketManager) AddConnection(requestID string, conn *websocket.Conn) *WebSocketConnection {
	m.mu.Lock()
	defer m.mu.Unlock()

	wsConn := &WebSocketConnection{
		Conn:      conn,
		RequestID: requestID,
	}

	if m.connections[requestID] == nil {
		m.connections[requestID] = []*WebSocketConnection{}
	}
	m.connections[requestID] = append(m.connections[requestID], wsConn)

	utils.Log("WebSocket connection added for request ID:", requestID, "total connections:", len(m.connections[requestID]))
	return wsConn
}

// RemoveConnection removes a WebSocket connection
func (m *WebSocketManager) RemoveConnection(requestID string, wsConn *WebSocketConnection) {
	m.mu.Lock()
	defer m.mu.Unlock()

	connections := m.connections[requestID]
	if connections == nil {
		return
	}

	// Find and remove the connection
	for i, conn := range connections {
		if conn == wsConn {
			m.connections[requestID] = append(connections[:i], connections[i+1:]...)
			break
		}
	}

	// Clean up empty lists
	if len(m.connections[requestID]) == 0 {
		delete(m.connections, requestID)
	}

	utils.Log("WebSocket connection removed for request ID:", requestID, "remaining connections:", len(m.connections[requestID]))
}

// NotifyUpdate notifies all clients subscribed to a request ID about an update
func (m *WebSocketManager) NotifyUpdate(requestID string, data interface{}) {
	m.mu.RLock()
	connections := m.connections[requestID]
	m.mu.RUnlock()

	if len(connections) == 0 {
		utils.Log("No WebSocket connections to notify for request ID:", requestID)
		return
	}

	message := map[string]interface{}{
		"type":       "update",
		"request_id": requestID,
		"data":       data,
		"timestamp":  time.Now().Unix(),
	}

	payload, err := json.Marshal(message)
	if err != nil {
		utils.Log("Failed to marshal WebSocket message:", err)
		return
	}

	utils.Log("Notifying", len(connections), "WebSocket connections for request ID:", requestID)

	// Send to all connections (create a copy to avoid holding the lock while sending)
	connsCopy := make([]*WebSocketConnection, len(connections))
	copy(connsCopy, connections)

	for _, wsConn := range connsCopy {
		go func(conn *WebSocketConnection) {
			conn.mu.Lock()
			defer conn.mu.Unlock()

			if err := conn.Conn.WriteMessage(websocket.TextMessage, payload); err != nil {
				utils.Log("Failed to send WebSocket message:", err)
				// Connection might be closed, remove it
				m.RemoveConnection(requestID, conn)
			}
		}(wsConn)
	}
}

// HandleConnection handles a WebSocket connection for a request ID
func (m *WebSocketManager) HandleConnection(conn *websocket.Conn, requestID string) {
	wsConn := m.AddConnection(requestID, conn)
	defer func() {
		m.RemoveConnection(requestID, wsConn)
		conn.Close()
	}()

	// Set read deadline and pong handler for connection health check
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	// Start ping ticker
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	done := make(chan struct{})

	// Goroutine to send ping messages
	go func() {
		for {
			select {
			case <-ticker.C:
				wsConn.mu.Lock()
				if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					wsConn.mu.Unlock()
					utils.Log("Failed to send ping:", err)
					close(done)
					return
				}
				wsConn.mu.Unlock()
			case <-done:
				return
			}
		}
	}()

	// Read messages from client (mainly to detect disconnection)
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				utils.Log("WebSocket error:", err)
			}
			close(done)
			break
		}
	}
}
