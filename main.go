package main

import (
  _ "embed"
  "fmt"
  "github.com/wailsapp/wails"
)

func basic(name string) string {
  return fmt.Sprintf("Hello %v!", name)
}

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 768,
    Title:  "Wails Test",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(basic)
  app.Run()
}
