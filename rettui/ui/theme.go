package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var Colors = map[string]tcell.Color{
    "white"     : tcell.NewHexColor(0xc8d3f5), // #c8d3f5
    "dark"      : tcell.NewHexColor(0x222436), // #222436
    "darker"    : tcell.NewHexColor(0x1b1d2b), // #1b1d2b
    "red"       : tcell.NewHexColor(0xff757f), // #ff757f
    "green"     : tcell.NewHexColor(0xc3e88d), // #c3e88d
    "orange"    : tcell.NewHexColor(0xffc777), // #ffc777
    "blue"      : tcell.NewHexColor(0x82aaff), // #82aaff
    "purple"    : tcell.NewHexColor(0xc099ff), // #c099ff
    "cyan"      : tcell.NewHexColor(0x86e1fc), // #86e1fc
    "darkblue"  : tcell.NewHexColor(0x444a73), // #444a73
    "grey"      : tcell.NewHexColor(0x828bb8), // #828bb8
}

func ReDefineTviewVars() {
    redefineColors()
    redefineBorders()
}

func redefineColors() {
    tview.Styles = tview.Theme{
        PrimitiveBackgroundColor:    Colors["dark"],
        ContrastBackgroundColor:     Colors["darkblue"],
        MoreContrastBackgroundColor: Colors["grey"],
        BorderColor:                 Colors["cyan"],
        TitleColor:                  Colors["cyan"],
        GraphicsColor:               Colors["white"],
        PrimaryTextColor:            Colors["white"],
        SecondaryTextColor:          Colors["orange"],
        TertiaryTextColor:           Colors["darkblue"],
        InverseTextColor:            Colors["grey"],
        ContrastSecondaryTextColor:  Colors["blue"],
    }
}

func redefineBorders() {
    tview.Borders = struct {
        Horizontal  rune
        Vertical    rune
        TopLeft     rune
        TopRight    rune
        BottomLeft  rune
        BottomRight rune

        LeftT   rune
        RightT  rune
        TopT    rune
        BottomT rune
        Cross   rune

        HorizontalFocus  rune
        VerticalFocus    rune
        TopLeftFocus     rune
        TopRightFocus    rune
        BottomLeftFocus  rune
        BottomRightFocus rune
    }{
        Horizontal:  tview.BoxDrawingsLightHorizontal,
        Vertical:    tview.BoxDrawingsLightVertical,
        TopLeft:     tview.BoxDrawingsLightDownAndRight,
        TopRight:    tview.BoxDrawingsLightDownAndLeft,
        BottomLeft:  tview.BoxDrawingsLightUpAndRight,
        BottomRight: tview.BoxDrawingsLightUpAndLeft,

        LeftT:   tview.BoxDrawingsLightVerticalAndRight,
        RightT:  tview.BoxDrawingsLightVerticalAndLeft,
        TopT:    tview.BoxDrawingsLightDownAndHorizontal,
        BottomT: tview.BoxDrawingsLightUpAndHorizontal,
        Cross:   tview.BoxDrawingsLightVerticalAndHorizontal,

        HorizontalFocus:  tview.BoxDrawingsLightHorizontal,
        VerticalFocus:    tview.BoxDrawingsLightVertical,
        TopLeftFocus:     tview.BoxDrawingsDoubleDownAndRight,
        TopRightFocus:    tview.BoxDrawingsDoubleDownAndLeft,
        BottomLeftFocus:  tview.BoxDrawingsDoubleUpAndRight,
        BottomRightFocus: tview.BoxDrawingsDoubleUpAndLeft,
    }
}
