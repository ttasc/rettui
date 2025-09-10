package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
    mainElements []tview.Primitive
    currentFocus int
)

func initalize(app *App) {
    mainElements = []tview.Primitive{
        app.UI.MainView.ReqSide.Params,
        app.UI.MainView.ReqSide.Headers,
        app.UI.MainView.ReqSide.BodyTypes,
        app.UI.MainView.ReqSide.Body,
        app.UI.MainView.ResSide,
    }
}

func (app *App) LoadKeyBinds() {
    initalize(app)

    app.UI.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Key() {
            case tcell.KeyEscape: app.UI.SetFocus(nil)

            case tcell.KeyCtrlJ:
                currentFocus = (currentFocus+1)%len(mainElements)
                app.UI.SetFocus(mainElements[currentFocus])
            case tcell.KeyCtrlK:
                currentFocus = (currentFocus-1+len(mainElements))%len(mainElements)
                app.UI.SetFocus(mainElements[currentFocus])
        }

        switch event.Rune() {
        case ':', '/':
            if app.UI.MainView.ReqSide.HasFocus() {
                break
            }
            app.Handle_Command()
        }

        return event
    })
}
