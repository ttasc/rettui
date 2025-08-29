package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type tabLine struct {
    textview *tview.TextView
    tabs     []string
}

func newTabLine() *tabLine {
    tabLine := &tabLine{
        tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetWrap(false),
        []string{"Tab 1", "Tab 2", "Tab 3"},
    }
    tabLine.textview.SetBackgroundColor(tcell.NewHexColor(0x10111a))
    tabLine.update()
    return tabLine
}

func (t *tabLine) addTo(layout *tview.Grid) {
    layout.AddItem(t.textview, 0, 0, 1, 2, 0, 0, false)
}

func (t *tabLine) addToFlex(layout *tview.Flex) {
    layout.AddItem(t.textview, 1, 1, false)
}

func (t *tabLine) update() {
    for index, tab := range t.tabs {
        fmt.Fprintf(t.textview, `["%d"][#222436] %s [white][""]`, index, tab)
    }
    t.textview.Highlight("0")
}

