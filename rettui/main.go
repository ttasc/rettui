package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/ttasc/rettui/rettui/ui"
)

func init() {
    ui.ReDefineTviewVars()
}

func main() {
    // Create Tview Application
    tviewApp := tview.NewApplication().EnableMouse(true)

    // Create UI
    mainUI := ui.NewUI(tviewApp)

    tviewApp.SetRoot(mainUI.Root, true)
    tviewApp.SetFocus(nil)

    tviewApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Key() {
        // Unfocus
        case tcell.KeyEscape:
            tviewApp.SetFocus(nil)
        // TODO: Tab navigation
        // ...
        // Request navigation
        case tcell.KeyCtrlP:
            tviewApp.SetFocus(mainUI.MainView.RequestSide.Params)
        case tcell.KeyCtrlH:
            tviewApp.SetFocus(mainUI.MainView.RequestSide.Headers)
        case tcell.KeyCtrlB:
            tviewApp.SetFocus(mainUI.MainView.RequestSide.Body)
        case tcell.KeyCtrlT:
            tviewApp.SetFocus(mainUI.MainView.RequestSide.BodyTypes)
        // Response navigation
        case tcell.KeyCtrlO:
            tviewApp.SetFocus(mainUI.MainView.ResponseSide.Container)
        }
        if event.Rune() == ':' {
            if tviewApp.GetFocus() == nil {
                tviewApp.SetFocus(mainUI.CmdLine)
            }
        }
        return event
    })

    if err := tviewApp.Run(); err != nil {
        panic(err)
    }
}
