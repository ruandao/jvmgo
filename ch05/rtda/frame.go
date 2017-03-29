package rtda

type Frame struct {
	lower 		*Frame
	LocalVars	LocalVars
	OperandStack	*OperandStack
	thread 		*Thread
	nextPC		int
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:thread,
		LocalVars:newLocalVars(maxLocals),
		OperandStack:newOperandStack(maxStack),
	}
}