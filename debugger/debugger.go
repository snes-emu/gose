package debugger

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/snes-emu/gose/config"
)

// Launch start the debugger and open the web page
func Launch() {
	s := createServer()

	go func() {
		err := s.ListenAndServe()
		fmt.Println(err)
	}()

	fmt.Println("open web browser")
	cmd := exec.Command("xdg-open", fmt.Sprintf("http://localhost:%d", config.DebugPort()))
	cmd.Start()
}

func createServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	return &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", config.DebugPort()),
		Handler: mux,
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
