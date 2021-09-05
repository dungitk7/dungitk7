package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func readNumberFromKeyboard(msg string) (int, error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	str, _ := reader.ReadString('\n')
	str = strings.Trim(str, "\n")
	number, err := strconv.Atoi(str)

	rand.Seed(time.Now().UnixNano())
	numrand := rand.Intn(100)

	switch {
	case number < numrand:
		fmt.Printf("Số bạn nhập nhỏ hơn : %v", numrand)
	case number > numrand:
		fmt.Printf("Số bạn nhập lớn hơn : %v", numrand)
	case number == numrand:
		fmt.Printf("Bạn đã đoán đúng : %v", numrand)
	}

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

func main() {
	r, err := readNumberFromKeyboard("Giá trị của số đã nhập: ")
	checkErr(err)
	fmt.Printf("Number: %v", r)

}
