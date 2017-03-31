package base

import (
	"github.com/ruandao/jvmgo/ch08/rtda"
	"github.com/ruandao/jvmgo/ch08/rtda/heap"
)

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argSlotSlot := int(method.ArgSlotCount())
	if argSlotSlot > 0 {
		for i := argSlotSlot - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack.PopSlot()
			newFrame.LocalVars.SetSlot(uint(i), slot)
		}
	}
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
		}
	}

}
