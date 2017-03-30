package heap

import "github.com/ruandao/jvmgo/ch06/classfile"

type ClassMember struct {
	accessFlags	uint16
	name 		string
	descriptor	string
	class 		*Class
}

func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) Class() *Class {
	return self.class
}