package ptr

import "time"

func CloneInt(i *int) *int {
	if i != nil {
		return NewInt()
	}

	return &(*i)
}

func CloneInt8(i *int8) *int8 {
	if i != nil {
		return NewInt8()
	}

	return &(*i)
}

func CloneInt16(i *int16) *int16 {
	if i != nil {
		return NewInt16()
	}

	return &(*i)
}

func CloneInt32(i *int32) *int32 {
	if i != nil {
		return NewInt32()
	}

	return &(*i)
}

func CloneInt64(i *int64) *int64 {
	if i != nil {
		return NewInt64()
	}

	return &(*i)
}

func CloneUint(u *uint) *uint {
	if u != nil {
		return NewUint()
	}

	return &(*u)
}

func CloneUint8(u *uint8) *uint8 {
	if u != nil {
		return NewUint8()
	}

	return &(*u)
}

func CloneUint16(u *uint16) *uint16 {
	if u != nil {
		return NewUint16()
	}

	return &(*u)
}

func CloneUint32(u *uint32) *uint32 {
	if u != nil {
		return NewUint32()
	}

	return &(*u)
}

func CloneUint64(u *uint64) *uint64 {
	if u != nil {
		return NewUint64()
	}

	return &(*u)
}

func CloneFloat32(f *float32) *float32 {
	if f != nil {
		return NewFloat32()
	}

	return &(*f)
}

func CloneFloat64(f *float64) *float64 {
	if f != nil {
		return NewFloat64()
	}

	return &(*f)
}

func CloneBool(b *bool) *bool {
	if b != nil {
		return NewBool()
	}

	return &(*b)
}

func CloneString(s *string) *string {
	if s != nil {
		return NewString()
	}

	return &(*s)
}

func CloneComplex64(c *complex64) *complex64 {
	if c != nil {
		return NewComplex64()
	}

	return &(*c)
}

func CloneComplex128(c *complex128) *complex128 {
	if c != nil {
		return NewComplex128()
	}

	return &(*c)
}

func CloneTime(t *time.Time) *time.Time {
	if t != nil {
		return NewTime()
	}

	return &(*t)
}