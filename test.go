package main

import (
	"fmt"
	"sort"
)

// 给定一个数n，如23121，给定一组数字a,如｛2，4，9｝，求a中元素组成的小于n的最大数，如小于23121的最大数为22999
func main() {
	var num, m int
	// 表示有 n 个节点，m 条关系描述
	fmt.Scan(&num, &m)
	// realyNum := num
	arrM := make([]int, 0)
	for m > 0 {
		arrM = append(arrM, m%10)
		m = m / 10
	}
	sort.Slice(arrM, func(i, j int) bool {
		return arrM[i] < arrM[j]
	})
	fmt.Println(arrM)
	arrN := make([]int, 0)
	for num > 0 {
		arrN = append(arrN, num%10)
		num = num / 10
	}
	fmt.Println(arrN)
	i := 0
	j := len(arrN) - 1
	for i < j {
		arrN[i], arrN[j] = arrN[j], arrN[i]
		i++
		j--
	}

	fmt.Println("arr:", arrN)
	lastSmall := false
	res := 0
	for i := 0; i < len(arrN); i++ {
		if !lastSmall {
			j := 0
			for ; j < len(arrM); j++ {
				if arrN[i] > arrM[j] {
					lastSmall = true
					break
				} else if arrM[j] == arrN[i] {
					break
				}
			}
			if lastSmall {
				res = res*10 + arrM[j]
			} else {
				res = res*10 + arrN[i]
			}
		} else {
			res = res*10 + arrM[len(arrM)-1]
		}
	}
	fmt.Println(res)
}
