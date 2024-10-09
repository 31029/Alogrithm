package others

import "testing"

func TestMain(t *testing.T)  {
	C := Constructor(3)
	C.Put(1, 1)
	C.Put(2, 2)
	C.Put(3, 3)
	C.Put(4, 4)
	C.Get(4)
	C.Get(3)
	C.Get(2)
	C.Get(1)
	C.Put(5, 5)
	C.Get(1)
	C.Get(2)
	C.Get(3)
	C.Get(4)
	C.Get(5)
}