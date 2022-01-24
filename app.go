package main

import (
	"context"
	"fmt"
)

// LauncherApp struct
type LauncherApp struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *LauncherApp {
	return &LauncherApp{}
}

// startup is called at application startup
func (b *LauncherApp) startup(ctx context.Context) {
	b.ctx = ctx

	// i := dropboxFetchIcons(ctx)
}

// domReady is called after the front-end dom has been loaded
func (b *LauncherApp) domReady(ctx context.Context) {
}

// shutdown is called at application termination
func (b *LauncherApp) shutdown(ctx context.Context) {
}

// Greet returns a greeting for the given name
func (b *LauncherApp) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetApps returns an array of icon urls from dropbox
func (b *LauncherApp) GetApps() []App {
	return dropboxFetchIcons(b.ctx)
}
