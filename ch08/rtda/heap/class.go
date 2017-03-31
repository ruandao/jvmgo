package heap

import (
	"github.com/ruandao/jvmgo/ch08/classfile"
	"strings"
)

type Class struct {
	accessFlags		uint16
	name			string
	superClassName		string
	interfaceNames		[]string
	constantPool		*ConstantPool
	fields			[]*Field
	methods			[]*Method
	loader			*ClassLoader
	superClass		*Class
	interfaces		[]*Class
	instanceSlotCount	uint
	staticSlotCount		uint
	staticVars		*Slots
	initStarted		bool
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(self.name)
	return self.loader.LoadClass(componentClassName)
}

func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := self; c!= nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic &&
			field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}


func (self *Class) IsPublic() bool {
	return 0 != (self.accessFlags & ACC_PUBLIC)
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}
func (self *Class) Name() string {
	return self.name
}

func (self *Class) ArrayClass() *Class {
	arrayClassName  := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}

func (self *Class) IsAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.getPackageName == other.getPackageName()
}

func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}
func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
		   method.name == method.descriptor == descriptor {
			return method
		}
	}
	return nil
}

func (self *Class) InitStarted() bool {
	return self.initStarted
}

func (self *Class) StartInit() {
	self.initStarted = true
}

func (self *Class) GetClinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}