def resultingFileChecksum() -> int:
    with open("./input.txt") as file:
        con = file.read()
    fileSystem = []
    file = 0

    # Create initial file system
    for i, char in enumerate(con):
        for _ in range(int(char)):
            fileSystem.append(file if i % 2 == 0 else -1)
        if i % 2 == 0:
            file += 1

    # Keep swapping empty spaces from bottom to top
    bottom, top = 0, len(fileSystem) - 1
    while bottom < top:
        if fileSystem[bottom] == -1:
            while fileSystem[top] == -1:
                top -= 1

            if top < bottom:  # Edge case to prevent unnecessary swaps
                break

            fileSystem[bottom], fileSystem[top] = fileSystem[top], -1

        bottom += 1

    # Find checksum of the file system
    totalSum = 0
    for i, value in enumerate(fileSystem):
        if value == -1:
            break
        totalSum += value * i
    print(totalSum)
    return totalSum



