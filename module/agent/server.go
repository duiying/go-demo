package agent

import (
	"github.com/duiying/go-demo/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Hub 维护所有的 connection
type Hub struct {
	// 当前在线的 connection
	online map[*WSConn]bool
	// 加入的 connection
	join chan *WSConn
	// 关闭的 connection
	close chan *WSConn
	// 广播数据
	broadCast chan []byte
}

var hub Hub

func InitHub() {
	hub = Hub{
		online: make(map[*WSConn]bool),
		join:   make(chan *WSConn),
		close:  make(chan *WSConn),
	}
}

func RunHub() {
	for {
		select {
		// 用户连接，添加 connection
		case conn := <-hub.join:
			hub.online[conn] = true
		// 删除指定用户连接
		case conn := <-hub.close:
			if _, ok := hub.online[conn]; ok {
				delete(hub.online, conn)
				close(conn.recChan)
				close(conn.sendChan)
				_ = conn.ws.Close()
			}
		// 广播
		case data := <-hub.broadCast:
			for conn := range hub.online {
				conn.recChan <- data
			}
		}
	}
}

func WS(c *gin.Context) {
	if !c.IsWebsocket() {
		response.Fail(c, response.NotWebSocketError)
		return
	}

	// 升级请求为 WebSocket 协议
	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer func() {
		_ = wsConn.Close()
	}()

	ip := c.ClientIP()
	data := []byte(c.Request.URL.RawQuery)

	conn := &WSConn{
		ws:         wsConn,
		data:       data,
		remoteAddr: ip,
		recChan:    make(chan []byte, 200),
		sendChan:   make(chan []byte, 200),
	}

	// connection 加入 hub 管理
	hub.join <- conn

	conn.wg.Add(3)
	go conn.rec()
	go conn.doRec()
	go conn.doSnd()
	conn.wg.Wait()
}
