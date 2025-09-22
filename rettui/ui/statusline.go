package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type statusLine struct {
    left  *left
    right *right
}

type left struct {
    textview    *tview.TextView
    method      string
    url         string
}
type right struct {
    textview    *tview.TextView
    timing      string
    size        string
    status      string
}

func newStatusLine(app *tview.Application) *statusLine {
    statusLine := &statusLine{
        left: &left{
            textview: tview.NewTextView().SetDynamicColors(true).SetTextAlign(tview.AlignLeft).SetChangedFunc(func(){app.Draw()}),
            method: "GET",
            // url: "󰌷 <URL>",
            url: "󰌷 example.com",
        },
        right: &right{
            textview:tview.NewTextView().SetDynamicColors(true).SetTextAlign(tview.AlignRight).SetChangedFunc(func(){app.Draw()}),
            timing: "522ms",
            size: "979B",
            status: "200 OK",
        },
    }
    statusLine.left.textview.SetBackgroundColor(tcell.NewHexColor(0x1b1d2b))
    statusLine.right.textview.SetBackgroundColor(tcell.NewHexColor(0x1b1d2b))
    statusLine.update()
    return statusLine
}

func (s *statusLine) addTo(layout *tview.Grid) {
    layout.AddItem(s.left.textview, 2, 0, 1, 1, 0, 0, false)
    layout.AddItem(s.right.textview, 2, 1, 1, 1, 0, 0, false)
}

func (s *statusLine) update() {
    leftStr := fmt.Sprintf("[black:blue] %s [blue:#444a73] %s [::]", s.left.method, s.left.url)
    rightStr := fmt.Sprintf("%s [blue:#444a73] %s [black:blue] %-3s [::]", s.right.timing, s.right.size, s.right.status)

    if leftStr != s.left.textview.GetText(false) {
        s.left.textview.SetText(leftStr)
    }
    if rightStr != s.right.textview.GetText(false) {
        s.right.textview.SetText(rightStr)
    }
}

func (s *statusLine) SetMethod(method string) {
    s.left.method = method
    s.update()
}

func (s *statusLine) SetURL(url string) {
    s.left.url = url
    s.update()
}

func (s *statusLine) SetStatus(status string) {
    s.right.status = status
    s.update()
}

func (s *statusLine) SetTiming(timing string) {
    s.right.timing = timing
    s.update()
}

func (s *statusLine) SetSize(size string) {
    s.right.size = size
    s.update()
}

