package main

import (
	"github.com/ttasc/rettui/rettui/app"
)

func main() {
    app := app.NewApp()
    if err := app.Run(); err != nil {
        panic(err)
    }
}
