package main

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

type GreetService struct {
	Window *application.WebviewWindow
}

func (g *GreetService) Greet(name string) string {
	return "Hello " + name + "!"
}

// MinimizeWindow 最小化窗口
func (g *GreetService) MinimizeWindow() {
	if g.Window != nil {
		g.Window.Minimise()
	}
}

// MaximizeWindow 最大化窗口
func (g *GreetService) MaximizeWindow() {
	if g.Window != nil {
		g.Window.Maximise()
	}
}

// CloseWindow 关闭窗口
func (g *GreetService) CloseWindow() {
	if g.Window != nil {
		g.Window.Close()
	}
}
