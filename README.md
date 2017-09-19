# FreePort

[![Coverage Status](https://coveralls.io/repos/github/Konstantin8105/FreePort/badge.svg?branch=master)](https://coveralls.io/github/Konstantin8105/FreePort?branch=master)
[![Build Status](https://travis-ci.org/Konstantin8105/FreePort.svg?branch=master)](https://travis-ci.org/Konstantin8105/FreePort)
[![Go Report Card](https://goreportcard.com/badge/github.com/Konstantin8105/FreePort)](https://goreportcard.com/report/github.com/Konstantin8105/FreePort)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Konstantin8105/FreePort/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/Konstantin8105/FreePort?status.svg)](https://godoc.org/github.com/Konstantin8105/FreePort)

Return a free tcp port.

Minimal example of using:

```golang
func main(){
	port, err := freeport.Get()
	if err != nil {
		panic(err)
	}
	if 0 < port && port < 65536 {
		fmt.Println("Found free tcp port")
	}
}
```
