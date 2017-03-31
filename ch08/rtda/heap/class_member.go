package heap

import "github.com/ruandao/jvmgo/ch08/classfile"

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

func (self *ClassMember) Name() string {
	return self.name
}
func (self *ClassMember) Descriptor() string {
	return  self.descriptor
}