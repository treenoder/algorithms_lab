package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, nums := input()
	fmt.Println(solution(nums, n))
}

// solution знаходить кількість різних чисел у заданому масиві.
//
// Аргументи:
// a ([]int): Масив цілих чисел.
// n (int): Кількість елементів у масиві.
//
// Повертає:
// int: Кількість різних чисел у масиві.
func solution(a []int, n int) int {
	// Використовуємо map для зберігання унікальних елементів
	elements := make(map[int]struct{})
	for i := 0; i < n; i++ {
		elements[a[i]] = struct{}{}
	}
	// Кількість унікальних елементів - це розмір map
	return len(elements)
}

func input() (int, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return n, nums
}
