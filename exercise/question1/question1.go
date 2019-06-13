package question1

import "fmt"

// Question1 solution O(n^2)
func Question1() {
	testSet := []int{1, 2, 3, 1, 3, 6, 6}
	setAns := make(map[int]int)
	for index, num := range testSet {
		if setAns[num] == 0 {
			setAns[num] = 1
		}
		for index2, num2 := range testSet {
			if num == num2 && index2 > index {
				count := setAns[num]
				setAns[num] = count + 1
			}
		}
	}

	for key, value := range setAns {
		fmt.Println(key, " has ", value)
	}

}
