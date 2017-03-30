package constants

import (
	"github.com/ruandao/jvmgo/ch05/instructions/base"
	"github.com/ruandao/jvmgo/ch04/rtda"
)

// push byte
type BIPUSH struct {
	val int8
}

// push short
type SIPUSH struct {
	val int16
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack.PushInt(i)
}
