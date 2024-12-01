### Function to parse the input file into two lists
def parseInput():
    file_path = "D:/Projects/advent-of-code/2024/Go/Day1/input.txt"
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


# Function to calculate the weighted sum based on counts (Part 2)
def similarityScore():
    # Sort both lists (even if already sorted in Part 1, explicitly do it here)
    left, right = parseInput()
    left.sort()
    right.sort()

    # Calculate the weighted sum
    weighted_sum = 0
    for l_val in left:
        count = right.count(l_val)  # Count occurrences of `l_val` in `right`
        weighted_sum += l_val * count

    return weighted_sum
