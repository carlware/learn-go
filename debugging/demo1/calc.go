package main

func finalCalc(itemOpeningStock int, itemSold int, itemReturned int, itemMissing int) int {

	var itemClosingStock int

	itemClosingStock = minus(itemOpeningStock, itemSold)
	itemClosingStock = plus(itemClosingStock, itemReturned)
	itemClosingStock = itemClosingStock - itemMissing

	return (itemClosingStock)
}
