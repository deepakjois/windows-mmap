// +build windows

package winmmap

import (
	"fmt"
	"math"
	"os"
	"syscall"
	"unsafe"
)

func trymmap(fd *os.File, size int64) ([]byte, error) {
	protect := syscall.PAGE_READONLY
	access := syscall.FILE_MAP_READ

	fi, err := fd.Stat()
	if err != nil {
		return nil, err
	}

	// Truncate the database to the size of the mmap.
	if fi.Size() < size {
		if err := fd.Truncate(size); err != nil {
			return nil, fmt.Errorf("truncate: %s", err)
		}
	}

	// Open a file mapping handle.
	sizelo := uint32(size >> 32)
	sizehi := uint32(size) & 0xffffffff

	// Create the memory map.
	handler, err := syscall.CreateFileMapping(syscall.Handle(fd.Fd()), nil,
		uint32(protect), sizelo, sizehi, nil)
	if err != nil {
		return nil, os.NewSyscallError("CreateFileMapping", err)
	}

	addr, err := syscall.MapViewOfFile(handler, uint32(access), 0, 0, uintptr(size))
	if addr == 0 {
		return nil, os.NewSyscallError("MapViewOfFile", err)
	}

	// Close mapping handle.
	if err := syscall.CloseHandle(syscall.Handle(handler)); err != nil {
		return nil, os.NewSyscallError("CloseHandle", err)
	}

	data := (*[math.MaxUint32]byte)(unsafe.Pointer(addr))[:size]
	return data, nil
}
