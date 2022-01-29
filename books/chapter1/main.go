package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/**
Напишите программу, которая выводит "I like Go!"
*/
func function1() {
	fmt.Println("I like Go!")
}

/**
Напишите программу, которая выведет "I like Go!" 3 раза.
*/
func function2() {
	for i := 0; i < 3; i++ {
		fmt.Println("I like Go!")
	}
}

/**
Напишите программу, которая последовательно делает следующие операции с введённым числом:

1. Число умножается на 2;
2. Затем к числу прибавляется 100.
*/
func function3() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a*2 + 100)
}

/**
По данному целому числу, найдите его квадрат.
*/
func function4() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a * a)
}

/**
Дано натуральное число, выведите его последнюю цифру.
*/
func function5() {
	var a string
	fmt.Scan(&a)
	fmt.Println(string(a[len(a)-1]))
}

/**
Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа).
*/
func function6() {
	var a string
	fmt.Scan(&a)
	fmt.Println(string(a[len(a)-2]))
}

/**
Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.
*/
func function7() {
	var a int
	fmt.Scan(&a)
	fmt.Println("It is", a/30, "hours", a%30/5, "minutes.")
}

/**
На ввод подается целое число. Если число положительное - вывести сообщение "Число положительное", если число отрицательное -
"Число отрицательное". Если подается ноль - вывести сообщение "Ноль". Выводить сообщение без кавычек.
*/
func function8() {
	var num int
	if fmt.Scan(&num); num == 0 {
		fmt.Println("Ноль")
	} else if num < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Число положительное")
	}
}

/**
По данному трехзначному числу определите, все ли его цифры различны.
*/
func function9() {
	var num string

	if fmt.Scan(&num); num[0] == num[1] || num[0] == num[2] || num[1] == num[2] {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}

/**
Дано неотрицательное целое число. Найдите и выведите первую цифру числа.
*/
func function10() {
	var num string

	fmt.Scan(&num)
	fmt.Println(string(num[0]))
}

/**
Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
*/
func function11() {
	var num string
	fmt.Scan(&num)
	i1, _ := strconv.Atoi(string(num[0]))
	i2, _ := strconv.Atoi(string(num[1]))
	i3, _ := strconv.Atoi(string(num[2]))
	i4, _ := strconv.Atoi(string(num[3]))
	i5, _ := strconv.Atoi(string(num[4]))
	i6, _ := strconv.Atoi(string(num[5]))

	if i1+i2+i3 == i4+i5+i6 {
		fmt.Println("YES")

	} else {
		fmt.Println("NO")
	}
}

/**
Требуется определить, является ли данный год високосным, напомним:
Год является високосным если он соответствует хотя бы одному из нижеперечисленных условий:
- кратен 400;
- кратен 4, но не кратен 100.
*/
func function12() {
	var num int
	if fmt.Scan(&num); num%400 == 0 || (num%4 == 0 && num%100 != 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

/**
Напишите программу, которая выводит квадраты натуральных чисел от 1 до 10. Квадрат каждого числа должен выводится в новой строке.
*/
func function13() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i * i)
	}
}

/**
Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B).
Вывести сумму всех чисел от A до B  включительно.
*/
func function14() {
	var a, b int
	count := 0

	for fmt.Scan(&a, &b); a <= b; a++ {
		count += a
	}

	fmt.Println(count)
}

/**
Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8. Программа в первой строке получает
на вход число n - количество чисел в последовательности, во второй строке -- n чисел, входящих в данную последовательность.
*/
func function15() {
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

/**
Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.
*/
func function16() {
	var a, count, max int

	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		if a > max {
			max = a
			count = 1
		} else if a == max {
			count++
		}

	}
	fmt.Println(count)
}

/**
Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.
*/
func function17() {
	var n, c, d int

	fmt.Scan(&n, &c, &d)

	for i := 0; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}
}

/**
Напишите программу, которая считывает целые числа с консоли по одному числу в строке.

Для каждого введённого числа проверить:

если число меньше 10, то пропускаем это число;
если число больше 100, то прекращаем считывать числа;
в остальных случаях вывести это число обратно на консоль в отдельной строке.
*/
func function18() {
	var n int

	for fmt.Scan(&n); true; fmt.Scan(&n) {
		if n < 10 {
			continue
		} else if n > 100 {
			break
		} else {
			fmt.Println(n)
		}
	}
}

/**
Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается. Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.
*/
func function19() {
	var x, p, y, count int

	for fmt.Scan(&x, &p, &y); x <= y; x = int(float32(x) * (1 + float32(p)/100)) {
		count++
	}
	fmt.Println(count)
}

/**
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
*/
func function20() {
	var x, y string
	fmt.Scan(&x, &y)
	for pos, l := range x {
		if strings.Contains(y, string(l)) {
			if pos != len(x)-1 {
				fmt.Print(string(l) + " ")
			} else {
				fmt.Print(string(l) + "\n")
			}
		}

	}
}

/**
На вход подается число типа float64. Вам нужно вывести преобразованное число по правилу: число возводится в квадрат, затем обрезается дробная часть так что остается 4 знака после запятой. Но если число равно или меньше нуля - выводить:
"число R не подходит", где R - исходное число с 2 цифрами после запятой и с 2 по ширине. А если число больше чем 10 000 - выводить исходное число в экспоненциальном представлении (знак экспоненты в нижнем регистре).
*/
func function21() {
	var x float64

	if fmt.Scan(&x); x <= 0 {
		fmt.Printf("число %2.2f не подходит\n", x)
	} else if x > 10000 {
		fmt.Printf("%e\n", x)
	} else {
		fmt.Printf("%.4f\n", x*x)
	}
}

/**
Напишите программу, принимающая на вход число N (N \geq 4)N(N≥4), а затем NN целых чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
*/
func function22() {
	var a int

	fmt.Scan(&a)
	t := make([]int, a, a)

	for i := 0; i < a; i++ {

		fmt.Scan(&t[i])

	}
	fmt.Println(t[3])
}

/**
Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0. Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).
*/
func function23() {
	var a int
	fmt.Scan(&a)
	t := make([]int, a)
	for i := range t {
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
func function24() {
	var a, count int
	fmt.Scan(&a)
	t := make([]int, a)
	for i := range t {

		if fmt.Scan(&t[i]); t[i] > 0 {
			count++
		}
	}

	fmt.Println(count)
}

/**
Дано трехзначное число. Найдите сумму его цифр.
*/
func function25() {
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
func function26() {
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

то h=3 и m=40.
*/
func function27() {
	var a int32
	fmt.Scan(&a)
	fmt.Println("It is", a/3600, "hours", a%3600/60, "minutes.")
}

/**
Заданы три числа - a, b, c (a < b < c)a,b,c(a<b<c) - длины сторон треугольника. Нужно проверить, является ли треугольник прямоугольным. Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"
*/
func function28() {
	var a, b, c int

	if fmt.Scan(&a, &b, &c); c*c == a*a+b*b {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}

/**
Даны три натуральных числа a, b, c. Определите, существует ли треугольник с такими сторонами.
*/
func functinon29() {
	var a, b, c int

	if fmt.Scan(&a, &b, &c); a < b+c && b < a+c && c < a+b {
		fmt.Println("Существует")

	} else {
		fmt.Println("Не существует")
	}
}

/**
Даны два числа. Найти их среднее арифметическое.
*/
func function30() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(float32(a+b) / 2)
}

/**
По данным числам, определите количество чисел, которые равны нулю.
*/
func function31() {
	var n, count int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var t int

		if fmt.Scan(&t); t == 0 {
			count++
		}

	}
	fmt.Println(count)
}

/**
Найдите количество минимальных элементов в последовательности.
*/
func function32() {
	var n, count, min int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var t int
		fmt.Scan(&t)
		if i == 0 {
			min = t
		}
		if min == t {
			count++
		} else if min > t {
			min = t
			count = 1
		}
	}
	fmt.Println(count)
}

/**
Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр, на каждой итерации которого для
подсчета суммы цифр берут результат, полученный на предыдущей итерации. Этот процесс повторяется до тех пор, пока не будет получена одна цифра.

Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 .
По данному числу определите его цифровой корень.
*/
func function33() {

	var a, result int
	fmt.Scan(&a)
	for {
		result = 0
		for _, r := range []rune(strconv.Itoa(a)) {
			parseResult, err := strconv.Atoi(string(r))
			if err != nil {
				fmt.Println(err)
			}
			result += parseResult
		}
		a = result
		if result < 10 {
			break
		}
	}
	fmt.Println(a)
}

/**
Найдите самое большее число на отрезке от a до b, кратное 7 .
*/
func function34() {
	var a, b int

	if fmt.Scan(&a, &b); b/7 == 0 {
		fmt.Println("NO")
	} else {
		if t := b - (b % 7); t > a {
			fmt.Println(t)
		} else {
			fmt.Println("NO")
		}
	}
}

/**
По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".
Входные данные
Дано число n (0<n<100).
Выходные данные
Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.
*/
func function35() {

}

/**
По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.
*/
func function36() {
	var (
		number int
		result float64 = math.Exp2(float64(0))
	)
	fmt.Scan(&number)
	for i := 1; int(result) <= number; i, result = i+1, math.Exp2(float64(i)) {
		fmt.Printf("%d ", int(result))
	}
}

/**
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.
*/
func function37() {

}

/**
Дано натуральное число N. Выведите его представление в двоичном виде.
*/
func function38() {
	var a int
	fmt.Scan(&a)
	fmt.Printf("%b", a)
}

/**
Из натурального числа удалить заданную цифру.
*/
func function39() {
	var a, repl int
	fmt.Scan(&a, &repl)
	result, _ := strconv.Atoi(strings.ReplaceAll(strconv.Itoa(a), strconv.Itoa(repl), ""))
	fmt.Println(result)
}

func main() {

}
