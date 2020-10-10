package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

// Задача состоит в следующем: на стандартный ввод подаются целые числа в диапазоне 0-100, каждое число подается на стандартный ввод с новой
// строки (количество чисел не известно). Требуется прочитать все эти числа и вывести в стандартный вывод их сумму.
func function1() {
	read := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	var result int

	defer func() {
		wr.WriteString(strconv.Itoa(result) + "\n")
		wr.Flush()
	}()

	for read.Scan() {

		readD, _ := strconv.Atoi(read.Text())
		result = result + readD
	}
}

// На стандартный ввод подаются данные о студентах университетской группы в формате JSON:
// В сведениях о каждом студенте содержится информация о полученных им оценках (Rating).
// Требуется прочитать данные, и рассчитать среднее количество оценок, полученное студентами группы.
// Ответ на задачу требуется записать на стандартный вывод в формате JSON в следующей форме:
type Student struct {
	LastName, FirstName, MiddleName, Birthday, Address, Phone string
	Rating                                                    []int
}
type Request struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

type Response struct {
	Average float32
}

func (r *Response) getAverageFromStudent(students []Student) {
	arr := make([]int, 0)
	var result int
	for _, stud := range students {

		arr = append(arr, len(stud.Rating))
	}

	for _, v := range arr {
		result = result + v
	}

	r.Average = float32(result) / float32(len(arr))
}

func function2() {

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		// ...
	}
	req := new(Request)

	err = json.Unmarshal(data, &req)
	if err != nil {
		fmt.Println(err)
	}

	resp := new(Response)
	resp.getAverageFromStudent(req.Students)

	respB, _ := json.MarshalIndent(resp, "", "    ")
	fmt.Println(string(respB))
}

// На стандартный ввод подается строковое представление даты и времени в следующем формате:
// 1986-04-16T05:20:00+06:00
// Ваша задача конвертировать эту строку в Time, а затем вывести в формате UnixDate:
// Wed Apr 16 05:20:00 +0600 1986

func function3(line string) {
	now, _ := time.Parse(time.RFC3339, line)
	fmt.Println(now.Format(time.RubyDate))

}

// На стандартный ввод подается строковое представление даты и времени определенного события в следующем формате:
// 2020-05-15 08:00:00
// Если время события до обеда (13-00), то ничего менять не нужно, достаточно вывести дату на стандартный вывод в том же формате.
// Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день, а затем вывести на стандартный вывод в том же формате.
func function4(line string) {
	timeFormat := "2006-01-02 15:04:05"
	date, err := time.Parse(timeFormat, line)
	if err != nil {
		fmt.Println(err)
	}

	if date.Hour() > 13 {

		date = date.Add(time.Hour * 24)
	}
	fmt.Println(date.Format(timeFormat))

}

/*
На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).
Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.
Sample Input:
13.03.2018 14:00:15,12.03.2018 14:00:15
*/
func function5(lineTimes string) {
	timeFormat := "02.01.2006 15:04:05"
	times := strings.Split(lineTimes, ",")
	time1, err1 := time.Parse(timeFormat, times[0])
	time2, err2 := time.Parse(timeFormat, times[1])
	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
	}
	var dur time.Duration

	if time1.After(time2) {
		dur = time1.Sub(time2)

	} else {
		dur = time2.Sub(time1)

	}
	fmt.Println(dur)

}

// Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал.
// Функция должна называться task().
func task(c chan int, N int) {
	c <- N + 1
}

// Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.
// Функция должна называться task2().
func task2(c chan string, s string) {
	for i := 0; i < 5; i++ {
		c <- s + " "
	}
}
func squares(c chan int) {
	for i := 0; i <= 100; i++ {
		c <- i
		fmt.Println("<Squares", i)
	}
	close(c)
}

func task3() {
	fmt.Println("main() started")

	c := make(chan int, 10)

	go squares(c)

	for val := range c {
		fmt.Println("> Main", val)

	}

	fmt.Println("main() stoped")
}

// Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет
// значения на следующий этап конвейера только если оно отличается от того, что пришло ранее.
// Ваша фукнция должна принимать два канала - inputStream и outputStream, в первый вы будете
// получать строки, во второй вы должны отправлять значения без повторов. В итоге в outputStream
// должны остаться значения, которые не повторяются подряд. Не забудьте закрыть канал
func removeDuplicates(input chan string, output chan string) {
	defer close(output)
	m := make(map[string]bool)
	for v := range input {
		if _, exist := m[v]; exist != true {
			output <- v
			m[v] = true

		}
	}
}

func main() {
	list := [...]string{"test", "relax", "test", "ping", "pong"}
	input := make(chan string)
	output := make(chan string)
	go func(input chan string) {
		defer close(input)

		for _, v := range list {

			input <- v
		}

	}(input)

	go removeDuplicates(input, output)

	for r := range output {
		fmt.Println(">>>", r)
	}

}
