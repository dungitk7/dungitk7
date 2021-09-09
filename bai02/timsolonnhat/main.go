package main

import "fmt"

func main() {
	array := []int{2, 1, 3, 4}

	fmt.Println("Tìm số lớn nhất trong mảng?", array)
	fmt.Print("Số lớn nhất là: ", findmaxsecondnum(array))
}

func findmaxsecondnum(e []int) (maxnumsecond int) {

	var maxnum int
	maxnum = 0
	for i := 0; i < len(e); i++ {
		if e[i] > maxnum {
			maxnum = e[i]
		}
	}

	for i := 0; i < len(e); i++ {
		if e[i] > maxnumsecond && e[i] < maxnum {
			maxnumsecond = e[i]
		}
	}
	return
}
