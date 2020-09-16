package main

import (
	"fmt"
)

/**
Напишите программу, принимающая на вход число N (N \geq 4)N(N≥4), а затем NN целых чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
*/
func getSlice1() {
	var a int

	fmt.Scan(&a)
	t := make([]int, a, a)

	for i := 0; i < a; i++ {

		fmt.Scan(&t[i])

	}
	fmt.Println(t[3])
}

/**
На ввод подаются пять чисел, которые записываются в массив. Однако эта часть программы уже написана.

Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.
*/
func getSlice2() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	// здесь ваш код
	// ...
	var b int
	for i, v := range array {
		if i == 0 {
			b = v
		}

		if v > b {
			b = v
		}
	}
	fmt.Println(b)
}

/**
Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0. Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).
*/
func getSlice3() {
	var a int
	fmt.Scan(&a)
	t := make([]int, a)
	for i, _ := range t {
		fmt.Scan(&t[i])
		if i%2 == 0 {
			if i != len(t)-1 {
				fmt.Printf("%d ", t[i])
			} else {
				fmt.Printf("%d", t[i])
			}
		}
	}

}

/**
Дана последовательность, состоящая из целых чисел. Напишите программу, которая подсчитывает количество положительных чисел среди элементов последовательности.
*/
func getSlice4() {
	var a, count int
	fmt.Scan(&a)
	t := make([]int, a)
	for i, _ := range t {
		fmt.Scan(&t[i])
		if t[i] > 0 {
			count++
		}
	}

	fmt.Println(count)

}

func main() {
	getSlice4()

}
