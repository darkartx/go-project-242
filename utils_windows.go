//go:build windows

package path_size

import (
	"syscall"
)

func isHidden(filename string) (bool, error) {
	ptr, err := syscall.UTF16PtrFromString(filename)
	if err != nil {
		return false, err
	}

	attributes, err := syscall.GetFileAttributes(ptr)
	if err != nil {
		return false, err
	}

	result := attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0

	return result, nil
}
