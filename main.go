package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"
)

const (
	APP_NAME    = "server"
	APP_VERSION = "0.5.0"
)

var (
	confFile    = flag.String("conf", "./server.conf", "specify config file")
	showVersion = flag.Bool("version", false, "print version string")
	profilePort = flag.Int("pprof_port", 6060, "specify profile port")
)

func main() {
	//todo log init
	//todo log version
	flag.Parse()

	//pprof
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("localhost:%d", *profilePort), nil)
		if err != nil {
			fmt.Println("profiler start error:", err.Error())
		}
	}()
	//
	runtime.GOMAXPROCS(runtime.NumCPU())
	//todo load configFile

	//todo module init
	//todo run
	//todo wait
	//todo before exit logic

}
