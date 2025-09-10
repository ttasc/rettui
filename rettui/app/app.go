package app

import (
	"github.com/ttasc/rettui/rettui/ui"
)

type App struct {
    UI          *ui.UI
}

func NewApp() *App {
    UI := ui.NewUI()

    app := &App{UI}

    app.LoadKeyBinds()

    return app
}

func (app *App) Run() error {
    return app.UI.Run()
}

func (app *App) Quit() {
    app.UI.Stop()
}

