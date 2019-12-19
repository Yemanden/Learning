package proxy

// realObject is a real object
type realSubject struct {
	data int
}

// Proxy contains reference on a realObject
type Proxy struct {
	object *realSubject
}

// SetData sets a value for realObject data
func (obj *Proxy) SetData(i int) {
	subject := obj.getRealObject()
	subject.setData(i)
}

func (obj *realSubject) setData(i int) {
	obj.data = i
}

// GetData returns data of realObject
func (obj *Proxy) GetData() int {
	subject := obj.getRealObject()
	return subject.getData()
}

func (obj *realSubject) getData() int {
	return obj.data
}

func (obj *Proxy) getRealObject() *realSubject {
	if obj.object == nil {
		obj.object = &realSubject{}
	}
	return obj.object
}
