package main

import (
	"log"
	"strings"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
)

const dropboxToken = "IlMtswMVJHQAAAAAAAAAAUwdNgm0R6QOWjgXE8pC4i3mUcXtRUFEOhWSp5grHTww"

type projectMetadata struct {
	Name     string
	IconURL  string
	Versions []string
}

func getProjectsMetadata() []projectMetadata {

	config := dropbox.Config{
		Token:    dropboxToken,
		LogLevel: dropbox.LogDebug,
	}

	dbx := files.New(config)
	listFolderArg := files.NewListFolderArg("")
	listFolderArg.Recursive = false
	resp, _ := dbx.ListFolder(listFolderArg)

	projects := make([]projectMetadata, 0)

	for _, entry := range resp.Entries {

		switch f := entry.(type) {
		case *files.FolderMetadata:
			projectName := f.PathDisplay[1:]
			projects = append(projects, getProjectMetadata(projectName, dbx))
		}

	}

	return projects
}

func getProjectMetadata(projectName string, dbx files.Client) projectMetadata {
	project := projectMetadata{}
	project.Name = projectName

	listFolderArg := files.NewListFolderArg("/" + projectName)
	listFolderArg.Recursive = false
	resp, _ := dbx.ListFolder(listFolderArg)

	versions := make([]string, 0)

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
				version := strings.ReplaceAll(f.Name, ".zip", "")
				versions = append(versions, version)
			}
		}
	}

	project.Versions = versions

	return project
}

// case *files.FileMetadata:
// 	var getTempLink = files.NewGetTemporaryLinkArg(f.PathLower)
// 	if strings.Contains(f.PathLower, "Icon.png") {

// 		res, err := dbx.GetTemporaryLink(getTempLink)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		var projectName = path.Base(f.PathLower)
// 		t := projects[projectName]
// 		t.iconURL = res.Link
// 		projects[projectName] = t
// 	} else if strings.Contains(f.PathLower, ".zip") {

// 		// fmt.Println(get)
// 		// fmt.Println(res.Link)
// 	}
