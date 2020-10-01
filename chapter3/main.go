package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

func main() {

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
