package griagent

import "golang.org/x/sys/windows"

// constant fo DLL and their methods

type (
	LPVOID  uintptr
	DWORD   uint32
	LPBYTE  byte
	PBYTE   byte
	LPDWORD uint32
	LPWSTR  uint16
	LPCWSTR uint16
)

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")
	// modversion  = windows.NewLazySystemDLL("version.dll")
	modpsapi = windows.NewLazySystemDLL("psapi.dll")
	// modntdll    = windows.NewLazySystemDLL("ntdll.dll")

	procGetSystemInfo = modkernel32.NewProc("GetSystemInfo")

	// procGetNativeSystemInfo       = modkernel32.NewProc("GetNativeSystemInfo")
	// procGetTickCount64            = modkernel32.NewProc("GetTickCount64")
	// procGetSystemTimes            = modkernel32.NewProc("GetSystemTimes")
	// procGlobalMemoryStatusEx      = modkernel32.NewProc("GlobalMemoryStatusEx")
	// procReadProcessMemory         = modkernel32.NewProc("ReadProcessMemory")
	// procGetProcessHandleCount     = modkernel32.NewProc("GetProcessHandleCount")
	// procGetFileVersionInfoW       = modversion.NewProc("GetFileVersionInfoW")
	// procGetFileVersionInfoSizeW   = modversion.NewProc("GetFileVersionInfoSizeW")
	// procVerQueryValueW            = modversion.NewProc("VerQueryValueW")
	// procGetProcessMemoryInfo      = modpsapi.NewProc("GetProcessMemoryInfo")
	// procGetProcessImageFileNameA  = modpsapi.NewProc("GetProcessImageFileNameA")
	procEnumProcesses = modpsapi.NewProc("EnumProcesses")
	// procNtQueryInformationProcess = modntdll.NewProc("NtQueryInformationProcess")
)
