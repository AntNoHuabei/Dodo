package main

import (
	"embed"
	_ "embed"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/services/sqlite"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed all:frontend/dist
var assets embed.FS

// getAppDataDir 获取应用程序数据目录，使用平台API确保可靠性
func getAppDataDir() string {
	switch runtime.GOOS {
	case "windows":
		key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Explorer\Shell Folders`, 0)
		if err != nil {
			userProfile := os.Getenv("USERPROFILE")
			return filepath.Join(userProfile, "AppData", "Roaming", "dodo")
		}
		defer key.Close()

		var value string
		value, _, err = key.GetStringValue("AppData")
		if err != nil || value == "" {
			userProfile := os.Getenv("USERPROFILE")
			return filepath.Join(userProfile, "AppData", "Roaming", "dodo")
		}
		return filepath.Join(value, "dodo")

	case "darwin":
		home := os.Getenv("HOME")
		if home == "" {
			if info, err := os.UserHomeDir(); err == nil {
				home = info
			}
		}
		return filepath.Join(home, "Library", "Application Support", "dodo")

	default:
		xdgConfig := os.Getenv("XDG_CONFIG_HOME")
		if xdgConfig == "" {
			home := os.Getenv("HOME")
			if home == "" {
				if info, err := os.UserHomeDir(); err == nil {
					home = info
				}
			}
			xdgConfig = filepath.Join(home, ".config")
		}
		return filepath.Join(xdgConfig, "dodo")
	}
}

// getLocalAppDataDir 获取本地应用程序数据目录（Windows）
func getLocalAppDataDir() string {
	switch runtime.GOOS {
	case "windows":
		key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Explorer\Shell Folders`, 0)
		if err != nil {
			userProfile := os.Getenv("USERPROFILE")
			return filepath.Join(userProfile, "AppData", "Local", "dodo")
		}
		defer key.Close()

		var value string
		value, _, err = key.GetStringValue("Local AppData")
		if err != nil || value == "" {
			userProfile := os.Getenv("USERPROFILE")
			return filepath.Join(userProfile, "AppData", "Local", "dodo")
		}
		return filepath.Join(value, "dodo")

	case "darwin":
		home := os.Getenv("HOME")
		if home == "" {
			if info, err := os.UserHomeDir(); err == nil {
				home = info
			}
		}
		return filepath.Join(home, "Library", "Application Support", "dodo")

	default:
		xdgData := os.Getenv("XDG_DATA_HOME")
		if xdgData == "" {
			home := os.Getenv("HOME")
			if home == "" {
				if info, err := os.UserHomeDir(); err == nil {
					home = info
				}
			}
			xdgData = filepath.Join(home, ".local", "share")
		}
		return filepath.Join(xdgData, "dodo")
	}
}

// hideFileWindows 仅在Windows上隐藏文件
func hideFileWindows(path string) {
	if runtime.GOOS == "windows" {
		attr, err := windows.GetFileAttributes(windows.StringToUTF16Ptr(path))
		if err == nil {
			windows.SetFileAttributes(windows.StringToUTF16Ptr(path), attr|windows.FILE_ATTRIBUTE_HIDDEN)
		}
	}
}

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {
	appDataPath := getAppDataDir()
	dbPath := filepath.Join(appDataPath, "data.db")

	if err := os.MkdirAll(appDataPath, 0755); err != nil {
		log.Fatalf("Failed to create app data directory: %v", err)
	}

	if runtime.GOOS == "windows" {
		hideFileWindows(appDataPath)
	}

	app := application.New(application.Options{
		Name:        "Dodo",
		Description: "Dodo",
		Services: []application.Service{
			application.NewService(sqlite.NewWithConfig(&sqlite.Config{DBSource: dbPath})),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
		Windows: application.WindowsOptions{
			AdditionalBrowserArgs: []string{"--disable-web-security", "--no-proxy-server", "--disable-features=ElasticOverscroll", "--disable-pinch"},
		},
	})

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Frameless' 设置为 true 以使用自定义标题栏
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:     "Dodo",
		Frameless: true,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
	})

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
