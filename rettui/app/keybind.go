package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/ttasc/rettui/rettui/ui"
)

var (
    mainElements []tview.Primitive
    currentFocus int
)

func setKeyBinds(app *tview.Application, ui *ui.UI) {

    mainElements = []tview.Primitive{
        ui.MainView.ReqSide.Params,
        ui.MainView.ReqSide.Headers,
        ui.MainView.ReqSide.BodyTypes,
        ui.MainView.ReqSide.Body,
    }

    app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Key() {
            case tcell.KeyEscape: app.SetFocus(nil)

            case tcell.KeyCtrlJ:
                currentFocus = (currentFocus+1)%len(mainElements)
                app.SetFocus(mainElements[currentFocus])
            case tcell.KeyCtrlK:
                currentFocus = (currentFocus-1+len(mainElements))%len(mainElements)
                app.SetFocus(mainElements[currentFocus])

            case tcell.KeyCtrlH: app.SetFocus(mainElements[currentFocus])
            case tcell.KeyCtrlL: app.SetFocus(ui.MainView.ResSide.Container)
        }

        if event.Rune() == ':' {
            if app.GetFocus() == nil {
                app.SetFocus(ui.CmdLine)
            }
        }
        return event
    })
}
