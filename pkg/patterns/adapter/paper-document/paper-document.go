package paper_document

import "github.com/Yemanden/Learning/pkg/models"

type text = string

// Getter ...
type Getter interface {
	Get() (text, error)
}

type document struct {
	text string
}

func (d *document) Get() (text text, error error) {
	if d.text == "" {
		error = models.DocumentMissingDataError
		return
	}
	text = d.text
	return
}

// NewPaperDocument returns a new paper document
func NewPaperDocument(text string) Getter {
	return &document{
		text: text,
	}
}
