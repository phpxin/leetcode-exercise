/*
 @Title 寻找数组的中心索引
 @Description https://leetcode-cn.com/leetbook/read/array-and-string/yf47s/
 @Author  Leo
 @Update  2020/10/20 12:23 下午
*/

package main

import (
	"fmt"
)

func main() {
	examples := make([][]int, 0)
	examples = append(examples, []int{1, 7, 3, 6, 5, 6})
	examples = append(examples, []int{-1, 3, 2, 1, 1, 0})
	examples = append(examples, []int{-1, -1, -1, -1, -1, 0})
	examples = append(examples, []int{-1, -1, -1, 0, 1, 1})
	examples = append(examples, []int{1, 2, 3})

	for _, nums := range examples {
		r := pivotIndex(nums)
		fmt.Println("result of ", nums, " is ", r)
	}
}

// 核心思路：前缀和算法 https://www.jianshu.com/p/d0dabea38302
func pivotIndex(nums []int) int {
	total := len(nums)
	if total < 3 {
		return -1
	}

	leftSum := 0  // 左侧和
	rightSum := 0 // 右侧和

	// 首先把数据分为三部分，左侧和为 0 ，中间索引为 0
	// 从索引 1 开始累加全部单元，计算出右侧和
	for i := 1; i < total; i++ {
		rightSum = rightSum + nums[i]
	}

	for i := 0; i < total; i++ {
		if leftSum == rightSum {
			return i
		}
		// 左侧和不断累加，右侧和不断减少
		// 直到左右相等，则 i 为结果
		leftSum = leftSum + nums[i]
		if i+1 == total {
			rightSum = 0
		} else {
			rightSum = rightSum - nums[i+1]
		}
	}

	// 无结果，返回 -1 未找到
	return -1
}
