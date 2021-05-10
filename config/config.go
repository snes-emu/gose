package config

import "flag"

var (
	debugServer bool
	debugLogs   bool
	debugPort   int
	nCycles     int
	imgPath     string
)

func init() {
	flag.BoolVar(&debugServer, "debug-server", false, "enable the debug server")
	flag.BoolVar(&debugLogs, "debug-logs", false, "enable debug logs")
	flag.IntVar(&debugPort, "debug-port", 6060, "port the debugger listens to")
	flag.IntVar(&nCycles, "n-cycles", 10_000_000, "number of CPU cycles to execute before exiting (only used in 'save screen as image mode')")
	flag.StringVar(&imgPath, "image-path", "", "where to store the screen as an image (the emulator will exit after ${n-cycles} CPU cycles)")
}

// Inits the config
func Init() {
	flag.Parse()
}

// DebugLogs is used to know whether to enable debug logs or not
func DebugLogs() bool {
	return debugLogs
}

// DebugServer is used to know whether to start the debug server
func DebugServer() bool {
	return debugServer
}

// DebugPort is the port the debugger listens to
func DebugPort() int {
	return debugPort
}

// ImagePath is the path to the file we should use to store the screen as an image (it can be empty -> non set)
func ImagePath() string {
	return imgPath
}

// NCycles is the number of CPU cycles we should execute before storing the screen as an image (only used when ImagePath is not empty)
func NCycles() int {
	return nCycles
}
