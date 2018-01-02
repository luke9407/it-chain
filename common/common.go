package common

import "crypto/sha1"


// Hashable을 구현한 모든 struct는 Hashing할 수 있다.
type Hashable interface {
	toByte() []byte
}


func Hashing(serializable Hashable) []byte{

	var s = serializable.toByte()
	h := sha1.New()
	h.Write(s)
	bs := h.Sum(nil)

	return bs
}