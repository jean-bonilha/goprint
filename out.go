// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package goprint

import "unsafe"
import "syscall"

var _ unsafe.Pointer

var (
	modwinspool = syscall.NewLazyDLL("winspool.drv")

	procIsValidDevmodeW    = modwinspool.NewProc("IsValidDevmodeW")
	procSetDefaultPrinterW = modwinspool.NewProc("SetDefaultPrinterW")
	procOpenPrinter2W      = modwinspool.NewProc("OpenPrinter2W")
)

func IsvalidDevMode(dev *DevMode, buf uint16) (b bool) {
	r0, _, _ := syscall.Syscall(procIsValidDevmodeW.Addr(), 2, uintptr(unsafe.Pointer(dev)), uintptr(buf), 0)
	b = r0 != 0
	return
}

func SetDefaultPrinter(printerName string) (b bool) {
	var p0 *uint16
	p0, _ = syscall.UTF16PtrFromString(printerName)
	if _ != nil {
		return
	}
	return _SetDefaultPrinter(p0)
}

func _SetDefaultPrinter(printerName *uint16) (b bool) {
	r0, _, _ := syscall.Syscall(procSetDefaultPrinterW.Addr(), 1, uintptr(unsafe.Pointer(printerName)), 0, 0)
	b = r0 != 0
	return
}

func OpenPrinter2(name *uint16, h *syscall.Handle, defaults uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procOpenPrinter2W.Addr(), 3, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(h)), uintptr(defaults))
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
