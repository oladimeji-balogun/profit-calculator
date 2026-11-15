package main

import (
	"fmt"
)

func calculateEBT(revenue, expenses float64) float64 {
	return revenue - expenses
}

func calculateProfit(ebt, taxRate float64) float64 {
	return ebt - ebt*taxRate/100
}

func calculateRatio(ebt, profit float64) float64 {
	return ebt / profit
}

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

	ebt := calculateEBT(revenue, expenses)
	profit := calculateProfit(ebt, taxRate)
	ratio := calculateRatio(ebt, profit)

	fmt.Printf("ebt: %.2f\n", ebt)
	fmt.Printf("profit: %.2f\n", profit)
	fmt.Printf("ratio: %.2f\n", ratio)
}
