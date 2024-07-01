package cheatgo

import (
	"github.com/lxn/win"       // 导入WinAPI包
	"golang.org/x/sys/windows" // 导入Windows包
)

// 引入Windows DLL
var (
	user32                     = windows.NewLazySystemDLL("user32.dll")       // user32.dll
	gdi32                      = windows.NewLazySystemDLL("gdi32.dll")        // gdi32.dll
	getSystemMetrics           = user32.NewProc("GetSystemMetrics")           // GetSystemMetrics函数
	setLayeredWindowAttributes = user32.NewProc("SetLayeredWindowAttributes") // SetLayeredWindowAttributes函数
	showCursor                 = user32.NewProc("ShowCursor")                 // ShowCursor函数
	setTextAlign               = gdi32.NewProc("SetTextAlign")                // SetTextAlign函数
	createFont                 = gdi32.NewProc("CreateFontW")                 // CreateFontW函数
	createCompatibleDC         = gdi32.NewProc("CreateCompatibleDC")          // CreateCompatibleDC函数
	createSolidBrush           = gdi32.NewProc("CreateSolidBrush")            // CreateSolidBrush函数
	createPen                  = gdi32.NewProc("CreatePen")                   // CreatePen函数
)

// CreateSolidBrush创建一个实心画刷
func (w *Window) CreateSolidBrushA(color uintptr) (uintptr, uintptr, error) {
	return createSolidBrush.Call(color) // 调用CreateSolidBrush函数
}

// DeleteObject删除一个GDI对象
func (w *Window) DeleteObject(obj uintptr) bool {
	return win.DeleteObject(win.HGDIOBJ(obj)) // 调用DeleteObject函数
}

// CreatePen创建一个画笔
func (w *Window) CreatePen(style int, width int, color RGB) (uintptr, uintptr, error) {
	colr := w.RGBToCOLORREF(color)
	return createPen.Call(uintptr(style), uintptr(width), colr) // 调用CreatePen函数
}

// CreateFont创建一个字体
func (w *Window) CreateFont() (uintptr, uintptr, error) {
	return createFont.Call(
		12,                                // 字体高度
		0,                                 // 字体宽度
		0,                                 // 逃逸方向
		0,                                 // 字体倾斜角度
		win.FW_HEAVY,                      // 字体重量
		0,                                 // 斜体
		0,                                 // 下划线
		0,                                 // 删除线
		win.DEFAULT_CHARSET,               // 字符集
		win.OUT_DEFAULT_PRECIS,            // 输出精度
		win.CLIP_DEFAULT_PRECIS,           // 裁剪精度
		win.DEFAULT_QUALITY,               // 输出质量
		win.DEFAULT_PITCH|win.FF_DONTCARE, // 字体间距和家族
		0,                                 // 字体名称
	)
}

// DeleteDC删除一个设备上下文
func (w *Window) DeleteDC(dc win.HDC) {
	win.DeleteDC(dc) // 调用DeleteDC函数
}

// Rectangle绘制一个矩形
func (w *Window) Rectangle(rect RECT, width int, color RGB, style int, fillColor ...RGB) {
	if len(fillColor) > 0 {
		// 使用实心画刷进行填充
		brush, _, _ := w.CreateSolidBrushA(w.RGBToCOLORREF(fillColor[0]))
		defer w.DeleteObject(brush)

		// 选择画刷
		oldBrush := win.SelectObject(w.hdc, win.HGDIOBJ(brush))
		defer win.SelectObject(w.hdc, win.HGDIOBJ(oldBrush))

		// 绘制填充矩形
		win.Rectangle_(w.hdc, rect.Left, rect.Top, rect.Right, rect.Bottom)
	} else {
		// 使用空心画笔绘制矩形框
		pen, _, _ := w.CreatePen(style, width, color)
		defer w.DeleteObject(pen)
		win.SelectObject(w.hdc, win.HGDIOBJ(pen))
		win.MoveToEx(w.hdc, int(rect.Left), int(rect.Top), nil)
		win.LineTo(w.hdc, rect.Right, rect.Top)
		win.LineTo(w.hdc, rect.Right, rect.Bottom)
		win.LineTo(w.hdc, rect.Left, rect.Bottom)
		win.LineTo(w.hdc, rect.Left, rect.Top)
	}
}

// Circle绘制一个圆形
func (w *Window) Circle(v Vector2, radius int, width int, color RGB, style int) {
	pen, _, _ := w.CreatePen(style, width, color)
	defer w.DeleteObject(pen)
	win.SelectObject(w.hdc, win.HGDIOBJ(pen))                                                                          // 选择画笔
	win.Ellipse(w.hdc, int32(int(v.X)-radius), int32(int(v.Y)-radius), int32(int(v.X)+radius), int32(int(v.Y)+radius)) // 绘制圆形
}

// Circle3绘制一个三维空间的圆形
func (w *Window) Circle3(v Vector3, radius int, width int, color RGB, style int) {
	pen, _, _ := w.CreatePen(style, width, color)
	defer w.DeleteObject(pen)
	win.SelectObject(w.hdc, win.HGDIOBJ(pen))
	win.Ellipse(w.hdc, int32(int(v.X)-radius), int32(int(v.Y)), int32(int(v.X)+radius), int32(int(v.Z)+radius)) // 选择画笔
}

// Text在指定位置绘制文本
func (w *Window) Text(txt string, rgb RGB, v Vector2) {
	text, _ := windows.UTF16PtrFromString(txt)                              // 转换字符串为UTF16指针
	win.SetTextColor(w.hdc, win.RGB(byte(rgb.R), byte(rgb.G), byte(rgb.B))) // 设置文本颜色
	win.TextOut(w.hdc, int32(v.X), int32(v.Y), text, int32(len(txt)))       // 绘制文本
}

// LineTo绘制一条线
func (w *Window) LineTo(v Vector2, width int, color RGB, style int) {
	pen, _, _ := w.CreatePen(style, width, color)
	defer w.DeleteObject(pen)
	win.SelectObject(w.hdc, win.HGDIOBJ(pen))
	win.MoveToEx(w.hdc, int(v.X), int(v.Y), nil)
	win.LineTo(w.hdc, int32(v.X), int32(v.Y))
}

// RGBToCOLORREF 将RGB结构体转换为COLORREF值
func (w *Window) RGBToCOLORREF(rgb RGB) uintptr {
	return uintptr(rgb.R) | (uintptr(rgb.G) << 8) | (uintptr(rgb.B) << 16)
}
