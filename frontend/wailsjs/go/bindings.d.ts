interface go {
  "main": {
    "App": {
		Greet(arg1:string):Promise<string>
		ShowDialog():Promise<void>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
