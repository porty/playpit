// +build windows

package netsend

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	kernel32, _        = syscall.LoadLibrary("kernel32.dll")
	getModuleHandle, _ = syscall.GetProcAddress(kernel32, "GetModuleHandleW")

	user32, _     = syscall.LoadLibrary("user32.dll")
	messageBox, _ = syscall.GetProcAddress(user32, "MessageBoxW")
)

func MessageBox(caption, text string, style uintptr) (result int, err error) {
	var nargs uintptr = 4
	ret, _, callErr := syscall.Syscall9(uintptr(messageBox),
		nargs,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		style,
		0,
		0,
		0,
		0,
		0)
	if callErr != 0 {
		//abort("Call MessageBox", callErr)
		err := fmt.Errorf("bad return value from MessageBoxW: %d", callErr)
		return 0, err
	}
	result = int(ret)
	return result, nil
}

func GetModuleHandle() uintptr {
	var nargs uintptr
	handle, _, callErr := syscall.Syscall(uintptr(getModuleHandle), nargs, 0, 0, 0)
	if callErr != 0 {
		panic(fmt.Errorf("bad result from GetModuleHandle: %d", callErr))
	}

	return handle
}
