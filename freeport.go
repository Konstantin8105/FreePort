package freeport

import (
	"fmt"
	"net"
)

// GetFreePort - return free open TCP port
func GetFreePort() (port int, err error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer func() {
		err2 := l.Close()
		if err2 != nil {
			if err != nil {
				err = fmt.Errorf("%v\n%v", err, err2)
			} else {
				err = err2
			}
		}
	}()
	return l.Addr().(*net.TCPAddr).Port, nil
}
