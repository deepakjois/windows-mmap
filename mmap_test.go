package winmmap

import "testing"

func TestMmap(t *testing.T) {
	t.Log("Trying mmap")
	err := trymmap()
	if err != nil {
		t.Errorf("mmap failed with error: %v", err)
	}
}
