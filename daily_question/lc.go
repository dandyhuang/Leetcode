package daily_question

import (
	"math"
	"sort"
)

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

// 输入：employees = [[1,5,[2,3]],[2,3,[]],[3,3,[]]], id = 1
// 输出：11
// 解释：
// 员工 1 自身的重要度是 5 ，他有两个直系下属 2 和 3 ，而且 2 和 3 的重要度均为 3 。
// 因此员工 1 的总重要度是 5 + 3 + 3 = 11 。

func getImportance(employees []*Employee, id int) int {
	mp := map[int]*Employee{}
	for _, employee := range employees {
		mp[employee.Id] = employee
	}
	total := 0
	var dfs func(int)
	dfs = func(id int) {
		e := mp[id]
		total += e.Importance
		for _, subId := range e.Subordinates {
			dfs(subId)
		}
	}
	dfs(id)
	return total
}

// 3144. 分割字符频率相等的最少子字符串
const inf = 0x3f3f3f3f

// 直接判断子串是否为平衡字符串复杂度较高，并且有大量的重复计算。因此考虑每次枚举 i 之后，再从 i 开始倒序枚举 j，
// 过程中维护一个哈希表 occ_cnt，用于存储每种字符出现的次数。另外，为了快速判断所有字符出现的次数是否相等，
// 我们需要维护所有字符出现次数的最大值 max_cnt，当满足如下条件时，所有字符出现的次数相等：
// max_cnt×len(occ_cnt)=i−j+1
func minimumSubstringsInPartition(s string) int {
	n := len(s)
	d := make([]int, n+1)
	for i := range d {
		d[i] = inf
	}
	d[0] = 0

	for i := 1; i <= n; i++ {
		maxCnt := 0
		occCnt := make(map[byte]int)
		for j := i; j >= 1; j-- {
			occCnt[s[j-1]]++
			if occCnt[s[j-1]] > maxCnt {
				maxCnt = occCnt[s[j-1]]
			}
			if maxCnt*len(occCnt) == (i-j+1) && d[j-1] != inf {
				if d[i] > d[j-1]+1 {
					d[i] = d[j-1] + 1
				}
			}
		}
	}
	return d[n]
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func minimumSubstringsInPartitionV2(s string) int {
	dp := make([]int, len(s)+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	// 设置 d[i] 为将以 i 结尾的前缀字符串划分平衡字符串的最少个数
	dp[0] = 0
	for i := 1; i <= len(s); i++ {
		mCount := make(map[uint8]int, 0)
		maxCount := 0
		for j := i; j >= 1; j-- {
			mCount[s[j-1]]++
			if mCount[s[j-1]] > maxCount {
				maxCount = mCount[s[j-1]]
			}
			if maxCount*len(mCount) == i-j+1 && dp[j-1] != math.MaxInt {
				dp[i] = min(dp[i], dp[j-1]+1)
			}
		}
	}
	return dp[len(s)]
}

// 你有一个数组 nums ，它只包含 正 整数，所有正整数的数位长度都 相同 。
// 两个整数的 数位差 指的是两个整数 相同 位置上不同数字的数目。
// 请你返回 nums 中 所有 整数对里，数位差之和。
// 输入：nums = [13,23,12]
// 输出：4
// - 13 和 23 的数位差为 1 。
// - 13 和 12 的数位差为 1 。
// - 23 和 12 的数位差为 2 。
// 所以所有整数数对的数位差之和为 1 + 1 + 2 = 4 。
func sumDigitDifferences(nums []int) int64 {
	var res int64
	for i := 0; i < len(nums)-1; i++ {
		now := nums[i]
		next := nums[i+1]
		for now != 0 {
			res += int64(math.Abs(float64(now%10) - float64(next%10)))
			now = now / 10
			next = next / 10
		}
	}
	return res
}

// 2708. 一个小组的最大实力值
// 输入：nums = [3,-1,-5,2,5,-9]
// 输出：1350
// 解释：一种构成最大实力值小组的方案是选择下标为 [0,2,3,4,5] 的学生。实力值为 3 * (-5) * 2 * 5 * (-9) = 1350 ，
// 这是可以得到的最大实力值。
func maxStrength(nums []int) int64 {
	var res int64
	res = 1
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	if len(nums) == 1 {
		return int64(nums[0])
	}

	negativeNum := 0
	positiveNum := 0
	zeroNum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			positiveNum++
			break
		} else if nums[i] < 0 {
			negativeNum++
		} else {
			zeroNum++
		}
	}
	if negativeNum%2 == 0 {
		for i := 0; i < len(nums); i++ {
			if nums[i] != 0 {
				res *= int64(nums[i])
			}
		}
	} else {
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				continue
			}
			if i < negativeNum-1 && nums[i] < 0 {
				res *= int64(nums[i])
			} else if nums[i] > 0 {
				res *= int64(nums[i])
			}
		}
	}
	if negativeNum == 1 && zeroNum == 0 && positiveNum == 0 {
		return int64(nums[0])
	}
	if negativeNum <= 1 && positiveNum == 0 {
		return int64(0)
	}

	return res
}
