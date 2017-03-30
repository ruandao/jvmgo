package extended

import (
	"github.com/ruandao/jvmgo/ch05/instructions/base"
	"github.com/ruandao/jvmgo/ch05/instructions/loads"
	"github.com/ruandao/jvmgo/ch05/instructions/math"
	"github.com/ruandao/jvmgo/ch05/rtda"
)

type WIDE struct {
	modifiedInstruction	base.Instruction
}

func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15: // iload
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x16: // lload
	case 0x17: // fload
	case 0x18: // dload
	case 0x19: // aload
	case 0x36: // istore
	case 0x37: // lstore
	case 0x38: // fstore
	case 0x39: // dstore
	case 0x3a: // astore
	case 0x84: // iinc
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0xa9: // ret
		panic("Unsupported opcode: 0xa9!")
	}
}

func (self *WIDE) Execute(frame *rtda.Frame) {
	self.modifiedInstruction.Execute(frame)
}