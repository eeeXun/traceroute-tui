package main

import (
	// "fmt"
	"github.com/rivo/tview"
	"time"
)

var (
	// TraceRoute
	ProtocolICMP      = 1 // from https://pkg.go.dev/golang.org/x/net/internal/iana
	MaxTTL       int8 = 30
	// UI
	app        = tview.NewApplication()
	input_box  = tview.NewInputField()
	output_box = NewOutputScreen()
	// Control
	stop_traceroute       bool = true
	traceroute_thread_cnt      = 0
	// Other
	sec = time.Second
)

func main() {
	UIInit()

	if err := app.SetRoot(
		tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(input_box, 0, 1, true).
			AddItem(output_box, 0, 6, false),
		true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
