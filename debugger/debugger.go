package debugger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/snes-emu/gose/log"
	"image/png"
	"net/http"
	"os/exec"
	"strconv"

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
	mux.HandleFunc("/resume", db.resume)
	mux.HandleFunc("/step", db.step)
	mux.HandleFunc("/breakpoint", db.breakpoint)

	db.s = &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}

func (db *Debugger) resume(w http.ResponseWriter, r *http.Request) {
	db.emu.Resume()
	// Wait for the next breakpoint to be reached
	breakp := <-db.emu.BreakpointCh
	if breakp.IsRegister {
		db.sendStateWithRegister(breakp, w)
	} else {
		db.sendState(w)
	}
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
	rawAddr := r.URL.Query().Get("address")
	registers := r.URL.Query().Get("registers")

	if rawAddr != "" {
		// Address breakpoint
		address, err := strconv.Atoi(rawAddr)
		if err != nil {
			log.Error("failed to set breakpoint", zap.Error(err))
		} else {
			log.Info("Setting address breakpoint", zap.Int("address", address))
			db.emu.SetBreakpoint(uint32(address))
		}
	}

	if registers != "" {
		// Register breakpoint
		log.Info("Setting register breakpoints", zap.String("breakpoints", registers))
		db.emu.SetRegisterBreakpoint(registers)
	}
}

func (db *Debugger) emulatorState() map[string]interface{} {
	res := make(map[string]interface{})
	res["palette"] = db.emu.PPU.Palette()
	res["cpu"] = db.emu.CPU.Export()

	sprites := db.emu.PPU.Sprites()
	// Will store base64 encoded sprite images
	encoded := make([][]byte, len(sprites))

	for i, sprite := range sprites {
		buf := &bytes.Buffer{}
		if err := png.Encode(buf, sprite); err != nil {
			log.Error("error encoding sprite, this sprite will be skipped", zap.Int("sprite_number", i), zap.Error(err))
			continue
		}

		encoded[i] = buf.Bytes()
	}

	res["sprites"] = encoded

	return res
}

func (db *Debugger) sendStateWithRegister(register core.BreakpointData, w http.ResponseWriter) error {
	res := db.emulatorState()
	res["register"] = register
	return db.send(res, w)
}

func (db *Debugger) sendState(w http.ResponseWriter) error {
	return db.send(db.emulatorState(), w)
}

func (db *Debugger) send(payload map[string]interface{}, w http.ResponseWriter) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshalling the emulator state: %w", err)
	}

	if _, err = w.Write(jsonPayload); err != nil {
		return fmt.Errorf("error writing response: %w", err)
	}

	return nil
}
