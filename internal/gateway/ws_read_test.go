package gateway_test

import (
	"fmt"
	"testing"
	"time"
)

func TestTicketHeartbeat(t *testing.T) {
	ticker := time.NewTicker(time.Second * 5)
	for {
		fmt.Println(1)
		select {
		case <-ticker.C:
			fmt.Println("t")
		default:
		}
	}
}
