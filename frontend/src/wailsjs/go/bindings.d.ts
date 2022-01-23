interface go {
    "main": {
        "App": {
            Greet(arg1: string): Promise<string>
            ShowDialog(): Promise<void>
            Test(arg1: string): Promise<string>
        },
    }

}

declare global {
    interface Window {
        go: go;
    }
}
