package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"github.com/mitchellh/ioprogress"
	"github.com/zserge/lorca"
)

//go:embed www/*
var gui embed.FS

func main() {
	webview, _ := lorca.New("", "", 1280, 720)
	defer webview.Close()

	webview.Bind("start", func() {
		log.Println("UI is ready")
	})

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	// Website data
	fsWWW, _ := fs.Sub(gui, "www")
	rootHandler := http.FileServer(http.FS(fsWWW))
	http.Handle("/", rootHandler)

	// Game downloads data
	games := os.DirFS("./games/")
	gamesHandler := http.FileServer(http.FS(games))
	handler := http.StripPrefix("/games/", gamesHandler)
	http.Handle("/games/", handler)

	go http.Serve(ln, nil)

	webview.Load(fmt.Sprintf("http://%s/", ln.Addr()))

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-webview.Done():
	}

	log.Println("exiting...")
}

func dropboxDownload() {

	config := dropbox.Config{
		Token:    "IlMtswMVJHQAAAAAAAAAAUwdNgm0R6QOWjgXE8pC4i3mUcXtRUFEOhWSp5grHTww",
		LogLevel: dropbox.LogDebug,
	}

	dbx := files.New(config)
	var listFolderArg = files.NewListFolderArg("")
	listFolderArg.Recursive = true
	resp, _ := dbx.ListFolder(listFolderArg)

	for _, entry := range resp.Entries {

		switch f := entry.(type) {
		case *files.FolderMetadata:
			err := os.MkdirAll("Games/"+f.PathDisplay, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}

		case *files.FileMetadata:
			var fileArg = files.NewDownloadArg(f.PathLower)

			res, content, err := dbx.Download(fileArg)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Games/" + f.PathDisplay)
			file, err2 := os.Create("Games/" + f.PathDisplay)
			if err2 != nil {
				log.Fatal(err2)
			}
			defer file.Close()

			progressbar := &ioprogress.Reader{
				Reader: content,
				DrawFunc: ioprogress.DrawTerminalf(os.Stderr, func(progress, total int64) string {
					return fmt.Sprintf("Downloading %.2f/%.2f", float32(progress)/float32(total)*100, 100.0)
				}),
				Size: int64(res.Size),
			}

			_, err = io.Copy(file, progressbar)
		}

	}
}
