package util

import "testing"

func TestRuntimeUtil_GetRuntimeDir(t *testing.T) {
	util := RuntimeUtil{}
	dir, err := util.GetRuntimeDir()
	if err != nil {
		t.Error(err)
	}

	t.Log(dir)
}
