package main

import (
	"context"
	"encoding/json"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var downloadsFolder string
var launcher *Launcher
var downloadPercent int

// Launcher struct
type Launcher struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *Launcher {
	launcher = &Launcher{}
	return launcher
}

// startup is called at application startup
func (b *Launcher) startup(ctx context.Context) {
	b.ctx = ctx

	Appdata, _ := os.UserConfigDir()
	downloadsFolder = path.Join(Appdata, "IrishBruse", "Launcher")
}

// domReady is called after the front-end dom has been loaded
func (b *Launcher) domReady(ctx context.Context) {
}

// shutdown is called at application termination
func (b *Launcher) shutdown(ctx context.Context) {
}

// PassThru test
type PassThru struct {
	io.Reader
	total    int64 // Total # of bytes transferred
	length   int64 // Expected length
	progress float64
}

// Read 'overrides' the underlying io.Reader's Read method.
// This is the one that will be called by io.Copy(). We simply
// use it to keep track of byte counts and then forward the call.
func (pt *PassThru) Read(p []byte) (int, error) {
	n, err := pt.Reader.Read(p)
	if n > 0 {
		pt.total += int64(n)
		percentage := float64(pt.total) / float64(pt.length) * float64(95)
		downloadPercent = int(percentage)
		runtime.EventsEmit(launcher.ctx, "downloadProgress", downloadPercent)
	}

	return n, err
}

// Play exe
func (b *Launcher) Play(folder string) {
	runtime.WindowMinimise(b.ctx)
	pattern := path.Join(downloadsFolder, folder, "/*.exe")
	executables, err := filepath.Glob(pattern)
	if err != nil {
		runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Download (dbx.Download) Error!",
			Message: err.Error(),
		})
		runtime.LogError(b.ctx, err.Error())
		return
	}

	app := exec.Command(executables[0])
	app.Dir = path.Join(downloadsFolder, folder)
	app.Run()

	runtime.WindowUnminimise(b.ctx)
}

// Download url
func (b *Launcher) Download(file string) {
	runtime.LogInfo(b.ctx, file)

	config := dropbox.Config{
		Token:    dropboxToken,
		LogLevel: dropbox.LogDebug,
	}

	dbx := files.New(config)
	downloadArg := files.NewDownloadArg(file)

	res, reader, err := dbx.Download(downloadArg)
	if err != nil {
		runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Download (dbx.Download) Error!",
			Message: err.Error(),
		})
		runtime.LogError(b.ctx, err.Error())
		return
	}

	defer reader.Close()

	os.MkdirAll(downloadsFolder+path.Dir(file), fs.ModeDir)

	readerpt := &PassThru{Reader: reader, length: int64(res.Size)}
	data, err := ioutil.ReadAll(readerpt)
	if err != nil {
		runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Download (ioutil.ReadAll) Error!",
			Message: err.Error(),
		})
		runtime.LogError(b.ctx, err.Error())
		return
	}

	os.WriteFile(downloadsFolder+file, data, 0644)

	versionFolder := path.Join(downloadsFolder, strings.Replace(file, ".zip", "", 1))
	Unzip(downloadsFolder+file, versionFolder)
	os.Remove(downloadsFolder + file)

	runtime.EventsEmit(b.ctx, "downloadProgress", 100)
}

// GetApps returns an array of icon urls from dropbox
func (b *Launcher) GetApps() string {
	apps := make([]ListItem, 0, 10)

	var wg sync.WaitGroup

	dbxinit()

	appNames := dropboxGetApps(b.ctx)

	for i := 0; i < len(appNames); i++ {
		apps = append(apps, ListItem{Name: appNames[i]})
		apps[i].Downloaded = make([]string, 0)
	}

	wg.Add(1)
	go func() {
		dropboxFetchIcons(b.ctx, apps)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		dropboxFetchVersions(b.ctx, apps)
		wg.Done()
	}()

	wg.Wait()

	downloads, err := os.ReadDir(downloadsFolder)
	if err != nil {
		runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Download (os.ReadDir(downloadsFolder)) Error!",
			Message: err.Error(),
		})
		runtime.LogError(b.ctx, err.Error())
		return ""
	}

	for _, d := range downloads {
		if d.IsDir() {
			versions, err := os.ReadDir(path.Join(downloadsFolder, d.Name()))
			if err != nil {
				runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
					Type:    runtime.ErrorDialog,
					Title:   "Download (os.ReadDir(path.Join(downloadsFolder, app.Name()))) Error!",
					Message: err.Error(),
				})
				runtime.LogError(b.ctx, err.Error())
				return ""
			}

			for _, version := range versions {
				for i := 0; i < len(apps); i++ {
					if apps[i].Name == d.Name() {
						apps[i].Downloaded = append(apps[i].Downloaded, version.Name())
						break
					}
				}
			}
		}
	}

	b2, err := json.Marshal(apps)
	if err != nil {
		runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Dropbox (json.Marshal(apps)) Error!",
			Message: err.Error(),
		})
		runtime.LogError(b.ctx, err.Error())
		return ""
	}

	// runtime.LogInfo(b.ctx, string(b2))

	return string(b2)
}
