package main

import (
	"embed"

	"github.com/dqfan2012/luaplayground/internal/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := app.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "Lua Playground",
		Width:         1024,
		Height:        768,
		MinWidth:      400,
		MinHeight:     300,
		DisableResize: false,
		Fullscreen:    false,
		StartHidden:   false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Bind: []interface{}{
			app,
		},
		OnStartup: app.StartUp,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
