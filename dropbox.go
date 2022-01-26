package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const dropboxTokenA = "8TduvHwqR_0AAAAAAAAAAdAZ72VyHDLp"
const dropboxTokenB = "PvmI6Ba4YOJdMVUL_FHf85hQyU9FXFle"
const dropboxToken = dropboxTokenA + dropboxTokenB

// App is an app entry on dropbox
type App struct {
	name     string
	iconURL  string
	versions []string
}

var dbx files.Client

func dbxinit() {
	config := dropbox.Config{
		Token:    dropboxToken,
		LogLevel: dropbox.LogDebug,
	}
	dbx = files.New(config)
}

func dropboxFetchIcons(ctx context.Context) []string {

	arg := files.NewSearchV2Arg("*/*icon.png")

	res, err := dbx.SearchV2(arg)

	if err != nil {
		runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Dropbox (SearchV2) Error!",
			Message: err.Error(),
		})
		runtime.LogError(ctx, err.Error())
	}

	iconUrls := getIcons(ctx, res, dbx)

	runtime.LogInfo(ctx, fmt.Sprint(len(iconUrls)))
	for _, v := range iconUrls {
		runtime.LogInfo(ctx, v)
	}

	return iconUrls
}

func dropboxGetApps() []string {
	// arg := files.NewListFolderArg("*/*icon.png")

	// res, err := dbx.SearchV2(arg)

	return nil
}

// getIcons
func getIcons(ctx context.Context, res *files.SearchV2Result, dbx files.Client) []string {
	iconUrls := make([]string, 0)
	var wg sync.WaitGroup
	var m sync.Mutex

	for _, v := range res.Matches {
		switch t := v.Metadata.Metadata.(type) {
		case *files.FileMetadata:
			wg.Add(1)
			go func(mut *sync.Mutex, wg *sync.WaitGroup) {
				defer wg.Done()
				gtlr, err2 := dbx.GetTemporaryLink(files.NewGetTemporaryLinkArg(t.PathLower))
				if err2 != nil {
					runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
						Type:    runtime.ErrorDialog,
						Title:   "Dropbox (GetTemporaryLink) Error!",
						Message: err2.Error(),
					})
					runtime.LogError(ctx, err2.Error())
				}
				m.Lock()
				iconUrls = append(iconUrls, gtlr.Link)
				m.Unlock()
			}(&m, &wg)
		}
	}

	wg.Wait()
	return iconUrls
}
