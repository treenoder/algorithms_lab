package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	k := input()
	fmt.Println(Solution(k))
	//for i := 1; i <= 20; i++ {
	//	println(Solution(i))
	//}
}

func Solution(index int) int64 {
	var sol int64
	var count int
	j := int64(1)
	for i := int64(1); i <= int64(index) && count < index; i++ {
		cube := i * i * i
		for ; j*j <= cube && count < index; j++ {
			if j*j == cube {
				j++
				break
			}
			sol = j * j
			count++
		}
		if count < index {
			sol = cube
			count++
		} else {
			break
		}
	}

	return sol
}

func input() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
