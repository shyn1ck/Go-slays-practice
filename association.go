package main

import (
	"fmt"
)

func main() {
	var n1, n2 int
	fmt.Print("Введите размер первого слайса: ")
	fmt.Scan(&n1)
	fmt.Print("Введите размер второго слайса: ")
	fmt.Scan(&n2)

	slice1 := make([]string, n1)
	fmt.Println("Введите элементы первого слайса:")
	for i := 0; i < n1; i++ {
		fmt.Scan(&slice1[i])
	}

	slice2 := make([]string, n2)
	fmt.Println("Введите элементы второго слайса:")
	for i := 0; i < n2; i++ {
		fmt.Scan(&slice2[i])
	}

	countMap := make(map[string]int)

	for _, item := range slice1 {
		countMap[item]++
	}

	for _, item := range slice2 {
		countMap[item]++
	}

	var resultSlice []string
	for item, count := range countMap {
		if count == 1 {
			resultSlice = append(resultSlice, item)
		}
	}
	fmt.Println("Результирующий слайс без дубликатов:", resultSlice)
}

//Задача 2. Объединение слайсов без дубликатов:
//Напишите программу, которая создает два слайса строк
//и объединяет их без дубликатов, то есть исключая
//повторяющиеся элементы. Выведите на экран результат
//объединения.
//input: [1, 2, 3], [3, 4, 5]
//output: [1, 2, 4, 5]
