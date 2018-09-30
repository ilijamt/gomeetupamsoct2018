package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func init() {
	runtime.SetBlockProfileRate(1)
}

func main() {
	log.Fatal(http.ListenAndServe(":6060", nil))
}
