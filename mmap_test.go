package winmmap

import "testing"
import "os"
import "github.com/dgraph-io/badger/y"
import "math"

func TestMmap(t *testing.T) {
	t.Log("Trying mmap")
	f, err := os.Open("README.md")
	y.Check(err)
	_, err = f.Stat()
	y.Check(err)
	size := int64(math.MaxUint32)
	_, err = trymmap(f, size)
	if err != nil {
		t.Errorf("mmap failed with error: %v", err)
	}
}
