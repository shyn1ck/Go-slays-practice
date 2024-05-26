package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("Введите количество элементов в слайсе:")
	fmt.Scan(&n)
	slice := make([]int, n)
	fmt.Println("Введите элементы слайса:")

	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	fmt.Println("Реверсированный слайс:", slice)
}

//Задача 1. Реверс слайса:
//Напишите программу, которая создает слайс строк и
//реверсирует его, то есть меняет порядок элементов на
//противоположный. Выведите на экран полученный
//реверсированный слайс.
//input: [1, 3, 4, 2, 6]
//output: [6, 2, 4, 3, 1]
