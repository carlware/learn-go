package main

import (
	"bufio"
	"crl/grocery/internal/grocery"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)



func main() {
	filePath := flag.String("input", "input.txt", "file input")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("failed to open %v", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	idx := 0
	var totalRs int
	customers := []*grocery.Customer{}

	// TODO: abstract this an create an interface
	for scanner.Scan() {
		text := scanner.Text()
		if idx == 0 {
			totalRs, err = strconv.Atoi(text)
			if err != nil {
				log.Fatalf("convert to int %v", err)
			}
		} else {
			tokens := strings.Split(text, " ")
			if len(tokens) != 3 {
				log.Fatalf("input format error")
			}
			kind := tokens[0]
			arrivalTime, err := strconv.Atoi(tokens[1])
			if err != nil {
				log.Fatalf("convert to int %v", err)
			}
			items, err := strconv.Atoi(tokens[2])
			if err != nil {
				log.Fatalf("convert to int %v", err)
			}
			c, err := grocery.CreateCustomer(kind, arrivalTime, items)
			if err != nil {
				log.Fatalf("create cusotmer err=%v", err)
			}
			customers = append(customers, c)
		}
		idx++
	}

	total := grocery.CalculateTotalTime(totalRs, customers)
	fmt.Printf("Finished at: t=%d minutes", total)
}