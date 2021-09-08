//go:generate goversioninfo
package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"time"

	"github.com/zserge/lorca"
)

const FolderName = "/IrishbruseLauncher/"

var CacheFile string

func main() {
	AppdataLocal, _ := os.UserCacheDir()
	CacheFile = AppdataLocal + FolderName + "Cache.json"

	initializeGui(onWebviewLoaded)
}

func onWebviewLoaded(webview lorca.UI) {

	webview.Bind("Download", Download)

	var data []byte

	data, err := os.ReadFile(CacheFile)

	// Has Cache
	if err == nil {
		// Pass Cache to ui
		webview.Eval(fmt.Sprintf("projectsMetadata=%s", data))
		time.Sleep(250 * time.Millisecond)
		webview.Eval("document.dispatchEvent(downloadedProjectsMetadata)")
	}

	// Download new version and once done pass that to ui to update
	data = DownloadMetadata()
	webview.Eval(fmt.Sprintf("projectsMetadata=%s", data))
	time.Sleep(250 * time.Millisecond)
	webview.Eval("document.dispatchEvent(downloadedProjectsMetadata)")
}

func Download(url string) {
	log.Printf("url: %v\n", url)
}

func DownloadMetadata() []byte {
	projectsMetadata := getProjectsMetadata()
	data, _ := json.Marshal(projectsMetadata)
	os.MkdirAll(path.Dir(CacheFile), fs.ModeDir)
	os.Create(CacheFile)
	os.WriteFile(CacheFile, data, fs.ModeAppend)

	return data
}
