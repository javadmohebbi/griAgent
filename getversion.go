package griagent

import (
	"log"
	"syscall"

	"golang.org/x/sys/windows"
)

func (wp *WinAPI) GetWinVersion() (byte, uint8, uint16) {
	h, err := windows.LoadLibrary("kernel32.dll")
	if err != nil {
		log.Fatalln("LoadLibrary", err)
	}
	defer windows.FreeLibrary(h)
	proc, err := windows.GetProcAddress(h, "GetVersion")
	if err != nil {
		log.Fatalln("GetProcAddress", err)
	}
	r, _, _ := syscall.Syscall(uintptr(proc), 0, 0, 0, 0)
	major := byte(r)
	minor := uint8(r >> 8)
	build := uint16(r >> 16)

	return major, minor, build

}
