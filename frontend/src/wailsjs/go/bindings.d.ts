export interface go {
  "main": {
    "App": {
		Greet(arg1:string):Promise<string>
		MsgFromUI():Promise<void>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
