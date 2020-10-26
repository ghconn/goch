package main

import (
	"fmt"
	"math/big"
)

func main() {
	distance := new(big.Int)
	distance.SetString("240000000000000000000", 10)
	fmt.Println(distance)

	lightSpeed := big.NewInt(200792)
	secondsPerDay := big.NewInt(86400)

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)

	fmt.Println("That is", days, "days of travel at light speeds.")
}
