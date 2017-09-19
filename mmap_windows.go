// +build windows

package winmmap

import (
	"os"
	"syscall"
	"unsafe"
)

func trymmap(fd *os.File, size int64) ([]byte, error) {
	protect := syscall.PAGE_READONLY
	access := syscall.FILE_MAP_READ

	// FIXME Bolt has this in its Windows mmap-ing code.
	// Not sure if we need it.
	//
	// Truncate the database to the size of the mmap.
	//if err := fd.Truncate(size); err != nil {
	//	return fmt.Errorf("truncate: %s", err)
	//}

	// Open a file mapping handle.
	sizelo := uint32(size >> 32)
	sizehi := uint32(size) & 0xffffffff

	// Create the memory map.
	handler, err := syscall.CreateFileMapping(syscall.Handle(fd.Fd()), nil,
		uint32(protect), sizelo, sizehi, nil)
	if err != nil {
		return nil, err
	}

	addr, err := syscall.MapViewOfFile(handler, uint32(access), 0, 0, uintptr(size))
	if addr == 0 {
		return nil, os.NewSyscallError("MapViewOfFile", err)
	}

	// Close mapping handle.
	if err := syscall.CloseHandle(syscall.Handle(h)); err != nil {
		return os.NewSyscallError("CloseHandle", err)
	}

	data := (*[size]byte)(unsafe.Pointer(addr))[:size]
	return data, nil
}
