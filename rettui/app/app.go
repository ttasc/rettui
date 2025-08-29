package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/ttasc/rettui/rettui/ui"
)

type App struct {
    tviewApp    *tview.Application
    ui          *ui.UI
}

func NewApp() *App {

    app := tview.NewApplication().EnableMouse(true)

    ui := ui.NewUI(app)

    app.SetRoot(ui.Root, true)
    app.SetFocus(nil)

    setKeyBinds(app, ui)

    return &App{app, ui}
}

func setKeyBinds(app *tview.Application, ui *ui.UI) {
    app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Key() {
            case tcell.KeyEscape: app.SetFocus(nil)

            case tcell.KeyCtrlJ: // TODO: Tab next
            case tcell.KeyCtrlK: // TODO: Tab previous

            case tcell.KeyCtrlP: app.SetFocus(ui.MainView.ReqSide.Params)
            case tcell.KeyCtrlH: app.SetFocus(ui.MainView.ReqSide.Headers)
            case tcell.KeyCtrlB: app.SetFocus(ui.MainView.ReqSide.Body)
            case tcell.KeyCtrlT: app.SetFocus(ui.MainView.ReqSide.BodyTypes)
            case tcell.KeyCtrlO: app.SetFocus(ui.MainView.ResSide.Container)
        }

        if event.Rune() == ':' {
            if app.GetFocus() == nil {
                app.SetFocus(ui.CmdLine)
            }
        }
        return event
    })
}

func (app *App) Run() error {
    return app.tviewApp.Run()
}
