package references

import (
	"github.com/ruandao/jvmgo/ch08/instructions/base"
	"github.com/ruandao/jvmgo/ch08/rtda"
	"github.com/ruandao/jvmgo/ch08/rtda/heap"
)

type ANEW_ARRAY struct {
	base.Index16Instruction
}

func (self *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	componentClass := classRef.ResolvedClass()

	stack := frame.OperandStack
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySzeException")
	}

	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}