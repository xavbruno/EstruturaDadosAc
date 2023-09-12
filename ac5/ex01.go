package main

import "fmt"

func acharPar(arr []int, target int) (int, int) {
	left, right := 0, len(arr)-1

	for left < right {
		sum := arr[left] + arr[right]

		switch {
		case sum == target:
			return arr[left], arr[right]
		case sum < target:
			left++
		default:
			right--
		}
	}

	return -1, -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 10

	num1, num2 := acharPar(arr, target)

	if num1 != -1 && num2 != -1 {
		fmt.Printf("Par encontrado: %d + %d = %d\n", num1, num2, target)
	} else {
		fmt.Println("Nenhum par encontrado.")
	}
}
