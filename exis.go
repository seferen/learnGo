package main

import (
	"fmt"
	"strconv"
)

/**
Дано трехзначное число. Найдите сумму его цифр.
*/
func ex1() {
	var a, result int
	fmt.Scan(&a)
	for _, v := range strconv.Itoa(a) {
		r, _ := strconv.Atoi(string(v))
		result += r

	}

	fmt.Println(result)
}

/**
Дано трехзначное число. Переверните его, а затем выведите.
*/
func ex2() {
	var a string
	fmt.Scan(&a)
	runes := []rune(a)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println(string(runes))

}

/**
Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток. Например, если
k=13257=3*3600+40*60+57,
*/
func ex3() {

	var a int32
	fmt.Scan(&a)
	fmt.Println("It is", a/3600, "hours", a%3600/60, "minutes.")

}

/**
Заданы три числа - a, b, c (a < b < c)a,b,c(a<b<c) - длины сторон треугольника. Нужно проверить, является ли треугольник прямоугольным. Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"
*/
func ex4() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if c*c == a*a+b*b {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}

func main() {
	ex4()

}
