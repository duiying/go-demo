package agent

import (
	"github.com/duiying/go-demo/pkg/logger"
	"github.com/gorilla/websocket"
	"sync"
)

type WSConn struct {
	ws         *websocket.Conn
	data       []byte
	remoteAddr string
	recChan    chan []byte
	sendChan   chan []byte
	wg         sync.WaitGroup
}

// 数据读取器
func (conn *WSConn) rec() {
	defer func() {
		hub.close <- conn
		conn.wg.Done()
	}()
	for {
		_, data, err := conn.ws.ReadMessage()
		if err != nil {
			logger.Error("ws read message error", "err", err)
			return
		}
		conn.recChan <- data
	}
}

// 数据读取处理
func (conn *WSConn) doRec() {
	defer func() {
		hub.close <- conn
		conn.wg.Done()
	}()
	for {
		select {
		case data, ok := <-conn.recChan:
			if !ok {
				return
			}
			if len(data) < 1 {
				logger.Error("ws rec data len error", "data", data)
				return
			}
			logger.Debug("received websocket data", data)
			hub.broadCast <- data
		}
	}
}

// 数据接收处理
func (conn *WSConn) doSnd() {
	defer func() {
		hub.close <- conn
		conn.wg.Done()
	}()

	for {
		select {
		case data, ok := <-conn.sendChan:
			if !ok {
				return
			}
			if len(data) < 1 {
				return
			}
			err := conn.ws.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				return
			}
		}
	}
}
