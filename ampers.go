package main

import (
	"fmt"
)

/**
Напишите функцию, которая умножает значения на которые ссылаются два указателя и выводит получившееся произведение в консоль. Ввод данных уже написан за вас.
*/
func test1(x1 *int, x2 *int) {
	// здесь пишите ваш код
	fmt.Println(*x1 * *x2)
}

/**
Поменяйте местами значения переменных на которые ссылаются указатели. После этого переменные нужно вывести.
*/
func test(x1 *int, x2 *int) {
	// здесь ваш код
	*x1, *x2 = *x2, *x1
	fmt.Println(*x1, *x2)
}

func main() {
	var a, b int = 2, 4
	test(&a, &b)
}
