package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*var (
		err error
		a   float64
		b   float64
		c   float64
	)

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter number a :")
		str, _ := reader.ReadString('\n')
		str = strings.TrimSuffix(str, "\n")
		fmt.Print(str)
		if a, err = strconv.ParseFloat(str, 64); err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Print("Enter number b :")
		str, _ = reader.ReadString('\n')
		str = strings.TrimSuffix(str, "\n")
		fmt.Print(str)
		if b, err = strconv.ParseFloat(str, 64); err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Print("Enter number c :")
		str, _ = reader.ReadString('\n')
		str = strings.TrimSuffix(str, "\n")
		fmt.Print(str)
		if c, err = strconv.ParseFloat(str, 64); err != nil {
			fmt.Println(err.Error())
			return
		}
	*/
	a, err := readNumberFromKeyboard("Giá trị của a: ")
	checkErr(err)

	b, err := readNumberFromKeyboard("Giá trị của b: ")
	checkErr(err)

	c, err := readNumberFromKeyboard("Giá trị của c: ")
	checkErr(err)

	_, _ = quadratics(a, b, c)

}

func readNumberFromKeyboard(msg string) (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	str, _ := reader.ReadString('\n')
	str = strings.Trim(str, "\n")
	number, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("[Error]: %s", err)
		os.Exit(1)
	}
}

func quadratics(a, b, c float64) (float64, float64) {
	delta := (b * b) - (4 * a * c)

	switch {
	case delta > 0:
		x1 := (-b + math.Sqrt(delta)) / 2 * a
		x2 := (-b - math.Sqrt(delta)) / 2 * a
		//fmt.Println("x1 = ", x1, "x2 = ", x2) // Sau khi khai báo x1,x2 -> Cần in ra giá trị
		fmt.Printf("Phương trình có 2 nghiệm: x1 = %v, x2 = %v", x1, x2)
	case delta == 0:
		x1 := (-b / 2 * a)
		x2 := (-b / 2 * a)
		fmt.Printf("Phương trình có cùng nghiệm: x1 = %v, x2 = %v", x1, x2)
	case delta < 0:
		x1 := complex(-b/(2*a), math.Sqrt(-delta)/(2*a))
		x2 := complex(-b/(2*a), -math.Sqrt(-delta)/(2*a))
		fmt.Printf("Phương trình có 2 nghiệm ở dạng số phức: x1 = %v, x2 = %v", x1, x2)
	}
	return 0, 0
}
