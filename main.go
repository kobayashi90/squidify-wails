package main

import (
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// OnStartup is called when the app starts up and is the best place to
// initialize any services you require. The context provided is the
// context for the lifetime of the application.
func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx
}

// GetWebURL returns the Squidify URL
func (a *App) GetWebURL() string {
	return "https://www.squidify.org"
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Squidify",
		Width:  1200,
		Height: 800,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		RGBA:              &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:         app.OnStartup,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
