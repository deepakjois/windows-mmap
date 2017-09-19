package winmmap

import "testing"
import "os"
import "github.com/dgraph-io/badger/y"

func TestMmap(t *testing.T) {
	t.Log("Trying mmap")
	f, err := os.OpenFile("README.md", os.O_WRONLY|os.O_TRUNC, 0777)
	y.Check(err)
	fi, err := f.Stat()
	y.Check(err)
	// size := int64(math.MaxUint32) FAILS!
	// size := int64(500 * 1024 * 1024) FAILS!
	// size := int64(10 * 1024 * 1024) FAILS!
	size := fi.Size()
	t.Logf("Size is : %v", size)
	_, err = trymmap(f, size+1)
	if err != nil {
		t.Errorf("mmap failed with error: %v", err)
	}
}
