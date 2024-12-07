package main

// Part2 solves part 2 using Add, Multiply, and Concatenate operators.
func elephantHidingTotalCalibration() (int, error) {
	calibrationEquations, err := ParseInputFile()
	if err != nil {
		return 0, err
	}

	total := 0
	for _, eq := range calibrationEquations {
		result := eq[0].(int)
		terms := eq[1].([]int)
		if ValidateEquation(result, terms, nil, []func(*int, int) int{Add, Multiply, Concatenate}) {
			total += result
		}
	}
	return total, nil
}
