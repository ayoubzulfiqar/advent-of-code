import 'dart:io';

void main() {
  try {
    final inputFile = File('input.txt');
    final inputStr = inputFile.readAsStringSync().trim();

    final stopwatch = Stopwatch()..start();
    final part2Result = elevsAndForkLifts(inputStr);
    stopwatch.stop();

    final elapsedSeconds = stopwatch.elapsedMicroseconds / 1000000;
    print(
      'Part 2 Result: $part2Result (Time: ${elapsedSeconds.toStringAsFixed(6)}s)',
    );
  } on FileSystemException {
    print('Error: input.txt not found');
  } catch (e) {
    print('Error reading input.txt: $e');
  }
}

String elevsAndForkLifts(String inputStr) {
  final lines = inputStr.trim().split('\n');
  final grid = lines.map((line) => line.split('')).toList();

  final height = grid.length;
  var totalRemoved = 0;

  const directions = [
    (-1, -1),
    (-1, 0),
    (-1, 1),
    (0, -1),
    (0, 1),
    (1, -1),
    (1, 0),
    (1, 1),
  ];

  while (true) {
    final cellsToRemove = <(int, int)>[];

    for (var row = 0; row < height; row++) {
      final currentRow = grid[row];
      final width = currentRow.length;

      for (var col = 0; col < width; col++) {
        if (currentRow[col] != '@') {
          continue;
        }

        var adjacentCount = 0;

        for (final dir in directions) {
          final neighborRow = row + dir.$1;
          final neighborCol = col + dir.$2;

          // Check bounds
          if (neighborRow >= 0 && neighborRow < height) {
            final neighborLine = grid[neighborRow];
            if (neighborCol >= 0 && neighborCol < neighborLine.length) {
              if (neighborLine[neighborCol] == '@') {
                adjacentCount++;
              }
            }
          }
        }

        if (adjacentCount < 4) {
          cellsToRemove.add((row, col));
        }
      }
    }

    if (cellsToRemove.isEmpty) {
      break;
    }

    for (final cell in cellsToRemove) {
      grid[cell.$1][cell.$2] = '.';
    }

    totalRemoved += cellsToRemove.length;
  }

  return totalRemoved.toString();
}


/*  With Tuple

import 'dart:io';

class Position {
  final int row;
  final int col;
  
  Position(this.row, this.col);
}

void main() {
  try {
    final inputFile = File('input.txt');
    final inputStr = inputFile.readAsStringSync().trim();
    
    final stopwatch = Stopwatch()..start();
    final part2Result = elevsAndForkLifts(inputStr);
    stopwatch.stop();
    
    final elapsedSeconds = stopwatch.elapsedMicroseconds / 1000000;
    print('Part 2 Result: $part2Result (Time: ${elapsedSeconds.toStringAsFixed(6)}s)');
  } on FileSystemException {
    print('Error: input.txt not found');
  } catch (e) {
    print('Error reading input.txt: $e');
  }
}

String elevsAndForkLifts(String inputStr) {
  final lines = inputStr.trim().split('\n');
  final grid = lines.map((line) => line.split('')).toList();
  
  final height = grid.length;
  var totalRemoved = 0;
  
  // Define all 8 neighbor directions as List of Positions
  final directions = [
    Position(-1, -1), Position(-1, 0), Position(-1, 1),
    Position(0, -1),                   Position(0, 1),
    Position(1, -1),  Position(1, 0),  Position(1, 1),
  ];
  
  while (true) {
    final cellsToRemove = <Position>[];
    
    for (var row = 0; row < height; row++) {
      final currentRow = grid[row];
      final width = currentRow.length;
      
      for (var col = 0; col < width; col++) {
        if (currentRow[col] != '@') {
          continue;
        }
        
        var adjacentCount = 0;
        
        for (final dir in directions) {
          final neighborRow = row + dir.row;
          final neighborCol = col + dir.col;
          
          if (neighborRow >= 0 && neighborRow < height) {
            final neighborLine = grid[neighborRow];
            if (neighborCol >= 0 && neighborCol < neighborLine.length) {
              if (neighborLine[neighborCol] == '@') {
                adjacentCount++;
              }
            }
          }
        }
        
        if (adjacentCount < 4) {
          cellsToRemove.add(Position(row, col));
        }
      }
    }
    
    if (cellsToRemove.isEmpty) {
      break;
    }
    
    for (final cell in cellsToRemove) {
      grid[cell.row][cell.col] = '.';
    }
    
    totalRemoved += cellsToRemove.length;
  }
  
  return totalRemoved.toString();
}


*/