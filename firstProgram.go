package main

import (
	"fmt"
	"strconv"
)

/**
Напишите программу, которая выводит "I like Go!"
*/
func firstProgramm() {
	fmt.Println("I like Go!")
}

func secondProgramm() {
	for i := 0; i < 3; i++ {
		fmt.Println("I like Go!")
	}
}

func thirdProgramm() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a*2 + 100)
}

/**
По данному целому числу, найдите его квадрат.
*/
func squadProgramm() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a * a)
}

/**
Дано натуральное число, выведите его последнюю цифру.
*/
func getLastCifr() {
	var a string
	fmt.Scan(&a)
	fmt.Println(string(a[len(a)-1]))
}

/**
Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа).
*/
func getPreLastCifr() {
	var a string
	fmt.Scan(&a)
	fmt.Println(string(a[len(a)-2]))
}

/**
Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.
*/
func getTimeFromGrade() {
	var a int
	fmt.Scan(&a)
	fmt.Println("It is", a/30, "hours", a%30/5, "minutes.")
}
func getFiletrNumber() {

	var (
		a      int
		b      int
		result int = 0
	)

	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if len(strconv.Itoa(b)) == 2 && b%8 == 0 {
			result += b
		}
	}

	fmt.Println(result)

}
func main() {
	getFiletrNumber()
}
