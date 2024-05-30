package main

import (
	"fmt"
	"math"
)

func main() {
	const expectedReturnRate = 5.5
	investmentAmount := 10000.0
	var years float64

	fmt.Print("Please input investment amout: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Please input years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
}
