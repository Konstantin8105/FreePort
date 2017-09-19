package freeport_test

import (
	"fmt"
	"net"
	"strconv"
	"testing"

	freeport "github.com/Konstantin8105/FreePort"
)

func TestGetFreePort(t *testing.T) {
	port, err := freeport.Get()
	if err != nil {
		t.Errorf("Error for founding to the free TCP port. err = %v", err)
	}
	address := "127.0.0.1:" + strconv.Itoa(port)
	conn, err := net.Listen("tcp", address)
	if err != nil {
		t.Errorf("Error for create the TCP connection on free TCP port. err = %v", err)
	}
	_ = conn.Close()
}

func TestHaveNotFreePort(t *testing.T) {
	// port : from 0 to 65535
	amountPorts := 65536
	conn := make([]net.Listener, amountPorts, amountPorts)
	counter := 0
	for port := 0; port < amountPorts; port++ {
		address := "127.0.0.1:" + strconv.Itoa(port)
		c, err := net.Listen("tcp", address)
		if err != nil {
			continue
		}
		conn[counter] = c
		counter++
	}
	port, err := freeport.Get()
	if err == nil {
		t.Errorf("Find free port, but all port must be busy. Port = ", port)
	}
	for i := 0; i < counter; i++ {
		_ = conn[i].Close()
	}
}

func ExampleGet() {
	port, err := freeport.Get()
	if err != nil {
		panic(err)
	}
	if 0 < port && port < 65535 {
		fmt.Println("Found free tcp port")
	}
	// Output: Found free tcp port
}
