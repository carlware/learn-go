package main

import (
	"log"
	"os"
)

func finalCalcLog(itemOpeningStock int, itemSold int, itemReturned int, itemMissing int) int {

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	var itemClosingStock int
	log.Println("itemClosingStock is ", itemClosingStock)
	itemClosingStock = minus(itemOpeningStock, itemSold)
	log.Println("itemSold is ", itemSold)
	log.Println("itemClosingStock is ", itemClosingStock)
	itemClosingStock = plus(itemClosingStock, itemReturned)
	log.Println("itemReturned is ", itemReturned)
	log.Println("itemClosingStock is ", itemClosingStock)
	itemClosingStock = itemClosingStock - itemMissing
	log.Println("itemMissing is ", itemMissing)
	log.Println("itemClosingStock is ", itemClosingStock)

	return (itemClosingStock)
}
