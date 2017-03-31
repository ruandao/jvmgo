package extended

import (
	"github.com/ruandao/jvmgo/ch05/instructions/base"
	"github.com/ruandao/jvmgo/ch05/rtda"
)

type IFNULL struct {
	base.BranchInstruction
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack.PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}
