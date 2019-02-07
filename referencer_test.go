package ptr

import (
	"testing"
	"time"
)

func TestIntPtr(t *testing.T) {
	var testData = int(12)
	resultData := IntPtr(testData)
	if *resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, *resultData)
	}
}

func TestInt8Ptr(t *testing.T) {
	var testData = int8(12)
	resultData := Int8Ptr(testData)
	if *resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, *resultData)
	}
}

func TestInt16Ptr(t *testing.T) {
	var testData = int16(12)
	resultData := Int16Ptr(testData)
	if *resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, *resultData)
	}
}

func TestInt32Ptr(t *testing.T) {
	var testData = int32(12)
	resultData := Int32Ptr(testData)
	if *resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, *resultData)
	}
}

func TestInt64Ptr(t *testing.T) {
	var testData = int64(12)
	resultData := Int64Ptr(testData)
	if *resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, *resultData)
	}
}

func TestUintPtr(t *testing.T) {
	var testData = uint(12)
	resultData := UintPtr(testData)
	if *resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, *resultData)
	}
}

func TestUint8Ptr(t *testing.T) {
	var testData = uint8(12)
	resultData := Uint8Ptr(testData)
	if *resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, *resultData)
	}
}

func TestUint16Ptr(t *testing.T) {
	var testData = uint16(12)
	resultData := Uint16Ptr(testData)
	if *resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, *resultData)
	}
}

func TestUint32Ptr(t *testing.T) {
	var testData = uint32(12)
	resultData := Uint32Ptr(testData)
	if *resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, *resultData)
	}
}

func TestUint64Ptr(t *testing.T) {
	var testData = uint64(12)
	resultData := Uint64Ptr(testData)
	if *resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, *resultData)
	}
}

func TestFloat32Ptr(t *testing.T) {
	var testData = float32(12)
	resultData := Float32Ptr(testData)
	if *resultData != testData {
		t.Errorf("Test data should be %f, instead of %f", testData, *resultData)
	}
}

func TestFloat64Ptr(t *testing.T) {
	var testData = float64(12)
	resultData := Float64Ptr(testData)
	if *resultData != testData {
		t.Errorf("Test data should be %f, instead of %f", testData, *resultData)
	}
}

func TestBoolPtr(t *testing.T) {
	var testData = bool(true)
	resultData := BoolPtr(testData)
	if *resultData != testData {
		t.Errorf("Test data should be %t, instead of %t", testData, *resultData)
	}
}

func TestStringPtr(t *testing.T) {
	var testData = string("test")
	resultData := StringPtr(testData)
	if *resultData != testData {
		t.Errorf("Test data should be %s, instead of %s", testData, *resultData)
	}
}

func TestComplex64Ptr(t *testing.T) {
	var testData = complex64(12)
	resultData := Complex64Ptr(testData)
	if *resultData != testData {
		t.Errorf("Test data should be %f, instead of %f", testData, *resultData)
	}
}

func TestComplex128Ptr(t *testing.T) {
	var testData = complex128(12)
	resultData := Complex128Ptr(testData)
	if *resultData != testData {
		t.Errorf("Test data should be %f, instead of %f", testData, *resultData)
	}
}

func TestTimePtr(t *testing.T) {
	var current = time.Now()
	resultCurrent := TimePtr(current)
	if *resultCurrent != current {
		t.Errorf("Test data should be %v, instead of %v", current, *resultCurrent)
	}
}
