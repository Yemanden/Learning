package adapter

import (
	"github.com/Yemanden/Learning/pkg/patterns/adapter/converter"
	"github.com/Yemanden/Learning/pkg/patterns/adapter/screen"
	"testing"

	"github.com/stretchr/testify/assert"

	elDocument "github.com/Yemanden/Learning/pkg/patterns/adapter/electronic-document"
	"github.com/Yemanden/Learning/pkg/patterns/adapter/paper-document"
	"github.com/Yemanden/Learning/pkg/patterns/adapter/reader"
)

const (
	testAdapterPass1Name  = "Using without Adapter"
	testAdapterPass1Input = "Testing text"

	testAdapterPass2Name = "Using with Adapter"

	testAdapterPass3Name  = "Read empty text"
	testAdapterPass3Input = ""
)

var (
	testAdapterPass2Input = []byte{1, 2, 3, 3, 4}
)

func TestAdapter(t *testing.T) {

	t.Run(testAdapterPass1Name, func(t *testing.T) {
		paperDoc := paper_document.NewPaperDocument(testAdapterPass1Input)
		human := reader.NewReader(paperDoc)
		error := human.Read()

		assert.NoError(t, error)
	})

	t.Run(testAdapterPass2Name, func(t *testing.T) {
		elDoc := elDocument.NewElectronicDocument(testAdapterPass2Input)
		converter := converter.NewConverter()
		adapter := screen.NewScreen(elDoc, converter)
		human := reader.NewReader(adapter)
		error := human.Read()

		assert.NoError(t, error)
	})

	t.Run(testAdapterPass3Name, func(t *testing.T) {
		paperDoc := paper_document.NewPaperDocument(testAdapterPass3Input)
		human := reader.NewReader(paperDoc)
		error := human.Read()

		assert.Error(t, error)
	})
}
