package main

import "os"

func ReflectionLineOnEachPattern() int {
	b, _ := os.ReadFile("input.txt")
	input := string(b)
	grids := parse(input)
	res := 0
	for i := range grids {
		j := verticalReflection(grids[i], false, -1)
		j = verticalReflection(grids[i], true, j)

		if j != -1 {
			res += j + 1
		} else {
			k := horizontalReflection(grids[i], false, -1)
			k = horizontalReflection(grids[i], true, k)
			if k != -1 {
				res += (k + 1) * 100
			}
		}
	}
	return res
}
