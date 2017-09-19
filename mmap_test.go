package winmmap

import "testing"
import "os"
import "github.com/dgraph-io/badger/y"

func TestMmap(t *testing.T) {
	t.Log("Trying mmap")
	f, err := os.Open("README.md")
	y.Check(err)
	fi, err := f.Stat()
	y.Check(err)
	_, err = trymmap(f, fi.Size())
	if err != nil {
		t.Errorf("mmap failed with error: %v", err)
	}
}
