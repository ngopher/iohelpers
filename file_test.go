package iohelpers

import (
	"testing"
)

func TestNewFile(t *testing.T) {
	f := NewFile(nil)
	if f == nil {
		t.Error("NewFile() should not return nil")
		t.Fail()
	}
	if f.File != nil {
		t.Error("NewFile() should return a file with nil file")
		t.Fail()
	}
}
