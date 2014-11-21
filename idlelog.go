package main

import (
    "log"
    "fmt"
    "github.com/BurntSushi/xgb"
    "github.com/BurntSushi/xgb/xproto"
    "github.com/BurntSushi/xgb/screensaver"
)

func main() {
    X, err := xgb.NewConn()
    screen := xproto.Setup(X).DefaultScreen(X)
    if err != nil {
        log.Fatal(err)
    }

    err = screensaver.Init(X)
    if err != nil {
        log.Fatal(err)
    }

    info, err := screensaver.QueryInfo(X, xproto.Drawable(screen.Root)).Reply()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Idle: %dms\n", info.MsSinceUserInput)
}
