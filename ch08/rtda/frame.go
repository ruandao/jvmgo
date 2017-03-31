package rtda

import "github.com/ruandao/jvmgo/ch07/rtda/heap"

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

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}