package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (init *App) startup(ctx context.Context) {
	// Perform your setup here
	init.ctx = ctx
	go init.MsgFromUI()
	//send time/date to UI
	go init.sendTimeToUI(init.ctx)

}

// domReady is called after the front-end dom has been loaded
func (init App) domReady(ctx context.Context) {

}

// shutdown is called at application termination
func (init *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (init *App) Greet(name string) string {
	fmt.Println("name", name)
	return fmt.Sprintf("Hello %s!", name)
}

func MsgToUI(ctx context.Context, topic string, data interface{}) {
	if ctx != nil {
		runtime.EventsEmit(ctx, topic, data)
	}
}

func (init *App) MsgFromUI() {
	if init.ctx != nil {
		runtime.EventsOn(init.ctx, "terminal-echo", func(optionalData ...interface{}) {
			fmt.Println("Event from UI to backend data: ", optionalData)
		})
	}
}
