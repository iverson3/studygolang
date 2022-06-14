package utils

func Equals(a interface{}, b interface{}) bool {
	sliceA, ok1 := a.([]byte)
	sliceB, ok2 := b.([]byte)
	if ok1 && ok2 {
		return BytesEquals(sliceA, sliceB)
	}
	return a == b
}

func BytesEquals(a []byte, b []byte) bool {
	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}