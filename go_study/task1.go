package main

import (
	"fmt"
	"slices"
	"strconv"
)

// 1、只出现一次的数字
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
//结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。

func getSingleNumber(arr []int) int {
	arrMap := make(map[int]int)
	for _, v := range arr {
		arrMap[v]++
	}

	// 遍历map,找到出现次数为1的元素
	for k, _ := range arrMap {
		if arrMap[k] == 1 {
			return k
		}
	}

	return -1

}

// 2、回文数
func isPalindrome(n int) bool {
	// 去掉无效数字
	if n < 0 || (n != 0 && n%10 == 0) {
		return false
	}
	// 将数字转换为字符串
	numStr := strconv.Itoa(n)
	strLen := len(numStr)
	for i := 0; i < strLen/2; i++ {
		if numStr[i] != numStr[strLen-i-1] {
			return false
		}
	}

	return true
}

// 3、给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func isValid(str string) bool {
	// 去掉空和奇数长度的字符串
	if len(str) == 0 || len(str)%2 == 1 {
		return false
	}
	// 用map记录左右括号的对应关系,byte为uint8类型
	strMap := make(map[byte]byte)
	strMap['('] = ')'
	strMap['{'] = '}'
	strMap['['] = ']'

	// 用切片来代替栈，记录左括号
	stack := make([]byte, 0)

	// 遍历字符串
	for i := 0; i < len(str); i++ {
		// 如果是左括号，压入栈
		if key, ok := strMap[str[i]]; ok {
			stack = append(stack, key)
		} else {
			// 如果是右括号，判断栈顶是否匹配
			if len(stack) == 0 || str[i] != stack[len(stack)-1] {
				return false
			}
			// 如果匹配，弹出栈顶
			stack = stack[:len(stack)-1] //Go中没有pop方法，只能用切片操作
		}

	}
	return len(stack) == 0
}

// 4、查找字符串数组中的最长公共前缀
func longCommonStrPrefix(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	// 遍历数组，将第一个元素与后续的元素进行比较，找到最长公共前缀
	for i := 0; i < len(arr[0]); i++ {
		for j := 1; j < len(arr); j++ {
			if len(arr[j]) < i || arr[j][i] != arr[0][i] {
				// 返回最长公共前缀
				return arr[0][:i]
			}
		}
	}

	return arr[0]
}

// 5、给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func plusOne(arr []int) []int {
	// 遍历数组，找到最后一个元素
	for i := len(arr) - 1; i >= 0; i-- {
		// 如果最后一个元素不是9，则加1并返回
		if arr[i] != 9 {
			arr[i]++
			return arr
		}
		// 如果最后一个元素是9，则置为0
		arr[i] = 0
	}

	// 如果所有的元素都是9，则需要新创建一个比原数组大1的数组，并在最前面添加1
	arr = make([]int, len(arr)+1)
	arr[0] = 1
	return arr
}

// 6、给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 默认第一个元素有效，是不重复的
	slow := 1
	//遍历数组，如果当前元素与前一个元素不同，则将当前元素赋值给slow，并将slow指针后移
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

// 7、合并区间
func mergerInterval(intervals [][]int) (ans [][]int) {
	// 将二维数组按照左端点从小到大排序
	slices.SortFunc(intervals, func(left, right []int) int { return left[0] - right[0] })
	// 合并区间
	for _, interval := range intervals {
		m := len(ans)
		if m > 0 && interval[0] < ans[m-1][1] {
			// 更新右端点
			ans[m-1][1] = max(interval[1], ans[m-1][1])
		} else {
			ans = append(ans, interval)
		}

	}

	return
}

// 8、给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回下标
func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j := i + 1; j <= target; j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 检查作业
func main() {
	// 1、只出现一次的数字
	arr1 := []int{5, 2, 2, 3, 3}
	result1 := getSingleNumber(arr1)
	if result1 != -1 {
		fmt.Println("出现一次的元素为：", result1)
	} else {
		fmt.Println("数组中没有只出现一次的元素")
	}

	// 2、回文数
	n := 12321
	result2 := isPalindrome(n)
	if result2 {
		fmt.Println(n, "是回文数")
	} else {
		fmt.Println(n, "不是回文数")
	}

	// 3、给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
	str3 := "()[]{}"
	result3 := isValid(str3)
	if result3 {
		fmt.Println(str3, "是有效的")
	} else {
		fmt.Println(str3, "是无效的")
	}

	// 4、查找字符串数组中的最长公共前缀
	arr4 := []string{"test1", "tes2", "te33"}
	result4 := longCommonStrPrefix(arr4)
	if result4 == "" {
		fmt.Println("数组中没有公共前缀")
	} else {
		fmt.Println("数组中最长公共前缀为：", result4)
	}

	// 5、给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
	arr5 := []int{9, 9, 8}
	result5 := plusOne(arr5)
	fmt.Println(result5)

	// 6、给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度
	nums6 := []int{1, 1, 2, 2, 3, 4}
	result6 := removeDuplicates(nums6)
	fmt.Println("去重后的数组长度为：", result6)

	// 7、合并区间
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result7 := mergerInterval(intervals)
	fmt.Println(result7)

	// 8、给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回下标
	nums := []int{1, 2, 3, 4, 5}
	result := twoSum(nums, 6)
	fmt.Println(result)

}
