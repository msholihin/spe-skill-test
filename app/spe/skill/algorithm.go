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

func main() {
	test := SpeSkillTest{}

	// NARCISSISTIC NUMBER
	fmt.Println(test.IsNarcissistic(153)) // true
	fmt.Println(test.IsNarcissistic(111)) // false

}
