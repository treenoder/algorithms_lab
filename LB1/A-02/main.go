package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, nums := input()
	sort(nums, n)
}

// sort сортує масив вставками і виводить стан масиву після кожної вставки.
// Якщо масив вже відсортований, виводить повідомлення "The array is already sorted".
//
// Аргументи:
// a ([]int64): Масив, що підлягає сортуванню.
// n (int): Кількість елементів у масиві.
// Складність:
// O(n^2) - у середньому та у найгіршому випадку.
// O(n) - у найкращому випадку.
func sort(a []int64, n int) {
	isSorted := true
	for i := 1; i < n; i++ {
		if a[i-1] > a[i] {
			isSorted = false
			// якщо знайдено невідсортований елемент
			// вставка елемента у відсортовану частину масиву
			for j := i; j > 0; j-- {
				if a[j-1] > a[j] {
					a[j], a[j-1] = a[j-1], a[j]
				} else {
					break
				}
			}
			// вивід поточного стану масиву
			for _, now := range a {
				fmt.Print(now, " ")
			}
			fmt.Println()
		}
	}
	if isSorted {
		fmt.Println("The array is already sorted")
	}
}

func input() (int, []int64) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	nums := make([]int64, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		nums = append(nums, num)
	}
	return n, nums
}
