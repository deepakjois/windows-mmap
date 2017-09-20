package winmmap

import "testing"
import "os"
import "github.com/dgraph-io/badger/y"

func TestMmap(t *testing.T) {
	t.Log("Trying mmap")
	flags := os.O_RDWR | os.O_CREATE | os.O_EXCL
	f, err := os.OpenFile("test.md", flags, 0666)
	defer f.Close()
	y.Check(err)
	// size := int64(math.MaxUint32) FAILS!
	size := int64(500 * 1024 * 1024)
	// size := int64(10 * 1024 * 1024) FAILS!
	// size := fi.Size()
	// size := int64(math.MaxUint32)
	t.Logf("Size is : %v", size)
	_, err = trymmap(f, size)
	if err != nil {
		t.Errorf("mmap failed with error: %v", err)
	}
}

func TestMmapAndTruncate(t *testing.T) {
	f, err := os.OpenFile("README.md", os.O_RDWR, 0777)
	defer f.Close()
	y.Check(err)
	fi, err := f.Stat()
	y.Check(err)
	size := fi.Size()
	_, err = trymmap(f, size)
	y.Check(err)
	err = f.Truncate(size - 10)
	if err != nil {
		t.Errorf("truncate after mmap: %v", err)
	}
}
