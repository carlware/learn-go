package main

func finalStockCalc(openingStock int, sold int, returned int) int {

	closingStock := minus(openingStock, sold)
	closingStock = plus(closingStock, returned)

	//closingStock := minus(plus(openingStock, sold), returned)

	return (closingStock)
}

func finalCostCalc(closingStockCount int, cost int) int {

	closingStockCost := multi(closingStockCount, cost)

	return (closingStockCost)
}
