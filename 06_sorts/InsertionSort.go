package main

import "fmt"

//插入排序
func InsertionSort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		tmp := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > tmp {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = tmp
	}
}

func main()  {
	a := []int{6,5,4,3,2,1}
	InsertionSort(a, 6)
	fmt.Println(a)
}

