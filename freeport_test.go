package freeport_test

import (
	"fmt"
	"net"
	"strconv"
	"testing"

	freeport "github.com/Konstantin8105/FreePort"
)

func TestGetFreePort(t *testing.T) {
	port, err := freeport.GetFreePort()
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
	for port := 0; port < 65536; port++ {
		address := "127.0.0.1:" + strconv.Itoa(port)
		fmt.Println("port = ", address)
		conn, _ := net.Listen("tcp", address)
		defer func() {
			_ = conn.Close()
		}()
	}
	port, err := freeport.GetFreePort()
	if err == nil {
		t.Errorf("Find free port, but all port must be busy. Port = ", port)
	}
}
