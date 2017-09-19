package freeport

import (
	"fmt"
	"net"
)

// Get - return free open TCP port
func Get() (port int, err error) {
	ln, err := net.Listen("tcp", "[::]:0")
	if err != nil {
		return 0, err
	}
	defer func() {
		err2 := ln.Close()
		if err2 != nil {
			if err != nil {
				err = fmt.Errorf("%v\n%v", err, err2)
			} else {
				err = err2
			}
		}
	}()
	return ln.Addr().(*net.TCPAddr).Port, nil
}
