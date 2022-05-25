package main

import (
	"log"
	"os"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func main() {
	// Initialize astilectron
	a, err := astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName:            "SofaLab",
		AppIconDefaultPath: "resources/icon.png",
		AppIconDarwinPath:  "resources/icon.icns",
		BaseDirectoryPath:  "electron/",
		VersionAstilectron: "0.33.0",
		VersionElectron:    "13.0.0",
	})
	if err != nil {
		panic(err)
	}

	defer a.Close()

	// Start astilectron
	err = a.Start()
	if err != nil {
		panic(err)
	}

	// Create a new window
	w, err := a.NewWindow("http://localhost:4200", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(1024),
		Width:  astikit.IntPtr(768),
	})
	if err != nil {
		panic(err)
	}
	err = w.Create()
	if err != nil {
		panic(err)
	}

	// Blocking pattern
	a.Wait()
}
