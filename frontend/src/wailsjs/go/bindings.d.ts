export interface go {
  "main": {
    "Launcher": {
		Delete(arg1:string):Promise<void>
		Download(arg1:string):Promise<void>
		GetApps():Promise<string>
		Play(arg1:string):Promise<void>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
