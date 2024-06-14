package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	_, nums := input()
	fmt.Println(solution(nums))
}

func input() (int, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	var nums []string
	for i := 0; i < n; i++ {
		scanner.Scan()
		nums = append(nums, scanner.Text())
	}
	return n, nums
}

// solution сортує масив чисел у порядку неспадання суми їх цифр.
// У разі рівності суми цифр, числа сортуються за значенням.
// Використовується алгоритм сортування вставками. Складність - O(n^2).
//
// Аргументи:
// a ([]string): Масив чисел у вигляді рядків.
//
// Повертає:
// string: Відсортований масив чисел у вигляді одного рядка.
func solution(a []string) string {
	for i := 1; i < len(a); i++ {
		if digitSum(a[i-1]) >= digitSum(a[i]) {
			for j := i; j > 0; j-- {
				// Порівняння за сумою цифр, а потім за значенням числа
				if digitSum(a[j-1]) > digitSum(a[j]) || (digitSum(a[j-1]) == digitSum(a[j]) && toInt(a[j-1]) > toInt(a[j])) {
					a[j], a[j-1] = a[j-1], a[j]
				} else {
					break
				}
			}
		}
	}
	return strings.Join(a, " ")
}

// toInt перетворює рядок у ціле число.
//
// Аргументи:
// s (string): Рядок, що містить ціле число.
//
// Повертає:
// int: Ціле число, отримане з рядка.
func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

// digitSum обчислює суму цифр числа, представленого у вигляді рядка.
//
// Аргументи:
// s (string): Рядок, що містить число.
//
// Повертає:
// int: Сума цифр числа.
func digitSum(s string) int {
	sum := 0
	for _, ch := range s {
		if ch == '-' {
			continue
		}
		digit, err := strconv.Atoi(string(ch))
		if err != nil {
			panic(err)
		}
		sum += digit
	}
	return sum
}
