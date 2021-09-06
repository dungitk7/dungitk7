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

func readNumberFromKeyboard(msg string) (int, error) { //func đọc số từ bàn phím

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	str, _ := reader.ReadString('\n')
	str = strings.Trim(str, "\n")    // cắt bỏ \n sau string nhập vào. Ví dụ: 5\n -> 5
	number, err := strconv.Atoi(str) // convert number từ string sang int

	if err != nil { // func trả ra l
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

	rand.Seed(time.Now().UnixNano())
	numrand := rand.Intn(100)

	switch {
	case r < numrand:
		fmt.Printf("Số bạn nhập nhỏ hơn : %v", numrand)
	case r > numrand:
		fmt.Printf("Số bạn nhập lớn hơn : %v", numrand)
	case r == numrand:
		fmt.Printf("Bạn đã đoán đúng : %v", numrand)
	}
}
