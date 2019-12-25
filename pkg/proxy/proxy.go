package proxy

// Proxy is interface, contains methods SetData and GetData
type Proxy interface {
	SetData(int)
	GetData() int
}

// Proxy contains reference on a realObject
type proxy struct {
	data   int
	object Proxy
}

// SetData sets a value for realObject data
func (p *proxy) SetData(i int) {
	subject := p.getRealObject()
	subject.SetData(i)
}

// GetData returns data of realObject
func (p *proxy) GetData() int {
	object := p.getRealObject()
	if p.data == 0 {
		return object.GetData()
	}
	return p.data
}

func (p *proxy) getRealObject() Proxy {
	if p.object == nil {
		p.object = newRealObject()
	}
	return p.object
}

// NewProxy is constructor. Returns a new object of proxy structure, implemented Proxy interface
func NewProxy() Proxy {
	return &proxy{}
}
