package constants

import (
	"github.com/ruandao/jvmgo/ch05/instructions/base"
	"github.com/ruandao/jvmgo/ch05/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame)  {

}
