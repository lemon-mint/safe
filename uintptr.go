package safe

import (
	"fmt"
	"strconv"
	"strings"
)

var ErrNotPointer = fmt.Errorf("not a pointer")

func ToUintptr(v interface{}) (uintptr, error) {
	if v == nil {
		return 0, nil
	}
	vPtrStr := fmt.Sprintf("%p", v)
	if vPtrStr == "<nil>" {
		return 0, nil
	}
	if strings.HasPrefix(vPtrStr, "0x") {
		vPtrStr = vPtrStr[2:]
	} else {
		return 0, ErrNotPointer
	}
	ptr, err := strconv.ParseUint(vPtrStr, 16, 64)
	if err != nil {
		return 0, err
	}
	return uintptr(ptr), nil
}
