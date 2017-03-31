package instructions

import "github.com/ruandao/jvmgo/ch05/instructions/base"

import . "github.com/ruandao/jvmgo/ch05/instructions/comparisons"
import . "github.com/ruandao/jvmgo/ch05/instructions/constants"
import . "github.com/ruandao/jvmgo/ch05/instructions/control"
import . "github.com/ruandao/jvmgo/ch05/instructions/conversions"
import . "github.com/ruandao/jvmgo/ch05/instructions/extended"
import . "github.com/ruandao/jvmgo/ch05/instructions/loads"
import . "github.com/ruandao/jvmgo/ch05/instructions/math"
import . "github.com/ruandao/jvmgo/ch05/instructions/stack"
import (
	. "github.com/ruandao/jvmgo/ch05/instructions/stores"
	"fmt"
)

var (
	nop	= &NOP{}
	aconst_null = &ACONST_NULL{}
)
func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconst_null
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}
