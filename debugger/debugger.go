package debugger

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/snes-emu/gose/log"

	"github.com/gobuffalo/packr/v2"
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
		log.Error("an error occurred with the debug server", zap.Error(err))
	}()

	url := fmt.Sprintf("http://%s", db.addr)
	log.Info("open web browser at", zap.String("url", url))
	for _, open := range []string{"xdg-open", "open"} {
		if err := openURL(open, url); err == nil {
			log.Debug("failed to open url",
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
	box := packr.New("front", "./static")
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(box))
	mux.HandleFunc("/pause", db.pause)
	mux.HandleFunc("/step", db.step)
	mux.HandleFunc("/breakpoint", db.breakpoint)

	db.s = &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}

func (db *Debugger) pause(w http.ResponseWriter, r *http.Request) {
	db.emu.TogglePause()
	db.sendState(w)
}

func (db *Debugger) step(w http.ResponseWriter, r *http.Request) {
	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		count = 1
	}

	db.emu.Step(count)
	if err = db.sendState(w); err != nil {
		log.Error("an error occurred while sending current state to the debugger", zap.Error(err))
		w.Write([]byte(err.Error()))
	}
}

func (db *Debugger) breakpoint(w http.ResponseWriter, r *http.Request) {
	log.Debug("/breakpoint")
	address, err := strconv.Atoi(r.URL.Query().Get("address"))
	if err != nil {
		log.Info("fail to set breakpoint", zap.Error(err))
		return
	}

	db.emu.SetBreakpoint(uint32(address))
}

func (db *Debugger) sendState(w http.ResponseWriter) error {
	res := make(map[string]interface{})

	res["palette"] = db.emu.PPU.ExportPalette()
	res["cpu"] = db.emu.CPU
	jsonRes, err := json.Marshal(res)
	if err != nil {
		return errors.Wrap(err, "error marshalling the emulator state")
	}

	_, err = w.Write(jsonRes)
	return errors.Wrap(err, "error writing response")
}
