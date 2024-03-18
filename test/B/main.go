package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	price := (a * 100) + b
	funds := (c * 100) + d
	count := funds / price
	left := funds - count*price
	leftCents := left % 100
	leftDollars := left / 100
	fmt.Printf("%d\n", count)
	fmt.Printf("%d %d", leftDollars, leftCents)
}
