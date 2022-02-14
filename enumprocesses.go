package griagent

import (
	"syscall"
	"unsafe"
)

func (wp *WinAPI) EnumProcesses() (pids []uint32, err error) {

	// var pids []uint32
	var __pids []uint32

	for nAlloc, nGot := uint32(500), uint32(0); ; nAlloc *= 2 {
		_pids := make([]uint32, nAlloc)
		if err = wp.enumProcesses(&_pids[0], nAlloc*4, &nGot); err != nil {
			return nil, err
		}
		if nGot/4 < nAlloc {
			__pids = _pids
			break
		}
	}

	for _, pid := range __pids {
		if pid != 0 {
			pids = append(pids, pid)
		}
	}

	return pids, nil

}

func (wp *WinAPI) enumProcesses(lpidProcess *uint32, cb uint32, lpcbNeeded *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procEnumProcesses.Addr(), 3, uintptr(unsafe.Pointer(lpidProcess)), uintptr(cb), uintptr(unsafe.Pointer(lpcbNeeded)))
	if r1 == 0 {
		if e1 != 0 {
			// err = errnoErr(e1)
			e := syscall.Errno(e1)
			switch e {
			case 0:
				return nil
			case 997:
				err = syscall.ERROR_IO_PENDING
			}
		} else {
			err = syscall.EINVAL
		}
	}
	return

}
