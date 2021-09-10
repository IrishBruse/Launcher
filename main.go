//go:generate goversioninfo
package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/zserge/lorca"
)

var downloadsFolder string
var cacheFile string
var settingsFile string

var webview lorca.UI

func main() {
	Appdata, _ := os.UserConfigDir()

	downloadsFolder = path.Join(Appdata, "IrishBruse", "Launcher")
	cacheFile = path.Join(downloadsFolder, "DropboxCache.json")
	settingsFile = path.Join(downloadsFolder, "Settings.json")

	// TODO: make an option in settings to change downloadsFolder

	webview, _ = lorca.New("", "", 1280, 720)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		webview.Close()
		os.Exit(1)
	}()

	initializeGui(webview, onWebviewLoaded)
}

func onWebviewLoaded() {
	webview.Bind("Download", Download)
	webview.Bind("Play", Play)

	var data []byte

	data, err := os.ReadFile(cacheFile)

	contents, _ := os.ReadDir(downloadsFolder)
	// Get what games are downloaded and passs to ui
	for _, game := range contents {
		if game.IsDir() {
			versions, _ := os.ReadDir(path.Join(downloadsFolder, game.Name()))
			for _, version := range versions {
				if version.IsDir() {
					webview.Eval(fmt.Sprintf(
						`
						if(DownloadedVersions["%s"]==null)
						{
							DownloadedVersions["%s"]=[];
						}
						DownloadedVersions["%s"].push("%s");
						`,
						game.Name(),
						game.Name(),
						game.Name(),
						version.Name(),
					))
				}
			}
		}
	}

	// Has Cache
	if err == nil {
		// Pass Cache to ui
		webview.Eval(fmt.Sprintf("DownloadMetadata=%s", data))
		time.Sleep(250 * time.Millisecond)
		webview.Eval("window.dispatchEvent(DownloadMetadataEvent)")
	}

	// Download new version and once done pass that to ui to update
	data = DownloadMetadata()
	webview.Eval(fmt.Sprintf("DownloadMetadata=%s", data))
	time.Sleep(250 * time.Millisecond)
	webview.Eval("window.dispatchEvent(DownloadMetadataEvent)")
}

func Download(file string) {
	versionFolder := path.Join(strings.Replace(downloadsFolder+file, ".zip", "/", 1))

	log.Println(versionFolder)

	err := downloadFromUrl(file)
	if err != nil {
		log.Fatal(err)
	}

	Unzip(downloadsFolder+file, versionFolder)

	os.Remove(downloadsFolder + file)

	gamePaths := strings.Split(versionFolder, "/")

	webview.Eval(fmt.Sprintf(
		`
			if(DownloadedVersions["%s"]==null)
			{
				DownloadedVersions["%s"]=[];
			}
			DownloadedVersions["%s"].push("%s");
			`,
		gamePaths[len(gamePaths)-2],
		gamePaths[len(gamePaths)-2],
		gamePaths[len(gamePaths)-2],
		gamePaths[len(gamePaths)-1],
	))
}

func Play(folder string) {
	pattern := path.Join(downloadsFolder, folder, "/*.exe")
	executables, _ := filepath.Glob(pattern)

	log.Println(executables[0])

	webview.SetBounds(lorca.Bounds{WindowState: lorca.WindowStateMinimized})

	app := exec.Command(executables[0])
	app.Run()

	webview.SetBounds(lorca.Bounds{WindowState: lorca.WindowStateNormal})
}

func DownloadMetadata() []byte {
	projectsMetadata := getProjectsMetadata()
	data, _ := json.Marshal(projectsMetadata)
	os.MkdirAll(path.Dir(cacheFile), fs.ModeDir)
	os.Create(cacheFile)
	os.WriteFile(cacheFile, data, fs.ModeAppend)

	return data
}

// Unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}
