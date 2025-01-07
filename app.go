package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// App struct
type App struct {
	ctx      context.Context
	settings settings
}

// NewApp creates a new App application struct
func NewApp(s settings) *App {
	return &App{
		settings: s,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetListOfImages returns a list of images from the given directories
func (a *App) GetListOfImages() []string {
	var images []string
	for i, dir := range a.settings.ImageDirectories {
		err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if !d.IsDir() {
				// Get the relative path from the base directory
				relPath, err := filepath.Rel(dir, path)
				if err != nil {
					return nil
				}
				images = append(images, filepath.Join(strconv.Itoa(i), relPath))
			}
			return nil
		})
		if err != nil {
			fmt.Println("Error walking directory: ", err.Error())
			return nil
		}
	}
	return images
}
