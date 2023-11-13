package leetcode_hot100

import "fmt"

// 组合排序
// 77. 组合
// 输入：n = 4, k = 2
// 输出：
// [[2,4], [3,4], [2,3],[1,2], [1,3], [1,4]]
func combine(n int, k int) [][]int {
	var res [][]int
	var arr []int
	var dfs func(int, int, int, []int)
	dfs = func(n, k, start int, arr []int) {
		if len(arr)+n-start+1 < k {
			return
		}
		if len(arr) == k {
			tmp := make([]int, k)
			copy(tmp, arr)
			res = append(res, tmp)
			// 直接赋值有问题  拷贝的都是最后一个值
			// res = append(res, arr)
			return
		}
		for i := start; i <= n; i++ {
			arr = append(arr, i)
			dfs(n, k, i+1, arr)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(n, k, 1, arr)
	return res
}

// 46. 全排列
func permute(nums []int) [][]int {
	var res [][]int
	var arr []int
	used := make([]int, len(nums))
	var dfs func([]int, []int, []int)
	dfs = func(nums, used []int, arr []int) {
		if len(nums) == len(arr) {
			tmp := make([]int, len(nums))
			copy(tmp, arr)
			res = append(res, tmp)
			return
		}
		// i始终从0开始递归
		for i := 0; i < len(nums); i++ {
			if used[i] == 1 {
				continue
			}
			used[i] = 1
			arr = append(arr, nums[i])
			dfs(nums, used, arr)
			arr = arr[:len(arr)-1]
			used[i] = 0
		}
	}
	dfs(nums, used, arr)
	return res
}

// 78. 子集
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
func subsets(nums []int) [][]int {
	var res [][]int
	var arr []int
	var dfs func(int, []int)
	dfs = func(start int, arr []int) {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		res = append(res, tmp)
		if start >= len(nums) {
			return
		}
		for i := start; i < len(nums); i++ {
			arr = append(arr, nums[i])
			fmt.Println("in:", i, arr)
			dfs(i+1, arr)
			arr = arr[:len(arr)-1]
			fmt.Println("out:", i, arr)
		}
	}
	dfs(0, arr)
	return res
}
