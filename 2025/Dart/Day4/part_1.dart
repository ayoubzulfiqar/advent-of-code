import 'dart:io';

void main() {
  try {
    final inputFile = File('input.txt');
    final inputStr = inputFile.readAsStringSync().trim();

    final stopwatch = Stopwatch()..start();
    final part1Result = forkListRolls(inputStr);
    stopwatch.stop();

    print(
      'Part 1 Result: $part1Result (Time: ${stopwatch.elapsedMicroseconds / 1000000}s)',
    );
  } on FileSystemException catch (e) {
    print('Error: ${e.message}');
  } catch (e) {
    print('Error reading input.txt: $e');
  }
}

String forkListRolls(String inputStr) {
  final lines = inputStr.trim().split('\n');

  final grid = lines.map((line) => line.split('')).toList();

  final height = grid.length;
  var accessibleCount = 0;

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

  for (var row = 0; row < height; row++) {
    final currentLine = grid[row];
    final width = currentLine.length;

    for (var col = 0; col < width; col++) {
      if (currentLine[col] != '@') {
        continue;
      }

      var adjacentCount = 0;

      for (final dir in directions) {
        final neighborRow = row + dir.$1;
        final neighborCol = col + dir.$2;

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
        accessibleCount++;
      }
    }
  }

  return accessibleCount.toString();
}
