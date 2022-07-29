package main

import (
	"github.com/gdamore/tcell/v2"
)

func UIInit() {
	bg := tcell.ColorDarkCyan
	fg := tcell.ColorLightPink

	// tview.Box method
	input_box.SetBorder(true).
		SetTitle("URL/IP").
		SetBackgroundColor(bg).
		SetBorderColor(fg).
		SetTitleColor(fg)
	// tview.InputField method
	input_box.SetLabel("URL/IP: ").
		SetFieldBackgroundColor(tcell.ColorRosyBrown).
		SetFieldTextColor(tcell.ColorLightGray).
		SetLabelColor(fg)
	// Handle key
	input_box.SetDoneFunc(InputDoneHandle)
	input_box.SetInputCapture(InputCaptureHandle)
	output_box.SetInputCapture(OutputCaptureHandle)

	// tview.Box method
	output_box.SetBorder(true).
		SetTitle(output_box.Title).
		SetBackgroundColor(bg).
		SetBorderColor(fg).
		SetTitleColor(fg)
}

func InputDoneHandle(key tcell.Key) {
	switch key {
	case tcell.KeyEnter:
		stop_traceroute = false
		dest := input_box.GetText()
		go TraceRoute(dest)
	case tcell.KeyTab, tcell.KeyBacktab:
		app.SetFocus(output_box)
	case tcell.KeyEscape:
		app.Stop()
	}
}

func InputCaptureHandle(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()
	
	switch key {
	case tcell.KeyCtrlD, tcell.KeyCtrlS:
		stop_traceroute = true
	case tcell.KeyCtrlQ:
		app.Stop()
	}

	return event
}

func OutputCaptureHandle(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()

	switch key {
	case tcell.KeyTab, tcell.KeyBacktab:
		app.SetFocus(input_box)
	case tcell.KeyCtrlD, tcell.KeyCtrlS:
		stop_traceroute = true
	case tcell.KeyCtrlL:
		output_box.ClearText()
	case tcell.KeyCtrlQ, tcell.KeyEscape:
		app.Stop()
	}

	return event
}
