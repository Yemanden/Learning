package main

import (
	"fmt"
	"github.com/Yemanden/Learning/pkg/patterns/adapter/converter"
	eDocument "github.com/Yemanden/Learning/pkg/patterns/adapter/electronic-document"
	pDocument "github.com/Yemanden/Learning/pkg/patterns/adapter/paper-document"
	"github.com/Yemanden/Learning/pkg/patterns/adapter/reader"
	"github.com/Yemanden/Learning/pkg/patterns/adapter/screen"
)

func main() {
	paperDocument := pDocument.NewPaperDocument("Привет")
	human := reader.NewReader(paperDocument)
	error1 := human.Read()
	if error1 == nil {
		fmt.Println("Paper read. Done!")
	}

	elDocument := eDocument.NewElectronicDocument([]byte{1, 2, 3, 3, 4})
	converter := converter.NewConverter()
	screen := screen.NewScreen(elDocument, converter)
	human2 := reader.NewReader(screen)
	error2 := human2.Read()
	if error2 == nil {
		fmt.Println("Electronic read. Done!")
	}
}
