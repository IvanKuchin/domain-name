package main

import (
	"log"

	"github.com/ivankuchin/domain-name/internal/pkg/server"
)

func SetLogFlags() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	SetLogFlags()

	server.Run()
}
