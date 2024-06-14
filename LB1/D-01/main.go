package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	_, nums := input()
	//nums := []int{5, 4, 3, 2, 1, 5, 8, 9}
	fmt.Println(Solution(nums))
}

func input() (int, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	var nums []int
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return n, nums
}

// Solution сортує масив цілих чисел у порядку неубування за допомогою швидкого сортування.
// Якщо масив вже відсортований, повертає його без змін.
// Cкладність - O(n*log(n)) у середньому та у кращому випадку, де n - кількість елементів у масиві.
// У найгіршому випадку складність - O(n^2).
//
// Аргументи:
// a ([]int): Масив цілих чисел.
//
// Повертає:
// string: Відсортований масив у вигляді рядка.
func Solution(a []int) string {
	if !isSorted(a) {
		// Перемішування масиву для уникнення найгіршого випадку для швидкого сортування
		shuffle(a)
		quickSort(a, 0, len(a)-1)
	}
	return fmt.Sprintf("%v", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), " "), "[]"))
}

// quickSort реалізує алгоритм швидкого сортування (QuickSort).
//
// Аргументи:
// a ([]int): Масив цілих чисел.
// low (int): Початковий індекс для сортування.
// high (int): Кінцевий індекс для сортування.
func quickSort(a []int, low, high int) {
	if low < high {
		// Отримання індексу розділення
		p := partition(a, low, high)
		// Рекурсивний виклик для лівої частини
		quickSort(a, low, p-1)
		// Рекурсивний виклик для правої частини
		quickSort(a, p+1, high)
	}
}

// partition ділить масив на дві частини навколо опорного елемента.
//
// Аргументи:
// a ([]int): Масив цілих чисел.
// low (int): Початковий індекс для розділення.
// high (int): Кінцевий індекс для розділення.
//
// Повертає:
// int: Індекс опорного елемента після розділення.
func partition(a []int, low, high int) int {
	pivot := a[high]
	i := low - 1

	for j := low; j < high; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}

	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}

// shuffle перемішує масив для уникнення найгіршого випадку для швидкого сортування.
//
// Аргументи:
// a ([]int): Масив цілих чисел.
func shuffle(a []int) {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
}

// isSorted перевіряє, чи масив вже відсортований у порядку неубування.
//
// Аргументи:
// a ([]int): Масив цілих чисел.
//
// Повертає:
// bool: true, якщо масив відсортований, і false в іншому випадку.
func isSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}
