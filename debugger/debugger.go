package debugger

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/snes-emu/gose/core"
	"go.uber.org/zap"
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
		zap.L().Error("an error occured with the debug server", zap.Error(err))
	}()

	url := fmt.Sprintf("http://%s", db.addr)
	zap.L().Info("open web browser at", zap.String("url", url))
	for _, open := range []string{"xdg-open", "open"} {
		if err := openURL(open, url); err == nil {
			zap.L().Debug("failed to open url",
				zap.String("url", url),
				zap.Error(err),
			)
			break
		}
	}

}

func openURL(program, url string) error {
	cmd := exec.Command(program, url)
	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
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
	w.Write(front)
}

func (db *Debugger) pause(w http.ResponseWriter, r *http.Request) {
	db.emu.TogglePause()

}

func (db *Debugger) step(w http.ResponseWriter, r *http.Request) {
	db.emu.Step(1)

}
