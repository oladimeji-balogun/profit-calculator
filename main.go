package main

import (
	"fmt"
)

func main() {

	var revenue float64
	fmt.Print("enter the revenue: ")
	fmt.Scan(&revenue)

	var expenses float64
	fmt.Print("enter the expenses here: ")
	fmt.Scan(&expenses)

	var taxRate float64
	fmt.Print("enter the tax rate here: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt - (ebt * taxRate / 100)
	ratio := ebt / profit

	fmt.Printf("ebt: %.2f\n", ebt)
	fmt.Printf("profit: %.2f\n", profit)
	fmt.Printf("ratio: %.2f\n", ratio)
}
