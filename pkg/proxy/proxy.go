package proxy

type Proxy interface {
	SetData(int)
	GetData() int
}

// Proxy contains reference on a realObject
type proxy struct {
	object Proxy
}

// SetData sets a value for realObject data
func (p *proxy) SetData(i int) {
	subject := p.getRealObject()
	subject.SetData(i)
}

// GetData returns data of realObject
func (p *proxy) GetData() int {
	subject := p.getRealObject()
	return subject.GetData()
}

func (p *proxy) getRealObject() Proxy {
	if p.object == nil {
		p.object = newRealObject()
	}
	return p.object
}

func NewProxy() Proxy {
	return &proxy{}
}
