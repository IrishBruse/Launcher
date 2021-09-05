package main

import (
	"encoding/json"
	"fmt"

	"github.com/zserge/lorca"
)

func main() {
	initializeGui(onWebviewLoaded)
}

func onWebviewLoaded(webview lorca.UI) {

	projectsMetadata := getProjectsMetadata()
	data, _ := json.Marshal(projectsMetadata)

	jsVariable := fmt.Sprintf("projectsMetadata=%s", data)
	fmt.Printf("jsVariable: %v\n", jsVariable)

	webview.Eval(jsVariable)
	webview.Eval("document.dispatchEvent(downloadedProjectsMetadata)")
}
