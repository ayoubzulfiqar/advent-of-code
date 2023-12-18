package main

func traverseAll(objectMap map[Location]*Object) int {
	maxEnergy := 0
	for x := 0; x <= LastCol; x++ {
		traverse(objectMap, Location{x, 0}, TOP)
		maxEnergy = max(maxEnergy, countEnergized(objectMap))
		traverse(objectMap, Location{x, LastRow}, BOTTOM)
		maxEnergy = max(maxEnergy, countEnergized(objectMap))
	}
	for y := 0; y <= LastRow; y++ {
		traverse(objectMap, Location{0, y}, LEFT)
		maxEnergy = max(maxEnergy, countEnergized(objectMap))
		traverse(objectMap, Location{LastCol, y}, RIGHT)
		maxEnergy = max(maxEnergy, countEnergized(objectMap))
	}
	return maxEnergy
}

func Run2() int {
	answer := 0
	objectMap := parseFile("input.txt")
	answer = traverseAll(objectMap)
	// 7716
	println(answer)
	return answer
}
