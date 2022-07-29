package main

import(
	"github.com/rivo/tview"
)

type OutputScreen struct {
	*tview.TextView
	Title   string
	Text string
}

func NewOutputScreen() *OutputScreen {
	return &OutputScreen{
		TextView: tview.NewTextView(),
		Title: "TraceRoute",
	}
}

func (screen *OutputScreen) AddText(s string) {
	if len(screen.Text) == 0 {
		screen.Text = s
	} else {
		screen.Text = s + "\n" + screen.Text
	}
}

func (screen OutputScreen) UpdateTitle() {
	screen.SetTitle(screen.Title)
}

// Concurrency(app.Draw), do not call in main thread
func (screen OutputScreen) RefreshText() {
	screen.SetText(screen.Text)
	screen.ScrollToBeginning()
	app.Draw()
}

func (screen *OutputScreen) ClearText() {
	screen.Text = ""
	screen.SetText(screen.Text)
}
