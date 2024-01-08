package hot100_2

// 53. 最大子数组和
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。!
func maxSubArray(nums []int) int {

}

// 56. 合并区间
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6]. !
func merge(intervals [][]int) [][]int {

}

// 189. 轮转数组
// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]
// 解释:
// 向右轮转 1 步: [7,1,2,3,4,5,6]
// 向右轮转 2 步: [6,7,1,2,3,4,5]
// 向右轮转 3 步: [5,6,7,1,2,3,4] !
func rotate(nums []int, k int) {

}

// 238. 除自身以外数组的乘积
// 输入: nums = [1,2,3,4]
// 输出: [24,12,8,6]
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	lPlus, rPlus := make([]int, len(nums)), make([]int, len(nums))
	plus := 1
	for i := range nums {
		lPlus[i] = plus
		plus *= nums[i]
	}
	plus = 1
	for i := len(nums) - 1; i >= 0; i-- {
		rPlus[i] = plus
		plus *= nums[i]
	}
	for i := 0; i < len(nums); i++ {
		res[i] = lPlus[i] * rPlus[i]
	}
	return res
}

// 41. 缺失的第一个正数
// 输入：nums = [3,4,-1,1]
// 输出：2
func firstMissingPositive(nums []int) int {
}
