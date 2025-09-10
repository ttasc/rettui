package app

import "github.com/gdamore/tcell/v2"

var CmdFuncsLong = map[string]func(app *App){
    "quit": cmd_quit,
}

var CmdFuncsShort = map[string]func(app *App){
    "q": cmd_quit,
}

func (app *App) Handle_Command() {
    app.UI.SetFocus(app.UI.CmdLine)
    app.UI.CmdLine.SetDoneFunc(func(key tcell.Key) {
        cmd := app.UI.CmdLine.GetText()[1:]
        if f, ok := CmdFuncsShort[cmd]; ok {
            f(app)
        } else if f, ok := CmdFuncsLong[cmd]; ok {
            f(app)
        } else {
            app.UI.CmdLine.SetText("Error: command not found")
            app.UI.SetFocus(nil)
        }
    })
}

func cmd_quit (app *App) {
    app.Quit()
}
