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
func (p *Proxy) SetData(i int) {
	subject := p.getRealObject()
	subject.setData(i)
}

func (s *realSubject) setData(i int) {
	s.data = i
}

// GetData returns data of realObject
func (p *Proxy) GetData() int {
	subject := p.getRealObject()
	return subject.getData()
}

func (s *realSubject) getData() int {
	return s.data
}

func (p *Proxy) getRealObject() *realSubject {
	if p.object == nil {
		p.object = &realSubject{}
	}
	return p.object
}
