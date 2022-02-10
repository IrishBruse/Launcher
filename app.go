package main

import (
	"context"
	"encoding/json"
	"io"
	"io/fs"
	"os"
	"path"
	"sync"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"github.com/mitchellh/ioprogress"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var downloadsFolder string

// Launcher struct
type Launcher struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *Launcher {
	return &Launcher{}
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

// Download url
func (b *Launcher) Download(url string) {
	runtime.LogInfo(b.ctx, "test")

	config := dropbox.Config{
		Token:    dropboxToken,
		LogLevel: dropbox.LogDebug,
	}

	dbx := files.New(config)
	downloadArg := files.NewDownloadArg(url)

	res, reader, err := dbx.Download(downloadArg)

	if err != nil {
		runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Dropbox (dbx.Download) Error!",
			Message: err.Error(),
		})
		runtime.LogError(b.ctx, err.Error())
	}

	defer reader.Close()

	os.MkdirAll(downloadsFolder+path.Dir(url), fs.ModeDir)

	f, _ := os.Create(downloadsFolder + url)
	defer f.Close()

	progressbar := &ioprogress.Reader{
		Reader: reader,
		Size:   int64(res.Size),
		DrawFunc: ioprogress.DrawTerminalf(os.Stderr, func(i1, i2 int64) string {
			percent := float32(i1) / float32(i2) * 95.0
			runtime.LogInfo(b.ctx, "test")
			runtime.EventsEmit(b.ctx, "downloadProgress", percent)
			return ""
		}),
	}

	io.Copy(f, progressbar)
}

// GetApps returns an array of icon urls from dropbox
func (b *Launcher) GetApps() string {
	apps := make([]App, 0, 10)

	var wg sync.WaitGroup

	dbxinit()

	appNames := dropboxGetApps(b.ctx)

	for i := 0; i < len(appNames); i++ {
		apps = append(apps, App{Name: appNames[i]})
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

	b2, err := json.Marshal(apps)
	if err != nil {
		runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Dropbox (json.Marshal(apps)) Error!",
			Message: err.Error(),
		})
		runtime.LogError(b.ctx, err.Error())
	}

	return string(b2)
}
