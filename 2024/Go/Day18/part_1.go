package main

/*


type Complex struct {
	Real, Imag int
}

func (c Complex) Add(other Complex) Complex {
	return Complex{c.Real + other.Real, c.Imag + other.Imag}
}

func parseInput(filename string) ([]Complex, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var bytes []Complex
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}

		x, err1 := strconv.Atoi(parts[0])
		y, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid input format")
		}

		bytes = append(bytes, Complex{Real: x, Imag: y})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return bytes, nil
}


*/

func minimumStepsToExit(bytes []Complex) int {
	start := Complex{Real: 0, Imag: 0}
	goal := Complex{Real: 70, Imag: 70}
	steps := 0

	walls := make(map[Complex]bool)
	for _, wall := range bytes[:1024] {
		walls[wall] = true
	}

	front := map[Complex]bool{start: true}
	seen := map[Complex]bool{start: true}

	directions := []Complex{
		{Real: 1, Imag: 0},
		{Real: -1, Imag: 0},
		{Real: 0, Imag: 1},
		{Real: 0, Imag: -1},
	}

	for len(front) > 0 {
		newFront := make(map[Complex]bool)
		steps++

		for pos := range front {
			for _, d := range directions {
				newPos := pos.Add(d)

				if newPos == goal {
					return steps
				}

				if newPos.Real < 0 || newPos.Real > 70 || newPos.Imag < 0 || newPos.Imag > 70 {
					continue
				}

				if walls[newPos] || seen[newPos] {
					continue
				}

				seen[newPos] = true
				newFront[newPos] = true
			}
		}

		front = newFront
	}
	return 0
}

// func main() {
// 	bytes, err := parseInput("input.txt")
// 	if err != nil {
// 		fmt.Println("Error reading input:", err)
// 		return
// 	}

// 	result := minimumStepsToExit(bytes)
// 	fmt.Println("Minimum steps to exit:", result)
// }
