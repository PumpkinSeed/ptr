package ptr

import "time"

func IntPtr(i int) *int {
	return &i
}

func Int8Ptr(i int8) *int8 {
	return &i
}

func Int16Ptr(i int16) *int16 {
	return &i
}

func Int32Ptr(i int32) *int32 {
	return &i
}

func Int64Ptr(i int64) *int64 {
	return &i
}

func UintPtr(u uint) *uint {
	return &u
}

func Uint8Ptr(u uint8) *uint8 {
	return &u
}

func Uint16Ptr(u uint16) *uint16 {
	return &u
}

func Uint32Ptr(u uint32) *uint32 {
	return &u
}

func Uint64Ptr(u uint64) *uint64 {
	return &u
}

func Float32Ptr(f float32) *float32 {
	return &f
}

func Float64Ptr(f float64) *float64 {
	return &f
}

func BoolPtr(b bool) *bool {
	return &b
}

func StringPtr(s string) *string {
	return &s
}

func Complex64Ptr(c complex64) *complex64 {
	return &c
}

func Complex128Ptr(c complex128) *complex128 {
	return &c
}

func TimePtr(t time.Time) *time.Time {
	return &t
}
