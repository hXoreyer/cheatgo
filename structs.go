package cheatgo

// 类型定义
type (
	ATOM            uint16  // 定义ATOM为uint16类型
	BOOL            int32   // 定义BOOL为int32类型
	COLORREF        uint32  // 定义COLORREF为uint32类型
	DWM_FRAME_COUNT uint64  // 定义DWM_FRAME_COUNT为uint64类型
	DWORD           uint32  // 定义DWORD为uint32类型
	HACCEL          HANDLE  // 定义HACCEL为HANDLE类型
	HANDLE          uintptr // 定义HANDLE为uintptr类型
	HBITMAP         HANDLE  // 定义HBITMAP为HANDLE类型
	HBRUSH          HANDLE  // 定义HBRUSH为HANDLE类型
	HCURSOR         HANDLE  // 定义HCURSOR为HANDLE类型
	HDC             HANDLE  // 定义HDC为HANDLE类型
	HDROP           HANDLE  // 定义HDROP为HANDLE类型
	HDWP            HANDLE  // 定义HDWP为HANDLE类型
	HENHMETAFILE    HANDLE  // 定义HENHMETAFILE为HANDLE类型
	HFONT           HANDLE  // 定义HFONT为HANDLE类型
	HGDIOBJ         HANDLE  // 定义HGDIOBJ为HANDLE类型
	HGLOBAL         HANDLE  // 定义HGLOBAL为HANDLE类型
	HGLRC           HANDLE  // 定义HGLRC为HANDLE类型
	HHOOK           HANDLE  // 定义HHOOK为HANDLE类型
	HICON           HANDLE  // 定义HICON为HANDLE类型
	HIMAGELIST      HANDLE  // 定义HIMAGELIST为HANDLE类型
	HINSTANCE       HANDLE  // 定义HINSTANCE为HANDLE类型
	HKEY            HANDLE  // 定义HKEY为HANDLE类型
	HKL             HANDLE  // 定义HKL为HANDLE类型
	HMENU           HANDLE  // 定义HMENU为HANDLE类型
	HMODULE         HANDLE  // 定义HMODULE为HANDLE类型
	HMONITOR        HANDLE  // 定义HMONITOR为HANDLE类型
	HPEN            HANDLE  // 定义HPEN为HANDLE类型
	HRESULT         int32   // 定义HRESULT为int32类型
	HRGN            HANDLE  // 定义HRGN为HANDLE类型
	HRSRC           HANDLE  // 定义HRSRC为HANDLE类型
	HTHUMBNAIL      HANDLE  // 定义HTHUMBNAIL为HANDLE类型
	HWND            HANDLE  // 定义HWND为HANDLE类型
	LPARAM          uintptr // 定义LPARAM为uintptr类型
	LRESULT         uintptr // 定义LRESULT为uintptr类型
	QPC_TIME        uint64  // 定义QPC_TIME为uint64类型
	SIZE_T          uintptr // 定义SIZE_T为uintptr类型
	TRACEHANDLE     uintptr // 定义TRACEHANDLE为uintptr类型
	ULONG_PTR       uintptr // 定义ULONG_PTR为uintptr类型
	WPARAM          uintptr // 定义WPARAM为uintptr类型
)

// 鼠标输入结构体
type MOUSEINPUT struct {
	Dx          int32   // 鼠标的水平移动量
	Dy          int32   // 鼠标的垂直移动量
	MouseData   uint32  // 额外的鼠标数据
	DwFlags     uint32  // 事件类型标志
	Time        uint32  // 事件时间戳
	DwExtraInfo uintptr // 额外信息
}

// 键盘输入结构体
type KEYBDINPUT struct {
	WVk         uint16  // 虚拟键码
	WScan       uint16  // 硬件扫描码
	DwFlags     uint32  // 事件类型标志
	Time        uint32  // 事件时间戳
	DwExtraInfo uintptr // 额外信息
}

// 硬件输入结构体
type HARDWAREINPUT struct {
	UMsg    uint32 // 消息
	WParamL uint16 // 参数1
	WParamH uint16 // 参数2
}

// 输入事件结构体
type INPUT struct {
	Type uint32        // 输入事件类型
	Mi   MOUSEINPUT    // 鼠标输入
	Ki   KEYBDINPUT    // 键盘输入
	Hi   HARDWAREINPUT // 硬件输入
}

// 消息结构体
type MSG struct {
	Hwnd    HWND    // 窗口句柄
	Message uint32  // 消息
	WParam  uintptr // 参数1
	LParam  uintptr // 参数2
	Time    uint32  // 时间戳
	Pt      POINT   // 鼠标位置
}

// 点结构体
type POINT struct {
	X, Y int32 // 坐标
}

// 矩形结构体
type RECT struct {
	Left, Top, Right, Bottom int32 // 矩形的边界
}

// 模块条目结构体
type MODULEENTRY32 struct {
	Size         uint32                        // 结构体大小
	ModuleID     uint32                        // 模块ID
	ProcessID    uint32                        // 进程ID
	GlblcntUsage uint32                        // 全局使用计数
	ProccntUsage uint32                        // 进程使用计数
	ModBaseAddr  *uint8                        // 模块基地址
	ModBaseSize  uint32                        // 模块大小
	HModule      HMODULE                       // 模块句柄
	SzModule     [MAX_MODULE_NAME32 + 1]uint16 // 模块名称
	SzExePath    [MAX_PATH]uint16              // 可执行文件路径
}

// 进程条目结构体
type PROCESSENTRY32 struct {
	Size            uint32          // 结构体大小
	Usage           uint32          // 使用计数
	ProcessID       uint32          // 进程ID
	DeafultHeapID   uintptr         // 默认堆ID
	ModuleID        uint32          // 模块ID
	Threads         uint32          // 线程数
	ParentProcessID uint32          // 父进程ID
	PriClassBase    uint32          // 优先级基数
	Flags           uint32          // 标志
	SzExeFile       [MAX_PATH]uint8 // 可执行文件名称
}

// 钩子过程函数类型定义
type HOOKPROC func(int, WPARAM, LPARAM) LRESULT
