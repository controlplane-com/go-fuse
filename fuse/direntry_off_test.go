package fuse

import (
	"testing"
	"unsafe"
)

func TestDirEntryListOff(t *testing.T) {
	l := NewDirEntryList(make([]byte, 4096), 0)
	l.AddDirEntry(DirEntry{Name: "a", Ino: 1, Off: 42})
	l.AddDirEntry(DirEntry{Name: "b", Ino: 2})
	first := (*_Dirent)(unsafe.Pointer(&l.buf[0]))
	second := (*_Dirent)(unsafe.Pointer(&l.buf[direntSize+8]))
	if first.Off != 42 || second.Off != 43 {
		t.Fatalf("offsets %d, %d; want 42, 43", first.Off, second.Off)
	}
}
