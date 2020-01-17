package screen

type (
	text        = string
	decimalData = []byte
)

type converter interface {
	Convert(decimalData) text
}

type electronicDocument interface {
	Get() (decimalData, error)
}

// Getter ...
type Getter interface {
	Get() (text, error)
}

type screen struct {
	electronicDocument electronicDocument
	converter          converter
}

func (s *screen) Get() (text text, error error) {
	var decimalData decimalData
	decimalData, error = s.electronicDocument.Get()
	if error != nil {
		return
	}
	text = s.converter.Convert(decimalData)
	return
}

// NewScreen ...
func NewScreen(electronicDocument electronicDocument, converter converter) Getter {
	return &screen{
		electronicDocument: electronicDocument,
		converter:          converter,
	}
}
