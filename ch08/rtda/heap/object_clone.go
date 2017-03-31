package heap

func (self *Object) Clone() *Object {
	return &Object{
		class:self.class,
		data:self.cloneData(),
	}
}

func (self *Object) cloneData() interface{} {
	switch self.data.(type) {
	case []int8:
	case []int16:
	case []uint16:
	case []int32:
	case []int64:
	case []float32:
	case []float64:
	case []*Object:
		elements := self.data.([]*Object)
		elements2 := make([]*Object, len(elements))
		copy(elements2, elements)
		return elements2
	default:
		slots := self.data.(Slots)
		slots2 := newSlots(uint(len(slots)))
		copy(slots2, slots)
		return slots2
	}
}