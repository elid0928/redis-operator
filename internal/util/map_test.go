package util

import "testing"

func TestCopyMap(t *testing.T) {
	src := map[string]int{"a": 1, "b": 2}
	dst := CopyMap(src)
	if len(dst) != len(src) {
		t.Fatalf("expected length %d, got %d", len(src), len(dst))
	}
	for k, v := range src {
		if dv, ok := dst[k]; !ok || dv != v {
			t.Errorf("key %s expected %d got %d", k, v, dv)
		}
	}
	dst["c"] = 3
	if _, ok := src["c"]; ok {
		t.Errorf("modifying copy should not affect source")
	}

	if CopyMap[string, int](nil) != nil {
		t.Errorf("expected nil when copying nil map")
	}
}
