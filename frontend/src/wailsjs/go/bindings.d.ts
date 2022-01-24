export interface go {
  "main": {
    "LauncherApp": {
		GetApps():Promise<Array<App>>
		Greet(arg1:string):Promise<string>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
