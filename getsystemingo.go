package griagent

import (
	"errors"
	"os"
	"strings"
)

type ProcArchType uint

const (
	PROC_ARCH_64 = iota
	PROC_ARCH_32

	PROC_ARCH_UNSUPPORTED
)

func (pa ProcArchType) String() string {
	return [...]string{
		"64Bit",
		"32Bit",
	}[pa]
}

func (ws *WinAPI) GetProcessorInfo() (ProcArchType, error) {

	if pa := os.Getenv("PROCESSOR_ARCHITECTURE"); pa != "" {
		switch strings.ToLower(pa) {
		case "amd64":
			return PROC_ARCH_64, nil
		case "ia64":
			return PROC_ARCH_64, nil
		case "x86":
			return PROC_ARCH_32, nil
		}
	}

	return PROC_ARCH_UNSUPPORTED, errors.New("unkown architecture")

}

// // see https://msdn.microsoft.com/en-us/library/windows/desktop/ms724958(v=vs.85).aspx
// type systeminfo struct {
// 	wProcessorArchitecture      uint16
// 	wReserved                   uint16
// 	dwPageSize                  uint32
// 	lpMinimumApplicationAddress uintptr
// 	lpMaximumApplicationAddress uintptr
// 	dwActiveProcessorMask       uintptr
// 	dwNumberOfProcessors        uint32
// 	dwProcessorType             uint32
// 	dwAllocationGranularity     uint32
// 	wProcessorLevel             uint16
// 	wProcessorRevision          uint16
// }

// // See https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/ns-sysinfoapi-system_info
// const (
// 	PROCESSOR_ARCHITECTURE_AMD64 = 9
// 	PROCESSOR_ARCHITECTURE_INTEL = 0
// 	PROCESSOR_ARCHITECTURE_ARM   = 5
// 	PROCESSOR_ARCHITECTURE_ARM64 = 12
// 	PROCESSOR_ARCHITECTURE_IA64  = 6
// )

// var sysinfo systeminfo

// func (ws *WinAPI) SystemInfo() (string, error) {

// 	syscall.Syscall(procGetSystemInfo.Addr(), 1, uintptr(unsafe.Pointer(&sysinfo)), 0, 0)

// 	log.Println(sysinfo.wProcessorArchitecture)

// 	switch sysinfo.wProcessorArchitecture {
// 	case PROCESSOR_ARCHITECTURE_AMD64:
// 		return "amd64", nil
// 	case PROCESSOR_ARCHITECTURE_IA64:
// 		return "amd64", nil
// 	case PROCESSOR_ARCHITECTURE_INTEL:
// 		return "386", nil
// 	case PROCESSOR_ARCHITECTURE_ARM:
// 		return "arm", errors.New("unsupported processor architecture")
// 	case PROCESSOR_ARCHITECTURE_ARM64:
// 		return "arm", errors.New("unsupported processor architecture")
// 	default:
// 		return "", errors.New("unknown processor architecture")
// 	}
// }
