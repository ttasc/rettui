package ui

import (
	"github.com/rivo/tview"
)

type mainView struct {
    ReqSide  *reqSide
    ResSide  *resSide
}

type reqSide struct {
    *tview.Flex
    Params      *tview.TextArea
    Headers     *tview.TextArea
    BodyTypes   *tview.DropDown
    Body        *tview.TextArea
}

type resSide struct {
    *tview.Pages
    Headers     *tview.Table
    Body        *tview.TextView
}

func newMainView() *mainView {
    mainView := &mainView{
        ReqSide:  newRequestSide(),
        ResSide: newResponseSide(),
    }
    return mainView
}

func (m *mainView) addTo(layout *tview.Grid) {
    layout.AddItem(m.ReqSide, 1, 0, 1, 1, 0, 0, true)
    layout.AddItem(m.ResSide, 1, 1, 1, 1, 0, 0, true)
}

func newRequestSide() *reqSide {
    requestSide := &reqSide{
        Flex:        tview.NewFlex().SetDirection(tview.FlexRow),
        Params:      tview.NewTextArea().SetPlaceholder("URL Params...").SetWordWrap(true),
        Headers:     tview.NewTextArea().SetPlaceholder("<key>: <value>\nExample:\n...\nCache-Control: no-cache\nServer: Microsoft-IIS/10.0\nContent-Type: application/json\n...").SetWordWrap(false),
        BodyTypes:   tview.NewDropDown().SetLabel("Body-type ").SetOptions([]string{"none", "form-data", "x-www-form-urlencoded", "raw"}, nil).SetCurrentOption(0),
        Body:        tview.NewTextArea().SetPlaceholder("Insert data here...").SetWordWrap(true),
    }

    requestSide.Params.SetBorder(true).SetTitle("Request Params").SetTitleAlign(tview.AlignLeft)
    requestSide.Headers.SetBorder(true).SetTitle("Request Headers").SetTitleAlign(tview.AlignLeft)
    requestSide.Body.SetBorder(true).SetTitle("Request Body").SetTitleAlign(tview.AlignLeft)

    requestSide.AddItem(requestSide.Params, 4, 1, false)
    requestSide.AddItem(requestSide.Headers, 10, 1, false)
    requestSide.AddItem(requestSide.BodyTypes, 1, 1, false)
    requestSide.AddItem(requestSide.Body, 0, 1, false)
    return requestSide
}

func newResponseSide() *resSide {
    responseSide := &resSide{
        Pages:       tview.NewPages(),
        Headers:     tview.NewTable(),
        Body:        tview.NewTextView().SetDynamicColors(true).SetWrap(false).SetTextColor(Colors["red"]),
    }

    responseSide.Body.SetText(responseTestText)

    responseSide.SetBorder(true).SetTitle("Response - Ctrl+H/Ctrl+B to switch Headers/Body").SetTitleAlign(tview.AlignLeft)
    responseSide.AddPage("Headers", responseSide.Headers, true, false)
    responseSide.AddPage("Body", responseSide.Body, true, true)
    return responseSide
}

var responseTestText = `<!doctype html>
<html>

<head>
    <title>Example Domain</title>

    <meta charset="utf-8" />
    <meta http-equiv="Content-type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style type="text/css">
        body {
            background-color: #f0f0f2;
            margin: 0;
            padding: 0;
            font-family: -apple-system, system-ui, BlinkMacSystemFont, "Segoe UI", "Open Sans", "Helvetica Neue", Helvetica, Arial, sans-serif;

        }

        div {
            width: 600px;
            margin: 5em auto;
            padding: 2em;
            background-color: #fdfdff;
            border-radius: 0.5em;
            box-shadow: 2px 3px 7px 2px rgba(0, 0, 0, 0.02);
        }

        a:link,
        a:visited {
            color: #38488f;
            text-decoration: none;
        }

        @media (max-width: 700px) {
            div {
                margin: 0 auto;
                width: auto;
            }
        }
    </style>
</head>

<body>
    <div>
        <h1>Example Domain</h1>
        <p>This domain is for use in illustrative examples in documents. You may use this
            domain in literature without prior coordination or asking for permission.</p>
        <p><a href="https://www.iana.org/domains/example">More information...</a></p>
    </div>
</body>

</html>
`
