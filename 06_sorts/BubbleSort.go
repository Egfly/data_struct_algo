package main

import "fmt"

//冒泡排序
func BubbleSort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		isChange := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				isChange = true
			}
		}
		if !isChange {
			break
		}
	}
}

func main()  {
	a := []int{6,5,4,3,2,1}
	BubbleSort(a, 6)
	fmt.Println(a)
}
