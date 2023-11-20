package main

import "fmt"

func main() {
	classes := [][]int8{}
	classes = append(classes, []int8{100, 40, 50, 55, 64, 71, 12, 90, 75, 25})
	classes = append(classes, []int8{68, 43, 34, 87, 89})
	classes = append(classes, []int8{23, 55, 45, 67, 45, 12, 34, 54, 35, 24, 96, 87, 100, 45, 65})

	for i, class := range classes {
		for _, mark := range class {
			if mark == 100 {
				fmt.Printf("Full mark with %d marks from class %d \n", mark, i)
			} else if mark >= 80 {
				fmt.Printf("Flying color with %d marks from class %d \n", mark, i)
			} else if mark >= 40 {
				fmt.Printf("Pass with %d marks from class %d \n", mark, i)
			} else {
				fmt.Printf("Fail with %d marks from class %d \n", mark, i)
			}

		}
		fmt.Println()
	}
}
