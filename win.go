package cheatgo

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/lxn/win"       // 导入WinAPI包
	"golang.org/x/sys/windows" // 导入Windows包
)

type Window struct {
	hwnd win.HWND
	msg  win.MSG
	hdc  win.HDC // 设备上下文句柄
	sw   int
	sh   int
}

func NewWindow(name string, screenWidth, screenHeight uintptr) (*Window, error) {
	hwnd, err := initWindows(name, screenWidth, screenHeight)
	return &Window{
		hwnd: hwnd,
		hdc:  win.GetDC(hwnd),
		sw:   int(screenWidth),
		sh:   int(screenHeight),
	}, err
}

// 窗口过程函数
func windowProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_TIMER: // 定时器消息
		return 0
	case win.WM_DESTROY: // 窗口销毁消息
		win.PostQuitMessage(0) // 发送退出消息
		return 0
	default:
		return win.DefWindowProc(hwnd, msg, wParam, lParam) // 调用默认窗口过程处理其他消息
	}
}

// 初始化窗口函数
func initWindows(name string, screenWidth, screenHeight uintptr) (win.HWND, error) {
	// 将窗口类名转换为UTF16指针
	className, err := windows.UTF16PtrFromString(fmt.Sprintf("%sWindow", name))
	if err != nil {
		return 0, err
	}
	// 将窗口标题转换为UTF16指针
	windowTitle, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return 0, err
	}

	// 注册窗口类
	wc := win.WNDCLASSEX{
		CbSize:        uint32(unsafe.Sizeof(win.WNDCLASSEX{})),                                  // 结构体大小
		Style:         win.CS_HREDRAW | win.CS_VREDRAW,                                          // 窗口样式
		LpfnWndProc:   syscall.NewCallback(windowProc),                                          // 窗口过程函数
		CbWndExtra:    0,                                                                        // 额外窗口内存
		HInstance:     win.GetModuleHandle(nil),                                                 // 当前实例句柄
		HIcon:         win.LoadIcon(0, (*uint16)(unsafe.Pointer(uintptr(win.IDI_APPLICATION)))), // 图标
		HCursor:       win.LoadCursor(0, (*uint16)(unsafe.Pointer(uintptr(win.IDC_ARROW)))),     // 光标
		HbrBackground: win.COLOR_WINDOW,                                                         // 背景颜色
		LpszMenuName:  nil,                                                                      // 菜单名
		LpszClassName: className,                                                                // 窗口类名
		HIconSm:       win.LoadIcon(0, (*uint16)(unsafe.Pointer(uintptr(win.IDI_APPLICATION)))), // 小图标
	}

	// 注册窗口类
	if atom := win.RegisterClassEx(&wc); atom == 0 {
		return 0, fmt.Errorf("%v", win.GetLastError())
	}

	// 创建窗口
	hInstance := win.GetModuleHandle(nil)
	hwnd := win.CreateWindowEx(
		win.WS_EX_TOPMOST|win.WS_EX_NOACTIVATE|win.WS_EX_LAYERED, // 窗口扩展样式
		className,           // 窗口类名
		windowTitle,         // 窗口标题
		win.WS_POPUP,        // 窗口样式
		0,                   // 窗口左上角X坐标
		0,                   // 窗口左上角Y坐标
		int32(screenWidth),  // 窗口宽度
		int32(screenHeight), // 窗口高度
		0,                   // 父窗口句柄
		0,                   // 菜单句柄
		hInstance,           // 实例句柄
		nil,                 // 附加参数
	)
	if hwnd == 0 {
		return 0, fmt.Errorf("%v", win.GetLastError())
	}

	// 设置分层窗口属性
	result, _, _ := setLayeredWindowAttributes.Call(uintptr(hwnd), 0x00000000, uintptr(128), 0x00000001)
	if result == 0 {
		fmt.Printf("Error setting layered window attributes %v\n", fmt.Errorf("%v", win.GetLastError()))
	}

	// 获取当前扩展窗口样式
	style := win.GetWindowLongPtr(hwnd, win.GWL_EXSTYLE)

	// // 添加WS_EX_TRANSPARENT样式
	style |= win.WS_EX_TRANSPARENT

	// 设置新的扩展窗口样式
	win.SetWindowLongPtr(hwnd, win.GWL_EXSTYLE, style)

	// 隐藏光标
	showCursor.Call(0)

	// 显示窗口
	win.ShowWindow(hwnd, win.SW_SHOWDEFAULT)
	return hwnd, nil
}

func (w *Window) Destroy() {
	win.DestroyWindow(w.hwnd)
	w.DeleteDC(w.hdc)
}

func (w *Window) Msg() *win.MSG {
	return &w.msg
}

func (w *Window) RunLoop(loop func()) {
	for win.GetMessage(&w.msg, 0, 0, 0) > 0 {
		win.TranslateMessage(&w.msg)
		win.DispatchMessage(&w.msg)

		win.SetTimer(w.hwnd, 1, 15, 0)

		bgBrush, _, _ := createSolidBrush.Call(uintptr(0x000000))

		memhdc, _, _ := createCompatibleDC.Call(uintptr(w.hdc))
		memBitmap := win.CreateCompatibleBitmap(w.hdc, int32(w.sw), int32(w.sh))
		win.SelectObject(win.HDC(memhdc), win.HGDIOBJ(memBitmap))
		win.SelectObject(win.HDC(memhdc), win.HGDIOBJ(bgBrush))
		win.SetBkMode(win.HDC(memhdc), win.TRANSPARENT)

		if loop != nil {
			loop()
		}

		win.BitBlt(w.hdc, 0, 0, int32(w.sw), int32(w.sh), win.HDC(memhdc), 0, 0, win.SRCCOPY)
		win.DeleteObject(win.HGDIOBJ(memBitmap))
		win.DeleteDC(win.HDC(memhdc))
	}
}
