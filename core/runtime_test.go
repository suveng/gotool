package util

import (
	"testing"
)

func TestRuntimeUtil_GetRuntimeDir(t *testing.T) {
	dir, err := GetRuntimeDir()
	if err != nil {
		t.Error(err)
	}

	t.Log(dir)
}
