package ws

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type ConnManager struct {
	conns []*websocket.Conn
	sync.Mutex
}

func (cm *ConnManager) AddConn(c *websocket.Conn) {
	cm.Lock()
	cm.conns = append(cm.conns, c)
	cm.Unlock()
	go cm.checkStatus(c)
}

func (cm *ConnManager) checkStatus(conn *websocket.Conn) {
	log.Printf("Connection established with: %s\n", conn.RemoteAddr())
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {

			log.Printf("Connection: %s Closed. Error Message:%s\n", conn.RemoteAddr(), err)
			cm.RemoveConn(conn)
			return
		}
	}
}

func (cm *ConnManager) RemoveConn(c *websocket.Conn) {
	idx := cm.FindConn(c)
	if idx < 0 {
		log.Printf("Connection index %d for %s not found, Total:%d\n",
			idx, c.RemoteAddr(), len(cm.conns))
		return
	}
	cm.Lock()
	defer cm.Unlock()
	if cm.conns[idx] != c {
		log.Printf("Connection Mismatch, index %d for %s, Total:%d\n",
			idx, c.RemoteAddr(), len(cm.conns))
		return
	}
	cm.conns[idx].Close()
	cm.conns = append(cm.conns[0:idx], cm.conns[idx+1:]...)
}

func (cm *ConnManager) FindConn(c *websocket.Conn) int {
	cm.Lock()
	defer cm.Unlock()

	for i, _ := range cm.conns {
		if c == cm.conns[i] {
			return i
		}
	}
	return -1
}

func (cm *ConnManager) Size() int {
	cm.Lock()
	defer cm.Unlock()
	return len(cm.conns)
}

func (cm *ConnManager) Broadcast(mType int, content []byte) int {
	count := 0
	cm.Lock()
	for _, c := range cm.conns {
		if err := c.WriteMessage(mType, content); err != nil {
			log.Printf("An Error occured when writing to connection\n%s\n", err)
			continue
		}
		count += 1
	}
	cm.Unlock()
	return count
}

func (cm *ConnManager) BroadcastJson(v interface{}) int {
	count := 0
	cm.Lock()
	for _, c := range cm.conns {
		if err := c.WriteJSON(v); err != nil {
			log.Printf("An Error occured when json writing to connection\n%s\n", err)
			continue
		}
		count += 1
	}
	cm.Unlock()
	return count
}
