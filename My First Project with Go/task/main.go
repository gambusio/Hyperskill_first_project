package main

import "fmt"

/* Item prices
const (
	bubblegumPrice     = 2.0
	toffeePrice        = 0.2
	icecreamPrice      = 5.0
	milkchocolatePrice = 4.0
	doughnutPrice      = 2.5
	pancakePrice       = 3.2
)*/

const (
	bubblegumTotalSales     = 202.0
	toffeeTotalSales        = 118
	icecreamTotalSales      = 2250
	milkchocolateTotalSales = 1680
	doughnutTotalSales      = 1075
	pancakeTotalSales       = 80
)

func main() {
	var income float32 = bubblegumTotalSales + toffeeTotalSales + icecreamTotalSales + milkchocolateTotalSales + doughnutTotalSales + pancakeTotalSales
	var staffExpenses float32
	var otherExpenses float32
	var netIncome float32

	fmt.Println("Earned amount:")
	fmt.Print("Bubblegum: $", bubblegumTotalSales, "\n")
	fmt.Print("Toffee: $", toffeeTotalSales, "\n")
	fmt.Print("Ice cream: $", icecreamTotalSales, "\n")
	fmt.Print("Milk chocolate: $", milkchocolateTotalSales, "\n")
	fmt.Print("Doughnut: $", doughnutTotalSales, "\n")
	fmt.Print("Pancake: $", pancakeTotalSales, "\n")

	fmt.Println("")
	fmt.Print("Income: $", income, "\n")

	fmt.Print("Staff expenses:\n")
	fmt.Scanf("%f", &staffExpenses)
	fmt.Print("Other expenses:\n")
	fmt.Scanf("%f", &otherExpenses)
	netIncome = income - staffExpenses - otherExpenses
	fmt.Printf("Net income: $%.0f\n", netIncome)
}
