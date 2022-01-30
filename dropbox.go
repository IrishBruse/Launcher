package main

import (
	"context"
	"path"
	"strings"
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
	Name     string
	IconURL  string
	Versions []string
}

var dbx files.Client

func dbxinit() {
	config := dropbox.Config{
		Token:    dropboxToken,
		LogLevel: dropbox.LogDebug,
	}
	dbx = files.New(config)
}

func dropboxFetchIcons(ctx context.Context, apps []App) {
	arg := files.NewSearchV2Arg("*/*icon.png")

	res, err := dbx.SearchV2(arg)
	if err != nil {
		runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Dropbox (SearchV2 */*icon.png) Error!",
			Message: err.Error(),
		})
		runtime.LogError(ctx, err.Error())
	}

	var wg sync.WaitGroup

	for _, v := range res.Matches {
		switch t := v.Metadata.Metadata.(type) {
		case *files.FileMetadata:
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				res, err := dbx.GetTemporaryLink(files.NewGetTemporaryLinkArg(t.PathLower))
				if err != nil {
					runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
						Type:    runtime.ErrorDialog,
						Title:   "Dropbox (GetTemporaryLink) Error!",
						Message: err.Error(),
					})
					runtime.LogError(ctx, err.Error())
				}

				for i := 0; i < len(apps); i++ {
					if strings.HasPrefix(t.PathDisplay[1:], apps[i].Name) {
						apps[i].IconURL = res.Link
						break
					}
				}
				wg.Done()
			}(&wg)

		}
	}

	wg.Wait()
}

func dropboxFetchVersions(ctx context.Context, apps []App) {
	arg := files.NewSearchV2Arg("*/*.zip")

	res, err := dbx.SearchV2(arg)
	if err != nil {
		runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Dropbox (SearchV2 */*.zip) Error!",
			Message: err.Error(),
		})
		runtime.LogError(ctx, err.Error())
	}

	for _, v := range res.Matches {
		switch t := v.Metadata.Metadata.(type) {
		case *files.FileMetadata:
			for i := 0; i < len(apps); i++ {
				if strings.HasPrefix(t.PathDisplay[1:], apps[i].Name) {
					file := path.Base(t.PathLower)
					apps[i].Versions = append(apps[i].Versions, strings.TrimSuffix(file, path.Ext(file)))
					break
				}
			}
		}
	}
}

func dropboxGetApps(ctx context.Context) []string {
	appNames := make([]string, 0, 10)

	arg := files.NewListFolderArg("")

	res, err := dbx.ListFolder(arg)
	if err != nil {
		runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Dropbox (ListFolder) Error!",
			Message: err.Error(),
		})
		runtime.LogError(ctx, err.Error())
	}

	for _, v := range res.Entries {
		switch f := v.(type) {
		case *files.FolderMetadata:
			appNames = append(appNames, f.PathDisplay[1:])
		}
	}

	return appNames
}
