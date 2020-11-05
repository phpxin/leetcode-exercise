/*
 @Title 删除排序数组中的重复项
 @Description https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/
 @Author  Leo
 @Update  2020/9/16 5:18 下午
*/

package main

import "fmt"

func main() {
	examples := make([][]int, 0)
	examples = append(examples, []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	examples = append(examples, []int{1, 1, 2})

	for _, srcNums := range examples {
		nums := make([]int, len(srcNums))
		copy(nums, srcNums)
		c := removeDuplicates(nums)
		nums = nums[:c]
		fmt.Println(fmt.Sprintf("result of %+v is %d %+v", srcNums, c, nums))
	}
}

// B
// 核心思想：双指针 https://www.cnblogs.com/kyoner/p/11087755.html
// 快指针 j 负责扫描重复项，慢指针 i 负责重新排列成员
// 程序开始 j 不停的向前移动
// 当 j 扫描到不重复数据，则 i 向前移动并且将 j 当前位置的值赋值到 i 位置
func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

// D
func removeDuplicates1(nums []int) int {
	c := len(nums)
	if c < 2 {
		return c
	}
	i := 0
	j := i + 1
	for {
		if nums[i] == nums[j] {
			j++

			if j >= c {
				return i + 1
			}
		} else {
			if j-i > 1 {
				for m, n := i+1, j; n < c; m++ {
					nums[m] = nums[n]
					n++
				}
				c = c - (j - i - 1)
			}
			i++
			j = i + 1
			if j >= c {
				return i + 1
			}
		}
	}
}
