package main

import (
	"embed"
	"fmt"
	"io"
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

// go:embed www
var fs embed.FS

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

	go http.Serve(ln, http.FileServer(http.FS(fs)))
	webview.Load(fmt.Sprintf("http://%s/www", ln.Addr()))

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
		Token:    "sl.A3S_ZRNOoini4aF37yLn1qv8yp5bUE23TV8VHwHNFU2qtVyKrR2l6Yy9 gAL1hqdH8AYhMjeTNPKnG2PzVSLlwviTeyMwEOTiH-VmEu_V14hjX4h1qPvQDJEpv60Mxjp_2fuJRSZp",
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
