package main

import (
	"fmt"
	"math"
	"strconv"
)

type SpeSkillTest struct{}

func (s *SpeSkillTest) IsNarcissistic(num int) bool {
	strNum := strconv.Itoa(num)
	var sum float64
	for _, digit := range strNum {
		d, _ := strconv.Atoi(string(digit))
		sum += s.pow(float64(d), float64(len(strNum)))
	}
	return int(sum) == num
}

func (s *SpeSkillTest) pow(x, y float64) float64 {
	return math.Pow(x, y)
}

func (s *SpeSkillTest) FindOutlier(arr []int) interface{} {
	evenCount := 0
	oddCount := 0
	evenNum := 0
	oddNum := 0

	for _, num := range arr {
		if num%2 == 0 {
			evenCount++
			evenNum = num
		} else {
			oddCount++
			oddNum = num
		}
		if evenCount > 1 && oddCount == 1 {
			return oddNum
		} else if oddCount > 1 && evenCount == 1 {
			return evenNum
		}
	}
	return false // no outlier found
}

func (s *SpeSkillTest) FindNeedle(haystack []string, needle string) int {
	for i, str := range haystack {
		if str == needle {
			return i
		}
	}
	return -1 // needle not found
}

func (s *SpeSkillTest) BlueOcean(list1 []int, list2 []int) []int {
	result := make([]int, 0)
	for _, num := range list1 {
		if !contains(list2, num) {
			result = append(result, num)
		}
	}
	return result
}

func contains(list []int, num int) bool {
	for _, n := range list {
		if n == num {
			return true
		}
	}
	return false
}

func main() {
	test := SpeSkillTest{}

	// NARCISSISTIC NUMBER
	fmt.Println(test.IsNarcissistic(153)) // true
	fmt.Println(test.IsNarcissistic(111)) // false

	// PARITY OUTLIER
	fmt.Println(test.FindOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36})) // 11
	fmt.Println(test.FindOutlier([]int{160, 3, 1719, 19, 11, 13, -21})) // 160
	fmt.Println(test.FindOutlier([]int{11, 13, 15, 19, 9, 13, -21}))    // false

	//NEEDLE IN THE HAYSTACK
	fmt.Println(test.FindNeedle([]string{"red", "blue", "yellow", "black", "grey"}, "blue")) // 1

	//THE BLUE OCEAN REVERSE
	fmt.Println(test.BlueOcean([]int{1, 2, 3, 4, 6, 10}, []int{1})) // [2, 3, 4, 6, 10]
	fmt.Println(test.BlueOcean([]int{1, 5, 5, 5, 5, 3}, []int{5}))  // [1, 3]
}
