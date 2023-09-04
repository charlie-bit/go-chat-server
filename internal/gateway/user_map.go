package gateway

import (
	"sync"

	"github.com/gorilla/websocket"
)

type UserMap struct {
	m sync.Map
}

var UserClientMap *UserMap

func NewUserMap() {
	UserClientMap = &UserMap{}
}

func (u *UserMap) Set(key string, v *websocket.Conn) {
	allClients, existed := u.m.Load(key)
	if existed {
		oldClients := allClients.([]*websocket.Conn)
		oldClients = append(oldClients, v)
		u.m.Store(key, oldClients)
	} else {
		var clients []*websocket.Conn
		clients = append(clients, v)
		u.m.Store(key, clients)
	}
}

func (u *UserMap) Get(key string) ([]*websocket.Conn, bool, bool) {
	allClients, userExisted := u.m.Load(key)
	if userExisted {
		var clients []*websocket.Conn
		for _, client := range allClients.([]*websocket.Conn) {
			clients = append(clients, client)
		}
		if len(clients) > 0 {
			return clients, userExisted, true
		}
		return clients, userExisted, false
	}
	return nil, userExisted, false
}
