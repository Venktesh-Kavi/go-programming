package dll_test

import (
	"dsa/dll"
	"testing"
)

func TestDllLazyInit(t *testing.T) {
	dl := dll.CList{}
	dl.PushFront(10)
	if dl.Len != 1 {
		t.Errorf("dl.Len expected 1, got %d", dl.Len)
	}
}

func TestDllAutoCreation(t *testing.T) {
	dl := dll.New()
	dl.PushFront(10)

	if dl.Len != 1 {
		t.Errorf("dl.Len expected 1, got %d", dl.Len)
	}
}
