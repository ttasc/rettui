package app

import (
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

func (app *App) Run() error {
    return app.tviewApp.Run()
}
