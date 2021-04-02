package main

import (
	"fmt"
	"strconv"
)

func sum(nos ...interface{}) int {
	result := 0
	for _, no := range nos {
		switch no.(type) {
		case int:
			result += no.(int)
		case string:
			if val, err := strconv.Atoi(no.(string)); err == nil {
				result += val
			}
		case []int:
			noList := no.([]int)
			intfList := make([]interface{}, len(noList))
			for idx, x := range noList {
				intfList[idx] = x
			}
			result += sum(intfList...)
		case []interface{}:
			result += sum(no.([]interface{})...)
		}

	}
	return result
}

/*
Use cases:
sum(10,20)
sum(10,20,30,40,50)
sum(10)
sum()
sum(10,20,30,"40") => 100
sum(10, "20", 30 , "abc") => 60

sum(10,20, []int{30,40,50})
sum(10,20, []int{30,40,"50"})
*/

func main() {
	//nos := []int{10, 20, 30, 40, 50}
	/* fmt.Println(sum(nos[0], nos[1], nos[2], nos[3], nos[4])) */
	//fmt.Println(sum(nos...))

	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30, 40, 50))
	fmt.Println(sum(10))
	fmt.Println(sum())
	fmt.Println(sum(10, 20, 30, "40"))
	fmt.Println(sum(10, "20", 30, "abc"))
	fmt.Println(sum(10, 20, []int{30, 40, 50}))
	fmt.Println(sum(10, 20, []interface{}{30, 40, "50"}))
	fmt.Println(sum(10, 20, []interface{}{30, 40, "50", []int{1, 2, 3}}))

}
