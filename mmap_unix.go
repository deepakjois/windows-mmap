// +build !windows

package winmmap

import "os"

func trymmap(fd *os.File, size int64) ([]byte, error) {
	return nil, nil
}
