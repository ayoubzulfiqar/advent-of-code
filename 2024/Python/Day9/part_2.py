def resultingFileSystemChecksum() -> int:
    with open("./input.txt") as file:
        con = file.read()
    fileSystem = []
    file = 0

    # Create file system with object representation
    for i, char in enumerate(con):
        count = int(char)
        if i % 2 == 0:
            fileSystem.append({"file": file, "count": count})
            file += 1
        else:
            fileSystem.append({"file": -1, "count": count})

    reducedFileSystem = []

    for i in range(len(fileSystem)):
        # If processing a gap, try to fill it with as many files from right to left
        if fileSystem[i]["file"] == -1:
            scan = len(fileSystem) - 1
            while fileSystem[i]["count"] > 0 and scan > i:
                if (
                    fileSystem[scan]["file"] != -1
                    and fileSystem[scan]["count"] <= fileSystem[i]["count"]
                ):
                    reducedFileSystem.append(fileSystem[scan].copy())
                    fileSystem[i]["count"] -= fileSystem[scan]["count"]
                    fileSystem[scan]["file"] = -1
                    scan = len(fileSystem) - 1
                scan -= 1

            # If gap still exists, reflect it in reduced
            if fileSystem[i]["count"] != 0:
                reducedFileSystem.append(fileSystem[i])
        elif fileSystem[i]["count"] != 0:
            reducedFileSystem.append(fileSystem[i].copy())

    # Compute checksum with gaps in mind
    index = 0
    totalSum = 0
    current = reducedFileSystem.pop(0) if reducedFileSystem else None

    while reducedFileSystem or current:
        if current is None:
            break

        if current["file"] != -1:
            totalSum += current["file"] * index
            index += 1
            current["count"] -= 1
        else:
            index += current["count"]
            current["count"] = 0

        if current["count"] == 0:
            current = reducedFileSystem.pop(0) if reducedFileSystem else None
    print(totalSum)
    return totalSum
