package main

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