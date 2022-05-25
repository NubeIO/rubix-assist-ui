

## About

## install wails
```
go install github.com/wailsapp/wails/v2/cmd/wails@latest
wails doctor
```

## Live Development

To run in live development mode, run `wails dev` in the project directory. In another terminal, go into the `frontend`
directory and run `npm run dev`. Navigate to http://localhost:34115
in your browser to connect to your application.



## Building

To build this project in debug mode, use `wails build`. For production, use ` wails build -platform windows/amd64`.


## Known Issues

