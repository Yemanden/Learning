package proxy

import "github.com/Yemanden/Learning/pkg/proxy/real_object"

// Proxer is interface, contains interfaces Setter and Getter
type Proxer interface {
	Setter
	Getter
}

// Setter ...
type Setter interface {
	SetData(int)
}

// Getter ...
type Getter interface {
	GetData() int
}

// proxy contains reference on a realObject
type proxy struct {
	data   int
	object Proxer
}

// SetData sets a value for realObject data
func (p *proxy) SetData(i int) {
	p.getRealObject().SetData(i)
}

// GetData returns data of realObject
func (p *proxy) GetData() int {
	if p.data == 0 {
		return p.getRealObject().GetData()
	}
	return p.data
}

func (p *proxy) getRealObject() Proxer {
	if p.object == nil {
		p.object = realobject.NewRealObject()
	}
	return p.object
}

// NewProxy is constructor. Returns a new object of proxy structure, implemented Proxy interface
func NewProxy() Proxer {
	return &proxy{}
}
