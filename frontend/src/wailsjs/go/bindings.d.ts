export interface go {
  "main": {
    "Launcher": {
		GetApps():Promise<string>
		Greet(arg1:string):Promise<string>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
