package debugger

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/snes-emu/gose/core"
)

type Debugger struct {
	emu  *core.Emulator
	addr string
	s    *http.Server
}

// New debugger instance
func New(emu *core.Emulator, addr string) *Debugger {
	db := &Debugger{
		emu:  emu,
		addr: addr,
	}
	db.createServer(addr)

	return db
}

//Start the debugger and open the web page
func (db *Debugger) Start() {
	go func() {
		err := db.s.ListenAndServe()
		fmt.Println(err)
	}()

	fmt.Printf("open web browser at %s\n", db.addr)
	cmd := exec.Command("xdg-open", fmt.Sprintf("http://%s", db.addr))
	cmd.Start()
}

func (db *Debugger) createServer(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", db.home)
	mux.HandleFunc("/pause", db.pause)
	mux.HandleFunc("/step", db.step)

	db.s = &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}

func (db *Debugger) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(
		`
<html>
	<head>
		<title>Gose debugger</title>
	</head>
	<body>
		<button onClick="fetch('/pause')">toggle pause</button>
		<button onClick="fetch('/step')">step</button>
	</body>
</html>
`))
}

func (db *Debugger) pause(w http.ResponseWriter, r *http.Request) {
	db.emu.TogglePause()

}

func (db *Debugger) step(w http.ResponseWriter, r *http.Request) {
	db.emu.Step(1)

}
