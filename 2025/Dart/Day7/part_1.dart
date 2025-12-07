import 'dart:convert';
import 'dart:io';

int beamSplitTime() {
  try {
    final file = File('input.txt');
    final lines = file.readAsLinesSync(encoding: utf8);

    if (lines.isEmpty) {
      print('Part 1: 0');
      return 0;
    }

    final firstLine = lines[0];
    int start = -1;
    for (int i = 0; i < firstLine.length; i++) {
      if (firstLine[i] == 'S') {
        start = i;
        break;
      }
    }

    if (start == -1) {
      print('Part 1: 0');
      return 0;
    }

    Map<int, bool> beams = <int, bool>{start: true};
    int count = 0;

    for (final line in lines.sublist(1)) {
      final splitters = <int, bool>{};
      for (int i = 0; i < line.length; i++) {
        if (line[i] == '^') {
          splitters[i] = true;
        }
      }

      final splits = <int, bool>{};
      int splitCount = 0;
      for (final pos in beams.keys) {
        if (splitters.containsKey(pos)) {
          splits[pos] = true;
          splitCount++;
        }
      }

      final newBeams = <int, bool>{};
      for (final pos in splits.keys) {
        if (pos + 1 < line.length) {
          newBeams[pos + 1] = true;
        }
        if (pos - 1 >= 0) {
          newBeams[pos - 1] = true;
        }
      }

      final remainingBeams = <int, bool>{};
      for (final pos in beams.keys) {
        if (!splits.containsKey(pos)) {
          remainingBeams[pos] = true;
        }
      }

      final finalBeams = <int, bool>{};
      for (final pos in remainingBeams.keys) {
        finalBeams[pos] = true;
      }
      for (final pos in newBeams.keys) {
        finalBeams[pos] = true;
      }

      count += splitCount;
      beams = finalBeams;
    }

    print('Part 1: $count');
    return count;
  } catch (e) {
    print('Error reading input file: $e');
    return 0;
  }
}

void main() {
  beamSplitTime();
}
