### Function to parse the input file into two lists
def parseInput():
    file_path = "D:/Projects/advent-of-code/2024/Python/Day1/input.txt"
    data = []
    with open(file_path, encoding="UTF-8") as file:
        for line in file:
            for n in line.split():
                data.append(int(n))

    left = []
    right = []

    # Distribute elements into left and right lists
    for index, value in enumerate(data):
        if index % 2 == 0:  # Even indices go to 'left'
            left.append(value)
        else:  # Odd indices go to 'right'
            right.append(value)

    return left, right




def totalDistance():
    # Sort both lists
    left, right = parseInput()
    left.sort()
    right.sort()

    # Calculate the sum of absolute differences
    total_difference = 0
    for l_val, r_val in zip(left, right):
        total_difference += abs(l_val - r_val)

    return total_difference
