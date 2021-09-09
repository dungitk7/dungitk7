package main

import "fmt"

func main() {
	array := []int{1, 2, 5, 2, 6, 2, 5}
	fmt.Println("Xoá phần tử lặp nhau trong mảng", array)
	fmt.Println("Kết quả: ", removeDuplicate(array))

}

func removeDuplicate(e []int) (output []int) {
	keys := map[int]bool{}
	for _, entry := range e {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			output = append(output, entry)

		}

	}
	return
}
