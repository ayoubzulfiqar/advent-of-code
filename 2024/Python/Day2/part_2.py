from copy import deepcopy

def readInput(filename: str) -> list[list[int]]:
    with open(filename, "r", encoding="utf-8") as file:
        # Read lines, split them, and convert to integers
        data = [list(map(int, line.split())) for line in file if line.strip()]
    return data

def test(l, index):  # noqa: E741
    # Brute force it
    for i in range(len(l)):
        line = deepcopy(l)
        line.pop(i)

        safe = True
        if line[1] < line[0]: # Decreasing
            for j in range(len(line)):
                if j > 0:
                    if (line[j] - line[j-1]) not in [-1, -2, -3]:
                        safe = False
                        break
            if safe is True:
                return True
        elif line[1] > line[0]: # Increasing
            for j in range(len(line)):
                if j > 0:
                    if (line[j] - line[j-1]) not in [1, 2, 3]:
                        safe = False
                        break
            if safe is True:
                return True
        else: # This means that the second and first number are the same
            continue

def singleLevelSafeReports() -> int:
    totalSafe = 0
    problems = []
    problemIndexes = []
    data = readInput("D:/Projects/advent-of-code/2024/Python/Day2/input.txt")
    for i, line in enumerate(data):
        line = [int(x) for x in line.split(" ")]
        problem_count = 0
        problemIndexes.append([])

        if line[1] < line[0]: # Decreasing
            for j in range(len(line)):
                if j > 0:
                    if (line[j] - line[j-1]) not in [-1, -2, -3]:
                        problem_count += 1
                        problemIndexes[i].append(j)
        elif line[1] > line[0]: # Increasing
            for j in range(len(line)):
                if j > 0:
                    if (line[j] - line[j-1]) not in [1, 2, 3]:
                        problem_count += 1
                        problemIndexes[i].append(j)
        else: # This means that the second and first number are the same
            problem_count += 1
            problemIndexes[i].append(0)
            if line[2] < line[1]: # Decreasing
                for j in range(len(line)):
                    if j > 1:
                        if (line[j] - line[j-1]) not in [-1, -2, -3]:
                            problem_count += 1
                            problemIndexes[i].append(j)
            elif line[2] > line[1]: # Increasing
                for j in range(len(line)):
                    if j > 1:
                        if (line[j] - line[j-1]) not in [1, 2, 3]:
                            problem_count += 1
                            problemIndexes[i].append(j)
            else:
                problem_count += 1
                problemIndexes[i].append(1)

        problems.append(problem_count)

    for i in range(len(problems)):
        if problems[i] == 0:
            totalSafe += 1
        else:
            line = [int(x) for x in data[i].split(" ")]
            if test(line, problemIndexes[i][0]):
                totalSafe += 1
        
    return totalSafe
