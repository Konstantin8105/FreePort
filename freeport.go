package freeport

import (
	"net"
)

// Get - return free open TCP port
func Get() (port int, err error) {
	ln, err := net.Listen("tcp", "[::]:0")
	if err != nil {
		return 0, err
	}
	port = ln.Addr().(*net.TCPAddr).Port
	err = ln.Close()
	return
}
