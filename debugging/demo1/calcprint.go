package main

import "fmt"

func finalCalcPrint(itemOpeningStock int, itemSold int, itemReturned int, itemMissing int) int {

	var itemClosingStock int
	fmt.Println("itemOpeningStock is ", itemOpeningStock)
	fmt.Println("itemClosingStock is ", itemClosingStock)
	itemClosingStock = minus(itemOpeningStock, itemSold)
	fmt.Println("itemSold is ", itemSold)
	fmt.Println("itemClosingStock is ", itemClosingStock)
	itemClosingStock = plus(itemClosingStock, itemReturned)
	fmt.Println("itemReturned is ", itemReturned)
	fmt.Println("itemClosingStock is ", itemClosingStock)
	itemClosingStock = itemClosingStock - itemMissing
	fmt.Println("itemMissing is ", itemMissing)
	fmt.Println("itemClosingStock is ", itemClosingStock)

	return (itemClosingStock)
}
