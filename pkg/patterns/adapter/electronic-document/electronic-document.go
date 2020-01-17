package electronic_document

import "github.com/Yemanden/Learning/pkg/models"

type decimalData = []byte

// Getter ...
type Getter interface {
	Get() (decimalData, error)
}

type document struct {
	decimalData decimalData
}

func (d *document) Get() (decimalData decimalData, error error) {
	if d.decimalData == nil {
		error = models.DocumentMissingDataError
		return
	}
	decimalData = d.decimalData
	return
}

// NewElectronicDocument returns a new electronic document
func NewElectronicDocument(decimalData decimalData) Getter {
	return &document{
		decimalData: decimalData,
	}
}
