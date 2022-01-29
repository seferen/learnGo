package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

/*
Напишите функцию f(), которая будет принимать строку text и выводить ее (печатать на экране).
*/
func f(text string) {
	//print your code here
	fmt.Println(text)
}

/*
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

/*
Напишите "функцию голосования", возвращающую то значение (0 или 1), которое среди значений ее аргументов x, y, z встречается чаще.
*/
func vote(x int, y int, z int) int {
	//print your code here
	if (x == 0 && y == 0) || (x == 0 && z == 0) {
		return 0

	}
	return 1

}

/*
Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1. Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ... Напишите функцию, которая по указанному натуральному n возвращает φn.
*/
func fibonacci(n int) int {
	//print your code here
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fibonacci(n-2) + fibonacci(n-1)
}

// Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.
func sumInt(a ...int) (int, sum int) {
	for _, v := range a {
		sum += v
	}

	return len(a), sum
}

// Напишите функцию, которая умножает значения на которые ссылаются два указателя и выводит получившееся произведение в консоль. Ввод данных уже написан за вас.
func function1(x1 *int, x2 *int) {
	// здесь пишите ваш код
	fmt.Println(*x1 * *x2)
}

// Поменяйте местами значения переменных на которые ссылаются указатели. После этого переменные нужно вывести.
func function2(x1 *int, x2 *int) {
	// здесь ваш код
	*x1, *x2 = *x2, *x1
	fmt.Println(*x1, *x2)
}

/*
Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно.
У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.
Если значение On == false, то оба метода вернут false.
Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true), если его нет,
то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.
*/
type Test struct {
	On    bool
	Ammo  int
	Power int
}

func (ts *Test) Shoot() bool {
	if ts.On == false {
		return false
	}
	if ts.Ammo > 0 {
		ts.Ammo--
		return true
	} else {
		return false
	}

}
func (ts *Test) RideBike() bool {
	if ts.On == false {
		return false
	}
	if ts.Power > 0 {
		ts.Power--
		return true
	} else {
		return false
	}

}

/* "StringTest ..."
На вход подается строка. Нужно определить, является ли она правильной или нет. Правильная строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная - вывести Right иначе - вывести Wrong
*/
func StringTest() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	test := []rune(text)
	if unicode.IsUpper(test[0]) && strings.HasSuffix(text, ".") {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}

// На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является паллиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.

func polindrom(line string) string {
	lineR := []rune(strings.ToLower(line))

	for i, j := 0, len(lineR)-1; i < j; i, j = i+1, j-1 {

		if lineR[i] != lineR[j] {
			return "Нет"
		}

	}
	return "Палиндром"
}

// Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в строке X. Если подстроки S нет в строке X - вывести -1
func function3(line, subline string) int {
	return strings.Index(line, subline)

}

// На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)
func function4(line *string) string {
	resultR := make([]rune, 0)
	for i, v := range []rune(*line) {
		if i%2 == 1 {
			resultR = append(resultR, v)
		}

	}
	return string(resultR)
}

// Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку
func function5(line *string) {
	for _, v := range []rune(*line) {
		if strings.Count(*line, string(v)) > 1 {
			*line = strings.ReplaceAll(*line, string(v), "")
		}
	}
}

// Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
// Длина пароля должна быть не менее 5 символов, он должен содержать только цифры и буквы латинского алфавита.
// На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
func function6(pas *string) {
	result, err := regexp.Match(`^[\w|\d]{5,}$`, []byte(*pas))
	if err != nil {
		fmt.Println(err)
	}
	if result {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}

}

// На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы
func function7(a, b *float64) int {

	return int(math.Sqrt(*a**a + *b**b))
}

// Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).
func function8(line *string) string {
	return strings.Join(strings.Split(*line, ""), "*")
}

// Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.
func function9(line *string) (max int) {

	for i, v := range []rune(*line) {
		r, _ := strconv.Atoi(string(v))
		if i == 0 || r > max {
			max = r
		}

	}
	return max
}

// На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.
func function10(a *int) {
	for _, v := range []rune(strconv.Itoa(*a)) {
		vInt, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d", vInt*vInt)
	}
	fmt.Println()

}

func main() {

}
