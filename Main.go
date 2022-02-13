package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func convToFloat(o string) (output float64, err error) {
	output, err = strconv.ParseFloat(o, 64)
	return
}
func main() {
	var s string
	var k, opi, ec int
	var it, numa, numb float64
	var err error
	//vector := []int{}
	fmt.Println("Введите число")
	fmt.Scanln(&s)
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	y := func(c rune) bool {
		return unicode.IsDigit(c)
	}
	fmt.Println(strings.FieldsFunc(s, f))
	num := strings.FieldsFunc(s, f)
	op := strings.FieldsFunc(s, y)
	//fmt.Print("Массив чисел ")
	//fmt.Printf("%T %v\n", num, num)
	//fmt.Print("Массив операций ")
	//fmt.Printf("%T %v\n", op, op)
	exm := num[0]
	for i := range op {
		exm += " " + op[i] + " " + num[i+1]
	}
	for i := range op {
		if op[i] == "*" || op[i] == "/" {
			opi++
		}
	}
	for i := range op {
		if i >= opi {
			break
		}
		for j := range op {
			if op[j] == "*" || op[j] == "/" {
				i = j
				break
			}
		}
		numa, err = convToFloat(num[i])
		if err != nil {
			fmt.Println("Ошибка в числе")
			ec++
			fmt.Println(err)
		}
		numb, err = convToFloat(num[i+1])
		if err != nil {
			fmt.Println("Ошибка в числе")
			ec++
			fmt.Println(err)
		}
		switch op[i] {
		case "*":
			{
				newn := numa * numb
				num[i] = fmt.Sprintf("%g", newn)
				for j := i + 1; j <= len(num)-2; j++ {
					op[j-1] = op[j]
					num[j] = num[j+1]
				}
				k++
			}
		case "/":
			{
				if numb == 0 {
					fmt.Println("Деление на 0")
					ec++
					break
				}
				newn := numa / numb
				num[i] = fmt.Sprintf("%g", newn)
				for j := i + 1; j <= len(num)-2; j++ {
					op[j-1] = op[j]
					num[j] = num[j+1]

				}
				k++
			}
		}
	}
	op = op[:len(op)-k]
	num = num[:len(num)-k]
	it, _ = convToFloat(num[0])
	for i := range op {
		numa, err = convToFloat(num[i])
		if err != nil {
			ec++
			fmt.Println("Ошибка в числе")
			fmt.Println(err)
		}
		switch op[i] {
		case "+":
			{
				numa, _ := convToFloat(num[i+1])
				it += numa
			}
		case "-":
			{
				numa, _ := convToFloat(num[i+1])
				it -= numa
			}
		default:
			{
				if len(op[i]) > 1 {
					fmt.Println("Ошибка в записи оператора")
					ec++
					break
				}
			}
		}
	}
	if ec == 0 {
		fmt.Println(exm, " = ", it)
	}
}
