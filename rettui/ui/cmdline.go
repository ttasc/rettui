package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type cmdLine struct{
    *tview.InputField
    shouldFocus bool
}

func newCmdLine(app *tview.Application) *cmdLine {
    cmdLine := &cmdLine{
        tview.NewInputField().SetFieldBackgroundColor(Colors["dark"]),
        false,
    }
    cmdLine.SetDoneFunc(func(key tcell.Key) {
        app.SetFocus(nil)
    })
    cmdLine.SetFocusFunc(func() {
        cmdLine.shouldFocus = true
        cmdLine.SetText("")
        cmdLine.shouldFocus = false
    })
    cmdLine.SetChangedFunc(func(text string) {
        if text == "" && !cmdLine.shouldFocus {
            app.SetFocus(nil)
        }
    })

    cmdLine.SetMouseCapture(func(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {
        return action, nil
    })
    return cmdLine
}

func (c *cmdLine) addTo(layout *tview.Grid) {
    layout.AddItem(c, 3, 0, 1, 2, 0, 0, false)
}

