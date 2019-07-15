package ptr

import "time"

// CloneInt makes a copy of the value to a new memory address
func CloneInt(i *int) *int {
	if i == nil {
		return NewInt()
	}

	return &(*i)
}

// CloneInt8 makes a copy of the value to a new memory address
func CloneInt8(i *int8) *int8 {
	if i == nil {
		return NewInt8()
	}

	return &(*i)
}

// CloneInt16 makes a copy of the value to a new memory address
func CloneInt16(i *int16) *int16 {
	if i == nil {
		return NewInt16()
	}

	return &(*i)
}

// CloneInt32 makes a copy of the value to a new memory address
func CloneInt32(i *int32) *int32 {
	if i == nil {
		return NewInt32()
	}

	return &(*i)
}

// CloneInt64 makes a copy of the value to a new memory address
func CloneInt64(i *int64) *int64 {
	if i == nil {
		return NewInt64()
	}

	return &(*i)
}

// CloneUint makes a copy of the value to a new memory address
func CloneUint(u *uint) *uint {
	if u == nil {
		return NewUint()
	}

	return &(*u)
}

// CloneUint8 makes a copy of the value to a new memory address
func CloneUint8(u *uint8) *uint8 {
	if u == nil {
		return NewUint8()
	}

	return &(*u)
}

// CloneUint16 makes a copy of the value to a new memory address
func CloneUint16(u *uint16) *uint16 {
	if u == nil {
		return NewUint16()
	}

	return &(*u)
}

// CloneUint32 makes a copy of the value to a new memory address
func CloneUint32(u *uint32) *uint32 {
	if u == nil {
		return NewUint32()
	}

	return &(*u)
}

// CloneUint64 makes a copy of the value to a new memory address
func CloneUint64(u *uint64) *uint64 {
	if u == nil {
		return NewUint64()
	}

	return &(*u)
}

// CloneFloat32 makes a copy of the value to a new memory address
func CloneFloat32(f *float32) *float32 {
	if f == nil {
		return NewFloat32()
	}

	return &(*f)
}

// CloneFloat64 makes a copy of the value to a new memory address
func CloneFloat64(f *float64) *float64 {
	if f == nil {
		return NewFloat64()
	}

	return &(*f)
}

// CloneBool makes a copy of the value to a new memory address
func CloneBool(b *bool) *bool {
	if b == nil {
		return NewBool()
	}

	return &(*b)
}

// CloneString makes a copy of the value to a new memory address
func CloneString(s *string) *string {
	if s == nil {
		return NewString()
	}

	return &(*s)
}

// CloneComplex64 makes a copy of the value to a new memory address
func CloneComplex64(c *complex64) *complex64 {
	if c == nil {
		return NewComplex64()
	}

	return &(*c)
}

// CloneComplex128 makes a copy of the value to a new memory address
func CloneComplex128(c *complex128) *complex128 {
	if c == nil {
		return NewComplex128()
	}

	return &(*c)
}

// CloneTime makes a copy of the value to a new memory address
func CloneTime(t *time.Time) *time.Time {
	if t == nil {
		return NewTime()
	}

	return &(*t)
}