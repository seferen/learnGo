package main

import (
	"fmt"
)

/**
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.
*/
func minimumFromFour() int {
	//print your code here
	var a, result int
	for i := 0; i < 4; i++ {
		fmt.Scan(&a)
		if i == 0 {
			result = a
		} else if result > a {
			result = a

		}
	}

	return result
}

/**
Напишите "функцию голосования", возвращающую то значение (0 или 1), которое среди значений ее аргументов x, y, z встречается чаще.
*/
func vote(x int, y int, z int) int {
	//print your code here
	return 0

}

/**
Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.
*/
func sumInt(a ...int) (count int, sum int) {
	for _, v := range a {
		sum += v
	}
	count = len(a)
	return count, sum
}

func main() {
	fmt.Println(sumInt(1, 2, 3, 4))
}
