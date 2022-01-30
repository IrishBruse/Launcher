package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

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

	// i := dropboxFetchIcons(ctx)
}

// domReady is called after the front-end dom has been loaded
func (b *Launcher) domReady(ctx context.Context) {
}

// shutdown is called at application termination
func (b *Launcher) shutdown(ctx context.Context) {
}

// Greet returns a greeting for the given name
func (b *Launcher) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
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
