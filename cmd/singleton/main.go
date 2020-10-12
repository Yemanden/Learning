package main

import (
	"fmt"

	"github.com/Yemanden/Learning/pkg/models"
	"github.com/Yemanden/Learning/pkg/patterns/singleton/buyer"
	"github.com/Yemanden/Learning/pkg/patterns/singleton/shop"
)

func main() {
	human1 := buyer.NewBuyer()
	human2 := buyer.NewBuyer()

	for i := 0; i < 6; i++ {
		error := human1.Buy(models.TShirt, shop.GetSeller())
		if error == nil {
			fmt.Println("Human1 Buy. Done!")
		} else {
			fmt.Println("Human1 Buy. " + error.Error())
		}
	}

	for i := 0; i < 6; i++ {
		error := human2.Buy(models.TShirt, shop.GetSeller())
		if error == nil {
			fmt.Println("Human2 Buy. Done!")
		} else {
			fmt.Println("Human2 Buy. " + error.Error())
		}
	}
}
