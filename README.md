# gose
![Build Status](https://github.com/snes-emu/gose/workflows/Go/badge.svg?branch=master)

![Gose logo](./logo.svg)

Gose is a WIP SNES emulator

## Usage

You can start by cloning the repo with: `git clone https://github.com/snes-emu/gose.git`

Then to build gose you can simply do `make` in the `gose` project directory, this should output a binary named `gose` in your current directory.

You can then run a ROM doing: `./gose <path_to_your_rom>`, for now only raw ROMs and `.zip`s are supported

To run the ROM with the debugger enabled you can do: `./gose -debug-server <path_to_your_rom>` if the web debugger did not open automatically, check in the logs for the URL to open in your browser.

### Testing

To run the tests simply run: `make test`
