package rtda

import "github.com/ruandao/jvmgo/ch06/rtda/heap"

type Slot struct {
	num 	int32
	ref 	*heap.Object
}
