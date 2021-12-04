package safe

import (
	"testing"
	"unsafe"
)

func TestToUintptr(t *testing.T) {
	type c struct {
		name    string
		v       interface{}
		want    uintptr
		wantErr bool
	}

	tests := []c{}

	var nilI *int
	tests = append(tests, c{
		name:    "nilI",
		v:       nilI,
		want:    uintptr(unsafe.Pointer(nilI)),
		wantErr: false,
	})
	tests = append(tests, c{
		name:    "nil",
		v:       nil,
		want:    uintptr(unsafe.Pointer(nil)),
		wantErr: false,
	})
	v0 := 0
	tests = append(tests, c{
		name:    "int",
		v:       &v0,
		want:    uintptr(unsafe.Pointer(&v0)),
		wantErr: false,
	})
	v1 := make([]int, 0)
	tests = append(tests, c{
		name:    "slice",
		v:       &v1,
		want:    uintptr(unsafe.Pointer(&v1)),
		wantErr: false,
	})
	v2 := make(map[string]int)
	tests = append(tests, c{
		name:    "map",
		v:       &v2,
		want:    uintptr(unsafe.Pointer(&v2)),
		wantErr: false,
	})
	v3 := make(chan int)
	tests = append(tests, c{
		name:    "chan",
		v:       &v3,
		want:    uintptr(unsafe.Pointer(&v3)),
		wantErr: false,
	})
	v4 := [2]int{}
	tests = append(tests, c{
		name:    "array",
		v:       &v4,
		want:    uintptr(unsafe.Pointer(&v4)),
		wantErr: false,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUintptr(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUintptr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUintptr() = %v, want %v", got, tt.want)
			}
		})
	}
}
