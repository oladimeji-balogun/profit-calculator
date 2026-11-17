package main

import (
	"errors"
	"fmt"
	"os"
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

func getUserInput(message string) float64 {
	var output float64
	fmt.Print(message)
	_, err := fmt.Scan(&output)
	if err != nil {
		fmt.Println(err, "please try again")
		getUserInput(message)

	}
	return output
}

func handleLessOrEqualToZero(amount float64) {
	if amount <= 0 {
		err := errors.New("expenses cannot be less than or equal to zero. please try again")
		fmt.Println(err)
		return
	}
}

func saveValuesToFile(fileName string, ebt float64, profit float64, ratio float64) {
	text := fmt.Sprintf("ebt: %.2f | profit: %.2f | ratio: %.2f", ebt, profit, ratio)

	os.WriteFile(fileName, []byte(text), os.FileMode(os.O_APPEND))

}
func main() {

	revenue := getUserInput("enter your revenue:")
	if revenue <= 0 {
		handleLessOrEqualToZero(revenue)
	}

	expenses := getUserInput("enter you expenses here:")
	if expenses <= 0 {
		handleLessOrEqualToZero((expenses))
	}

	taxRate := getUserInput("enter your tax rate here:")
	if taxRate <= 0 {
		handleLessOrEqualToZero((taxRate))
	}

	ebt := calculateEBT(revenue, expenses)
	profit := calculateProfit(ebt, taxRate)
	ratio := calculateRatio(ebt, profit)

	saveValuesToFile("file.txt", ebt, profit, ratio)
	fmt.Printf("ebt: %.2f\n", ebt)
	fmt.Printf("profit: %.2f\n", profit)
	fmt.Printf("ratio: %.2f\n", ratio)
}
