package main

import "fmt"

func main() {
	BubbleShort()

}

func BubbleShort() []int {

	arr := []int{34, 15, 29, 8}

	//Short(arr)

	for i := 0; i < len(arr); i++ {
		//fmt.Println(arr[i])

		for j := 0; j < len(arr)-1; j++ {

			if arr[j] > arr[j+1] {
				arr = Swap(arr, j, j+1)

			}

		}

	}
	fmt.Println(arr)
	return arr
}

func Swap(arr []int, a, b int) []int {

	temp := arr[a]

	arr[a] = arr[b]

	arr[b] = temp

	return arr

}
