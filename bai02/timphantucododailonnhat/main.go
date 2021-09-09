package main

import "fmt"

// đo độ dài của từng phần tử => vứt ra 1 chỗ nào đó trước
// từ các số vừa được vứt ra từ task 1 => tìm ra số lớn nhất trong đó => lại lưu lại số đó ra 1 chỗ nào đó
// duyệt lại mảng ban đầu => so sánh  để tìm ra phần tử của chuỗi đó có chiều dài bằng số vừa lưu ra ở task

func main() {
	arrs := []string{"abc", "aa", "ad", "vcd"}
	fmt.Println("Phần tử có độ dài lớn nhất: ", findmaxmaxLenElement(arrs))
}

func findmaxmaxLenElement(e []string) (maxLenElement []string) {

	// Duyệt mảng và tìm ra độ dài của phần tử lớn nhất
	maxnum := 0
	for i := 0; i < len(e[i]); i++ {
		if len(e[i]) > maxnum {
			maxnum = len(e[i])
		}
	}
	// Duyệt mảng và in ra phần tử có độ dài bằng độ dài lớn nhất
	for _, v := range e {
		if len(v) == maxnum {
			maxLenElement = append(maxLenElement, v)
		}
	}
	return
}

// cách dùng 1 vòng for
// maxLength := len(e[0])

// for _, v := range e {
// 	switch {
// 	case len(v) < maxLength:
// 		continue
// 	case len(v) == maxLength:
// 		maxLenElement = append(maxLenElement, v)
// 	case len(v) > maxLength:
// 		maxLength = len(v)
// 		maxLenElement = maxLenElement[:0]
// 		maxLenElement = append(maxLenElement, v)
// 	}
// }
// return
