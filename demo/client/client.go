// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":2023")
	if err != nil {
		panic(err)
	}

	fmt.Println(conn)

	var message = "hello"
	fmt.Println(conn.Write([]byte(message)))
}
