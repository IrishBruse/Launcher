export interface go {
  "main": {
    "Launcher": {
		Download(arg1:string):Promise<void>
		GetApps():Promise<string>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
