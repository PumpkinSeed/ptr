package ptr

import "time"

// NewInt initiate a new pointer value with the default value of the certain type
func NewInt() *int {
	var i = int(0)
	return &i
}

// NewInt8 initiate a new pointer value with the default value of the certain type
func NewInt8() *int8 {
	var i = int8(0)
	return &i
}

// NewInt16 initiate a new pointer value with the default value of the certain type
func NewInt16() *int16 {
	var i = int16(0)
	return &i
}

// NewInt32 initiate a new pointer value with the default value of the certain type
func NewInt32() *int32 {
	var i = int32(0)
	return &i
}

// NewInt64 initiate a new pointer value with the default value of the certain type
func NewInt64() *int64 {
	var i = int64(0)
	return &i
}

// NewUint initiate a new pointer value with the default value of the certain type
func NewUint() *uint {
	var i = uint(0)
	return &i
}

// NewUint8 initiate a new pointer value with the default value of the certain type
func NewUint8() *uint8 {
	var i = uint8(0)
	return &i
}

// NewUint16 initiate a new pointer value with the default value of the certain type
func NewUint16() *uint16 {
	var i = uint16(0)
	return &i
}

// NewUint32 initiate a new pointer value with the default value of the certain type
func NewUint32() *uint32 {
	var i = uint32(0)
	return &i
}

// NewUint64 initiate a new pointer value with the default value of the certain type
func NewUint64() *uint64 {
	var i = uint64(0)
	return &i
}

// NewFloat32 initiate a new pointer value with the default value of the certain type
func NewFloat32() *float32 {
	var i = float32(0)
	return &i
}

// NewFloat64 initiate a new pointer value with the default value of the certain type
func NewFloat64() *float64 {
	var i = float64(0)
	return &i
}

// NewBool initiate a new pointer value with the default value of the certain type
func NewBool() *bool {
	var b = false
	return &b
}

// NewString initiate a new pointer value with the default value of the certain type
func NewString() *string {
	var s = ""
	return &s
}

// NewComplex64 initiate a new pointer value with the default value of the certain type
func NewComplex64() *complex64 {
	var c = complex64(0)
	return &c
}

// NewComplex128 initiate a new pointer value with the default value of the certain type
func NewComplex128() *complex128 {
	var c = complex128(0)
	return &c
}

// NewTime initiate a new pointer value with the default value of the certain type
func NewTime() *time.Time {
	var t = time.Time{}
	return &t
}
