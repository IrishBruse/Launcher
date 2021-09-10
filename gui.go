package main

import (
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/zserge/lorca"
)

//go:embed embed/*
var gui embed.FS

func initializeGui(webview lorca.UI, callback func()) {
	webview.SetBounds(lorca.Bounds{Left: 320, Top: 180, Width: 1280, Height: 720})

	// This functions should be called once the website is done loading
	// and is called from js
	webview.Bind("start", func() {
		callback()
	})

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	// Website data
	// fsWWW, _ := fs.Sub(gui, "embed")
	//  fsWWW, _ := fs.(gui, "embed")
	rootHandler := http.FileServer(http.Dir("embed/"))
	http.Handle("/", rootHandler)

	// Game downloads data
	games := os.DirFS("./games/")
	gamesHandler := http.FileServer(http.FS(games))
	handler := http.StripPrefix("/games/", gamesHandler)
	http.Handle("/games/", handler)

	go http.Serve(ln, nil)

	webview.Load(fmt.Sprintf("http://%s/", ln.Addr()))

	<-webview.Done()
}
