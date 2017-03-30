package rtda

import "github.com/ruandao/jvmgo/ch06/rtda/heap"

type Frame struct {
	lower 		*Frame
	LocalVars	LocalVars
	OperandStack	*OperandStack
	thread 		*Thread
	method		*heap.Method
	nextPC		int
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:thread,
		method:method,
		LocalVars:newLocalVars(method.MaxLocals()),
		OperandStack:newOperandStack(method.MaxStack()),
	}
}

func (self *Frame) Method() *heap.Method {
	return self.method
}