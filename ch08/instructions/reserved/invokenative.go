package reserved

import (
	"github.com/ruandao/jvmgo/ch08/instructions/base"
	"github.com/ruandao/jvmgo/ch08/rtda"
	"github.com/ruandao/jvmgo/ch08/native"
	_ "github.com/ruandao/jvmgo/ch08/native/java/lang"
)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}
	nativeMethod(frame)
}
