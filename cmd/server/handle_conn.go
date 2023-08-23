package main

import (
	"net"
	"sync"
	"time"
)

var (
	enterChannel   = make(chan *User, 100)
	leaveChannel   = make(chan *User, 100)
	MessageChannel = make(chan string, 100)
	globalID       int
	IDLock         sync.Mutex
)

type User struct {
	ID      int       `json:"id"`
	EnterAt time.Time `json:"enter_at"`
}

func getUID() int {
	IDLock.Lock()
	defer IDLock.Unlock()

	globalID++
	return globalID
}

func HandleConn(conn net.Conn) {
	defer conn.Close()

	var user = &User{
		ID:      getUID(),
		EnterAt: time.Now(),
	}

	enterChannel <- user
}
