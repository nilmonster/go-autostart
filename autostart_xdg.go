//go:build linux

package autostart

import (
	"os"
	"path/filepath"
	"text/template"
)

const desktopTemplate = `[Desktop Entry]
Type=Application
Name={{.DisplayName}}
Comment=Autostart service for {{.DisplayName}}
Exec={{.Exec}}
{{- if .Icon}}
Icon={{.Icon}}{{end}}
X-GNOME-Autostart-enabled=true
`

var autostartDir string

func init() {
	if os.Getenv("XDG_CONFIG_HOME") != "" {
		autostartDir = os.Getenv("XDG_CONFIG_HOME")
	} else {
		autostartDir = filepath.Join(os.Getenv("HOME"), ".config")
	}
	autostartDir = filepath.Join(autostartDir, "autostart")
}

func (a *app) path() string {
	return filepath.Join(autostartDir, a.Name+".desktop")
}

// Check if the app is enabled on startup.
func (a *app) IsEnabled() bool {
	_, err := os.Stat(a.path())
	return err == nil
}

// Enable this app on startup.
func (a *app) Enable() error {
	t := template.Must(template.New("desktop").Parse(desktopTemplate))

	if err := os.MkdirAll(autostartDir, 0777); err != nil {
		return err
	}
	f, err := os.Create(a.path())
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, a)
}

// Disable this app on startup.
func (a *app) Disable() error {
	return os.Remove(a.path())
}
