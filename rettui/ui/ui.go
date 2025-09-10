package ui

import (
	"github.com/rivo/tview"
)

type UI struct {
    *tview.Application

    Root        *tview.Grid

    TabLine     *tabLine
    MainView    *mainView
    StatusLine  *statusLine
    CmdLine     *cmdLine
}

func init() {
    redefineTviewVars()
}

func NewUI() *UI {
    root := tview.NewGrid()

    app := tview.NewApplication().EnableMouse(true).SetRoot(root, true).SetFocus(nil)

    ui := &UI{
        app,
        root,
        newTabLine(),
        newMainView(),
        newStatusLine(app),
        newCmdLine(app),
    }
    ui.Root.SetRows(
        1,  // tabline
        0,  // main view
        1,  // statusline
        1,  // commandline
    )
    ui.TabLine.addTo(ui.Root)
    ui.MainView.addTo(ui.Root)
    ui.StatusLine.addTo(ui.Root)
    ui.CmdLine.addTo(ui.Root)

    return ui
}

func (ui *UI) Run() error {
    return ui.Application.Run()
}
