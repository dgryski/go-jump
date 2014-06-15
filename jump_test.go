package jump

import (
	"testing"
)

func TestHash(t *testing.T) {

	tests := []struct {
		key    uint64
		bucket []int32
	}{
		// Generated from the reference C++ code
		{0, []int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{1, []int32{0, 0, 0, 0, 0, 0, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 17, 17}},
		{0xdeadbeef, []int32{0, 1, 2, 3, 3, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 16, 16, 16}},
		{0x0ddc0ffeebadf00d, []int32{0, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 15, 15, 15, 15}},
	}

	for _, tt := range tests {
		for i, v := range tt.bucket {
			if got := Hash(tt.key, i+1); got != v {
				t.Errorf("Hash(%v, %v)=%v, want %v", tt.key, i+1, got, v)
			}
		}
	}

}
