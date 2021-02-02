package Array

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) //字典声明，[key]value
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok { //_是当前key对应的value，ok是字典是否有这个key
			return []int{m[another], i} //数组声明，[len]{values}
		}
		m[nums[i]] = i
	}
	return nil
}

//func main() {
//	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
//	if age := 20; age > 18 {
//		fmt.Println("已经成年了")
//	}
//}
