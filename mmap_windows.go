package winmmap

import (
	"os"
	"syscall"
	"unsafe"
)

func trymmap(fd *os.File, size int64) ([]byte, error) {
	protect := syscall.PAGE_READONLY
	access := syscall.FILE_MAP_READ

	handler, err := syscall.CreateFileMapping(syscall.Handle(fd.Fd()), nil,
		uint32(protect), uint32(size>>32), uint32(size), nil)
	if err != nil {
		return nil, err
	}
	defer syscall.CloseHandle(handler)

	mapData, err := syscall.MapViewOfFile(handler, uint32(access), 0, 0, 0)
	if err != nil {
		return nil, err
	}

	data := (*[1 << 30]byte)(unsafe.Pointer(mapData))[:size]
	return data, nil
}
