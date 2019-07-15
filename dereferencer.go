package ptr

import "time"

// Int returns the value behind the pointer value
func Int(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

// Int8 returns the value behind the pointer value
func Int8(i *int8) int8 {
	if i == nil {
		return 0
	}
	return *i
}

// Int16 returns the value behind the pointer value
func Int16(i *int16) int16 {
	if i == nil {
		return 0
	}
	return *i
}

// Int32 returns the value behind the pointer value
func Int32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

// Int64 returns the value behind the pointer value
func Int64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// Uint returns the value behind the pointer value
func Uint(u *uint) uint {
	if u == nil {
		return 0
	}
	return *u
}

// Uint8 returns the value behind the pointer value
func Uint8(u *uint8) uint8 {
	if u == nil {
		return 0
	}
	return *u
}

// Uint16 returns the value behind the pointer value
func Uint16(u *uint16) uint16 {
	if u == nil {
		return 0
	}
	return *u
}

// Uint32 returns the value behind the pointer value
func Uint32(u *uint32) uint32 {
	if u == nil {
		return 0
	}
	return *u
}

// Uint64 returns the value behind the pointer value
func Uint64(u *uint64) uint64 {
	if u == nil {
		return 0
	}
	return *u
}

// Float32 returns the value behind the pointer value
func Float32(f *float32) float32 {
	if f == nil {
		return 0
	}
	return *f
}

// Float64 returns the value behind the pointer value
func Float64(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

// Bool returns the value behind the pointer value
func Bool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// String returns the value behind the pointer value
func String(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// Complex64 returns the value behind the pointer value
func Complex64(c *complex64) complex64 {
	if c == nil {
		return 0
	}
	return *c
}

// Complex128 returns the value behind the pointer value
func Complex128(c *complex128) complex128 {
	if c == nil {
		return 0
	}
	return *c
}

// Time returns the value behind the pointer value
func Time(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}
