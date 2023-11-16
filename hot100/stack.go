package leetcode_hot100

// 20. 有效的括号
func isValid(s string) bool {
	stack := make([]byte, 0)
	m := map[int32]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	for i, str := range s {
		if v, ok := m[str]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == v {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
