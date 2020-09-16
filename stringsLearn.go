package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/**
На вход подается строка. Нужно определить, является ли она правильной или нет. Правильная строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная - вывести Right иначе - вывести Wrong
*/
func StringTest() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	test := []rune(text)
	if unicode.IsUpper(test[0]) && strings.HasSuffix(text, ".\n") {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
func main() {

	StringTest()
}
