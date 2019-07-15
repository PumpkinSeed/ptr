package ptr

import "time"

// IntPtr returns the pointer value of the value
func IntPtr(i int) *int {
	return &i
}

// Int8Ptr returns the pointer value of the value
func Int8Ptr(i int8) *int8 {
	return &i
}

// Int16Ptr returns the pointer value of the value
func Int16Ptr(i int16) *int16 {
	return &i
}

// Int32Ptr returns the pointer value of the value
func Int32Ptr(i int32) *int32 {
	return &i
}

// Int64Ptr returns the pointer value of the value
func Int64Ptr(i int64) *int64 {
	return &i
}

// UintPtr returns the pointer value of the value
func UintPtr(u uint) *uint {
	return &u
}

// Uint8Ptr returns the pointer value of the value
func Uint8Ptr(u uint8) *uint8 {
	return &u
}

// Uint16Ptr returns the pointer value of the value
func Uint16Ptr(u uint16) *uint16 {
	return &u
}

// Uint32Ptr returns the pointer value of the value
func Uint32Ptr(u uint32) *uint32 {
	return &u
}

// Uint64Ptr returns the pointer value of the value
func Uint64Ptr(u uint64) *uint64 {
	return &u
}

// Float32Ptr returns the pointer value of the value
func Float32Ptr(f float32) *float32 {
	return &f
}

// Float64Ptr returns the pointer value of the value
func Float64Ptr(f float64) *float64 {
	return &f
}

// BoolPtr returns the pointer value of the value
func BoolPtr(b bool) *bool {
	return &b
}

// StringPtr returns the pointer value of the value
func StringPtr(s string) *string {
	return &s
}

// Complex64Ptr returns the pointer value of the value
func Complex64Ptr(c complex64) *complex64 {
	return &c
}

// Complex128Ptr returns the pointer value of the value
func Complex128Ptr(c complex128) *complex128 {
	return &c
}

// TimePtr returns the pointer value of the value
func TimePtr(t time.Time) *time.Time {
	return &t
}
