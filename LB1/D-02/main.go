package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	k := input()
	fmt.Println(solution(k))
	//for i := 1; i <= 20; i++ {
	//	println(Solution(i))
	//}
}

// solution обчислює k-й елемент об'єднаної та відсортованої послідовності C
// з послідовностей A та B, де A складається з квадратів натуральних чисел, а B
// складається з кубів натуральних чисел.
// Складність - O(k^2), де k - номер шуканого елемента.
func solution(index int) int64 {
	// Змінна для збереження k-го елемента послідовності C
	var sol int64

	// Лічильник для відстеження елементів у послідовності C
	var count int

	// Змінна для ітерації через послідовність A
	j := int64(1)

	for i := int64(1); i <= int64(index) && count < index; i++ {
		// Обчислення поточного куба для послідовності B
		cube := i * i * i

		// Ітерація через послідовність A та об'єднання з послідовністю B
		for ; j*j <= cube && count < index; j++ {
			// Якщо квадрат дорівнює поточному кубу, пропустити його
			if j*j == cube {
				// Перехід до наступного квадрата
				j++
				break
			}
			// Додавання поточного квадрата до послідовності C
			sol = j * j
			count++
		}
		if count < index {
			// Додавання поточного куба до послідовності C
			sol = cube
			count++
		} else {
			break
		}
	}

	// Повернення k-го елемента послідовності C
	return sol
}

func input() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
