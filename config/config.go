package config

import "flag"

var (
	debugServer bool
	debugPort   int
)

func init() {
	flag.BoolVar(&debugServer, "debug-server", false, "enable the debug server")
	flag.IntVar(&debugPort, "debug-port", 6060, "port the debugger listens to")
}

// Inits the config
func Init() {
	flag.Parse()
}

// DebugServer is used to know whether to start the debug server
func DebugServer() bool {
	return debugServer
}

// DebugPort is the port the debugger listens to
func DebugPort() int {
	return debugPort
}
