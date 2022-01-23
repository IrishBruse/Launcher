package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (b *App) startup(ctx context.Context) {
	b.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (b *App) domReady(ctx context.Context) {
}

// shutdown is called at application termination
func (b *App) shutdown(ctx context.Context) {
}

// Greet returns a greeting for the given name
func (b *App) Greet(name string) string {

	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Test returns a greeting for the given name
func (b *App) Test(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// ShowDialog Shows a Dialog
func (b *App) ShowDialog() {
	_, err := runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   "Native Dialog from Go",
		Message: "This is a Native Dialog send from Go.",
	})

	if err != nil {
		panic(err)
	}
}
