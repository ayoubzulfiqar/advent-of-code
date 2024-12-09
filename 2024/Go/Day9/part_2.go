package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type FileBlock struct {
	File  int
	Count int
}

func resultingFileSystemChecksum() int {
	// Open the input file
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the input
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	con := scanner.Text()
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Create file system with object representation
	var fileSystem []FileBlock
	fileID := 0

	for i, char := range con {
		count, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			fileSystem = append(fileSystem, FileBlock{File: fileID, Count: count})
			fileID++
		} else {
			fileSystem = append(fileSystem, FileBlock{File: -1, Count: count})
		}
	}

	var reducedFileSystem []FileBlock

	for i := 0; i < len(fileSystem); i++ {
		// If processing a gap, try to fill it with as many files from right to left
		if fileSystem[i].File == -1 {
			scan := len(fileSystem) - 1
			for fileSystem[i].Count > 0 && scan > i {
				if fileSystem[scan].File != -1 && fileSystem[scan].Count <= fileSystem[i].Count {
					reducedFileSystem = append(reducedFileSystem, fileSystem[scan])
					fileSystem[i].Count -= fileSystem[scan].Count
					fileSystem[scan].File = -1
					scan = len(fileSystem) - 1
				}
				scan--
			}

			// If gap still exists, reflect it in reduced
			if fileSystem[i].Count != 0 {
				reducedFileSystem = append(reducedFileSystem, fileSystem[i])
			}
		} else if fileSystem[i].Count != 0 {
			reducedFileSystem = append(reducedFileSystem, fileSystem[i])
		}
	}

	// Compute checksum with gaps in mind
	index := 0
	totalSum := 0

	for len(reducedFileSystem) > 0 {
		current := reducedFileSystem[0]
		reducedFileSystem = reducedFileSystem[1:]

		if current.File != -1 {
			totalSum += current.File * index
			index++
			current.Count--
		} else {
			index += current.Count
			current.Count = 0
		}

		if current.Count > 0 {
			reducedFileSystem = append([]FileBlock{current}, reducedFileSystem...)
		}
	}

	fmt.Println(totalSum)
	return totalSum
}

