//go:build windows
// +build windows

package griagent

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// processort custom type
type procArchType uint

const (
	PROC_ARCH_64 = iota
	PROC_ARCH_32

	PROC_ARCH_UNSUPPORTED
)

// convert procArchType to human readable string
func (pa procArchType) String() string {
	return [...]string{
		"64Bit",
		"32Bit",
		"Unsupported",
	}[pa]
}

// get processor architecture type
// currently only 64bit and 32bit are supported
func getProcessorInfo() (procArchType, error) {
	pa := os.Getenv("PROCESSOR_ARCHITECTURE")
	if pa != "" {
		switch strings.ToLower(pa) {
		case "amd64":
			return PROC_ARCH_64, nil
		case "ia64":
			return PROC_ARCH_64, nil
		case "x86":
			return PROC_ARCH_32, nil
		}

	}

	return PROC_ARCH_UNSUPPORTED, errors.New(fmt.Sprintf(
		"unkown architecture: %v", pa,
	))

}
