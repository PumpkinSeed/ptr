package ptr

func NewString() *string {
	var s = ""
	return &s
}

func NewUint64() *uint64 {
	var u = uint64(0)
	return &u
}
