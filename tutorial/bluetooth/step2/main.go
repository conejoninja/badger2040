package main

import (
	"fmt"
	"time"

	"tinygo.org/x/bluetooth"
	"tinygo.org/x/tinyfont/proggy"
	"tinygo.org/x/tinyterm"
	"tinygo.org/x/tinyterm/displays"
)

var (
	terminal *tinyterm.Terminal

	adapter = bluetooth.DefaultAdapter
)

func main() {
	initTerminal()

	terminalOutput("enable interface...")

	must("enable BLE interface", adapter.Enable())
	time.Sleep(time.Second)

	terminalOutput("start scan...")

	must("start scan", adapter.Scan(scanHandler))

	for {
		time.Sleep(time.Minute)
		terminalOutput("scanning...")
	}
}

func scanHandler(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
	msg := fmt.Sprintf("%s %d %s", device.Address.String(), device.RSSI, device.LocalName())
	terminalOutput(msg)
}

func must(action string, err error) {
	if err != nil {
		for {
			terminalOutput("failed to " + action + ": " + err.Error())

			time.Sleep(time.Second)
		}
	}
}

func terminalOutput(s string) {
	println(s)
	fmt.Fprintf(terminal, "\n%s", s)

	terminal.Display()
}

var (
	font = &proggy.TinySZ8pt7b
)

func initTerminal() {
	display := displays.Init()

	terminal = tinyterm.NewTerminal(display)
	terminal.Configure(&tinyterm.Config{
		Font:              font,
		FontHeight:        10,
		FontOffset:        6,
		UseSoftwareScroll: true,
	})
}
