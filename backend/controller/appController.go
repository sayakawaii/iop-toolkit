/*
# ------------------------------------------------------------
# -- appController.go
# --
# -- Huang Minghe
# -- 2022-7-12
# ------------------------------------------------------------
*/

package controller

import (
	"sync"

	"github.com/gorilla/websocket"
)

type AppController struct {
	wsConn  *websocket.Conn
	connMux sync.Mutex
	logChan chan string
}
