package autostart

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

// An application that will be started when the user logs in.
type app struct {
	// Unique identifier for the app.
	Name string
	// The command to execute, followed by its arguments.
	Exec string
	// The app name.
	DisplayName string
	// The app icon.
	Icon string
}

func New(name, exec, displayName string, icon []byte) *app {
	iconPath := fmt.Sprintf("%s/%s", os.TempDir(), "appicon.png")
	if _, err := os.Stat(iconPath); os.IsNotExist(err) {
		_, err := byteToPng(icon, iconPath)
		if err != nil {
			log.Printf("icon error: %s\n", err)
			return nil
		}
	}

	return &app{
		Name:        name,
		Exec:        exec,
		DisplayName: displayName,
		Icon:        iconPath,
	}
}

func byteToPng(b []byte, output string) (out *os.File, err error) {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return
	}

	out, _ = os.Create(output)
	defer out.Close()

	err = png.Encode(out, img)
	if err != nil {
		return
	}
	return
}
