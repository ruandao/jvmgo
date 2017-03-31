package lang

import (
	"github.com/ruandao/jvmgo/ch08/native"
	"github.com/ruandao/jvmgo/ch08/rtda"
	"github.com/ruandao/jvmgo/ch08/rtda/heap"
)

func init() {
	native.Register("java/lang/String", "intern", "()Ljava/lang/String;" intern)
}

func intern(frame *rtda.Frame) {
	this := frame.LocalVars.GetThis()
	interned := heap.InternString(this)
	frame.OperandStack.PushRef(interned)
}
