package main

import (
	"context"
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
func (b *Launcher) GetApps() []App {
	runtime.LogDebug(b.ctx, "test")
	apps := make([]App, 0)

	var wg sync.WaitGroup

	var links []string

	dbxinit()

	wg.Add(1)
	go func() {
		links = dropboxFetchIcons(b.ctx)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		links = dropboxFetchIcons(b.ctx)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		links = dropboxFetchIcons(b.ctx)
		wg.Done()
	}()

	wg.Wait()

	for _, v := range links {
		runtime.LogDebug(b.ctx, "Test: "+v)
	}

	return apps
}
