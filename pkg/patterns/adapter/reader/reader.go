package reader

type text = string

type getter interface {
	Get() (text, error)
}

// Reader is interface, contains method Read
type Reader interface {
	Read() error
}

type human struct {
	document getter
}

func (h *human) Read() (error error) {
	_, error = h.document.Get()
	return
}

// NewReader ...
func NewReader(document getter) Reader {
	return &human{
		document: document,
	}
}
