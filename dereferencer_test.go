package ptr

import "testing"

func TestInt(t *testing.T) {
	var testData = int(12)
	resultData := Int(&testData)
	if resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, resultData)
	}
}

func TestInt8(t *testing.T) {
	var testData = int8(12)
	resultData := Int8(&testData)
	if resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, resultData)
	}
}

func TestInt16(t *testing.T) {
	var testData = int16(12)
	resultData := Int16(&testData)
	if resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, resultData)
	}
}

func TestInt32(t *testing.T) {
	var testData = int32(12)
	resultData := Int32(&testData)
	if resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, resultData)
	}
}

func TestInt64(t *testing.T) {
	var testData = int64(12)
	resultData := Int64(&testData)
	if resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, resultData)
	}
}

func TestUint(t *testing.T) {
	var testData = uint(12)
	resultData := Uint(&testData)
	if resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, resultData)
	}
}

func TestUint8(t *testing.T) {
	var testData = uint8(12)
	resultData := Uint8(&testData)
	if resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, resultData)
	}
}

func TestUint16(t *testing.T) {
	var testData = uint16(12)
	resultData := Uint16(&testData)
	if resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, resultData)
	}
}

func TestUint32(t *testing.T) {
	var testData = uint32(12)
	resultData := Uint32(&testData)
	if resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, resultData)
	}
}

func TestUint64(t *testing.T) {
	var testData = uint64(12)
	resultData := Uint64(&testData)
	if resultData != testData {
		t.Errorf("Test data should be %d, instead of %d", testData, resultData)
	}
}

func TestFloat32(t *testing.T) {
	var testData = float32(12)
	resultData := Float32(&testData)
	if resultData != testData {
		t.Errorf("Test data should be %f, instead of %f", testData, resultData)
	}
}

func TestFloat64(t *testing.T) {
	var testData = float64(12)
	resultData := Float64(&testData)
	if resultData != testData {
		t.Errorf("Test data should be %f, instead of %f", testData, resultData)
	}
}

func TestBool(t *testing.T) {
	var testData = bool(true)
	resultData := Bool(&testData)
	if resultData != testData {
		t.Errorf("Test data should be %t, instead of %t", testData, resultData)
	}
}

func TestString(t *testing.T) {
	var testData = string("test")
	resultData := String(&testData)
	if resultData != testData {
		t.Errorf("Test data should be %s, instead of %s", testData, resultData)
	}
}

func TestComplex64(t *testing.T) {
	var testData = complex64(12)
	resultData := Complex64(&testData)
	if resultData != testData {
		t.Errorf("Test data should be %f, instead of %f", testData, resultData)
	}
}

func TestComplex128(t *testing.T) {
	var testData = complex128(12)
	resultData := Complex128(&testData)
	if resultData != testData {
		t.Errorf("Test data should be %f, instead of %f", testData, resultData)
	}
}
