package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"github.com/mitchellh/ioprogress"
)

const dropboxToken = "IlMtswMVJHQAAAAAAAAAAUwdNgm0R6QOWjgXE8pC4i3mUcXtRUFEOhWSp5grHTww"

type VersionMetadata struct {
	Version string
	Url     string
}

type ProjectMetadata struct {
	Name     string
	IconURL  string
	Versions []VersionMetadata
}

func getProjectsMetadata() []ProjectMetadata {

	config := dropbox.Config{
		Token:    dropboxToken,
		LogLevel: dropbox.LogDebug,
	}

	dbx := files.New(config)
	listFolderArg := files.NewListFolderArg("")
	listFolderArg.Recursive = false
	resp, _ := dbx.ListFolder(listFolderArg)

	projects := make([]ProjectMetadata, 0)

	for _, entry := range resp.Entries {

		switch f := entry.(type) {
		case *files.FolderMetadata:
			projectName := f.PathDisplay[1:]
			projects = append(projects, getProjectMetadata(projectName, dbx))
		}

	}

	return projects
}

func getProjectMetadata(projectName string, dbx files.Client) ProjectMetadata {
	project := ProjectMetadata{}
	project.Name = projectName

	listFolderArg := files.NewListFolderArg("/" + projectName)
	listFolderArg.Recursive = false
	resp, _ := dbx.ListFolder(listFolderArg)

	versions := make([]VersionMetadata, 0)

	for _, entry := range resp.Entries {

		switch f := entry.(type) {
		case *files.FileMetadata:
			if strings.Contains(f.PathLower, "icon.png") {
				getTempLink := files.NewGetTemporaryLinkArg(f.PathLower)
				res, err := dbx.GetTemporaryLink(getTempLink)
				if err != nil {
					log.Fatal(err)
				}

				project.IconURL = res.Link
			} else if strings.Contains(f.PathLower, ".zip") {
				version := VersionMetadata{}
				version.Version = strings.ReplaceAll(f.Name, ".zip", "")
				version.Url = "/" + projectName + "/" + f.Name
				versions = append(versions, version)
			}
		}
	}

	project.Versions = versions

	return project
}

func downloadFromUrl(url string) error {
	config := dropbox.Config{
		Token:    dropboxToken,
		LogLevel: dropbox.LogDebug,
	}

	dbx := files.New(config)
	downloadArg := files.NewDownloadArg(url)

	res, reader, err := dbx.Download(downloadArg)

	if err != nil {
		return err
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

			webview.Eval(fmt.Sprintf("DownloadProgress=%.2f", percent))
			webview.Eval("window.dispatchEvent(DownloadProgressEvent)")

			return ""
		}),
	}

	io.Copy(f, progressbar)

	return nil
}
