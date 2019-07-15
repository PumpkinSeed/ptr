package ptr

import (
	"testing"
)

func TestCloneInt(t *testing.T) {
	var i = NewInt()
	cI := CloneInt(i)
	if i == cI {
		t.Errorf("Address should NOT be equal ( %p != %p )", i, cI)
	}
}

func TestCloneInt8(t *testing.T) {
	var i = NewInt8()
	cI := CloneInt8(i)
	if i == cI {
		t.Errorf("Address should NOT be equal ( %p != %p )", i, cI)
	}
}

func TestCloneInt16(t *testing.T) {
	var i = NewInt16()
	cI := CloneInt16(i)
	if i == cI {
		t.Errorf("Address should NOT be equal ( %p != %p )", i, cI)
	}
}

func TestCloneInt32(t *testing.T) {
	var i = NewInt32()
	cI := CloneInt32(i)
	if i == cI {
		t.Errorf("Address should NOT be equal ( %p != %p )", i, cI)
	}
}

func TestCloneInt64(t *testing.T) {
	var i = NewInt64()
	cI := CloneInt64(i)
	if i == cI {
		t.Errorf("Address should NOT be equal ( %p != %p )", i, cI)
	}
}

func TestCloneUint(t *testing.T) {
	var u = NewUint()
	cU := CloneUint(u)
	if u == cU {
		t.Errorf("Address should NOT be equal ( %p != %p )", u, cU)
	}
}

func TestCloneUint8(t *testing.T) {
	var u = NewUint8()
	cU := CloneUint8(u)
	if u == cU {
		t.Errorf("Address should NOT be equal ( %p != %p )", u, cU)
	}
}

func TestCloneUint16(t *testing.T) {
	var u = NewUint16()
	cU := CloneUint16(u)
	if u == cU {
		t.Errorf("Address should NOT be equal ( %p != %p )", u, cU)
	}
}

func TestCloneUint32(t *testing.T) {
	var u = NewUint32()
	cU := CloneUint32(u)
	if u == cU {
		t.Errorf("Address should NOT be equal ( %p != %p )", u, cU)
	}
}

func TestCloneUint64(t *testing.T) {
	var u = NewUint64()
	cU := CloneUint64(u)
	if u == cU {
		t.Errorf("Address should NOT be equal ( %p != %p )", u, cU)
	}
}

func TestCloneFloat32(t *testing.T) {
	var f = NewFloat32()
	cF := CloneFloat32(f)
	if f == cF {
		t.Errorf("Address should NOT be equal ( %p != %p )", f, cF)
	}
}

func TestCloneFloat64(t *testing.T) {
	var f = NewFloat64()
	cF := CloneFloat64(f)
	if f == cF {
		t.Errorf("Address should NOT be equal ( %p != %p )", f, cF)
	}
}

func TestCloneBool(t *testing.T) {
	var b = NewBool()
	cB := CloneBool(b)
	if b == cB {
		t.Errorf("Address should NOT be equal ( %p != %p )", b, cB)
	}
}

func TestCloneString(t *testing.T) {
	var s = NewString()
	cS := CloneString(s)
	if s == cS {
		t.Errorf("Address should NOT be equal ( %p != %p )", s, cS)
	}
}

func TestCloneComplex64(t *testing.T) {
	var c = NewComplex64()
	cC := CloneComplex64(c)
	if c == cC {
		t.Errorf("Address should NOT be equal ( %p != %p )", c, cC)
	}
}

func TestCloneComplex128(t *testing.T) {
	var c = NewComplex128()
	cC := CloneComplex128(c)
	if c == cC {
		t.Errorf("Address should NOT be equal ( %p != %p )", c, cC)
	}
}

func TestCloneTime(t *testing.T) {
	var t1 = NewTime()
	cT1 := CloneTime(t1)
	if t1 == cT1 {
		t.Errorf("Address should NOT be equal ( %p != %p )", t1, cT1)
	}
}

func BenchmarkCloneUint(b *testing.B) {
	var u = NewUint()
	var cU *uint
	for i := 0; i<b.N;i++ {
		cU = CloneUint(u)
	}
	if u == cU {
		b.Errorf("Address should NOT be equal ( %p != %p )", u, cU)
	}
}