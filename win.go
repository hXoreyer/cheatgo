package cheatgo

import "golang.org/x/sys/windows"

var (
	user32                     = windows.NewLazySystemDLL("user32.dll")
	gdi32                      = windows.NewLazySystemDLL("gdi32.dll")
	getSystemMetrics           = user32.NewProc("GetSystemMetrics")
	setLayeredWindowAttributes = user32.NewProc("SetLayeredWindowAttributes")
	showCursor                 = user32.NewProc("ShowCursor")
	setTextAlign               = gdi32.NewProc("SetTextAlign")
	createFont                 = gdi32.NewProc("CreateFontW")
	createCompatibleDC         = gdi32.NewProc("CreateCompatibleDC")
	createSolidBrush           = gdi32.NewProc("CreateSolidBrush")
	createPen                  = gdi32.NewProc("CreatePen")
)
