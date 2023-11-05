package main

import (
	"fmt"
)

func main() {
	itemOpeningStock := 500
	itemSold := 100
	itemReturned := 50
	itemMissing := 5
	var itemClosingStockCalc int

	itemActualClosingStock := 445
	fmt.Println("Available Iventory (Check): ", itemActualClosingStock)
	fmt.Println("--------------------------------")

	itemClosingStockCalc = finalCalc(itemOpeningStock, itemSold, itemReturned, itemMissing)
	//itemClosingStockCalc = finalCalcPrint(itemOpeningStock, itemSold, itemReturned, itemMissing)
	//itemClosingStockCalc = finalCalcLog(itemOpeningStock, itemSold, itemReturned, itemMissing)
	fmt.Println("Available Iventory (Calc) : ", itemClosingStockCalc)

	fmt.Println("--------------------------------")
	if itemClosingStockCalc != itemActualClosingStock {
		fmt.Println("Warning: Calcualtion ERROR!")
	} else {
		fmt.Println("Great Work: All Good!")
	}
}
