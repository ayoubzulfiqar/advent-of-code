import 'dart:io';

int singleTachyonParticleTimelines() {
  try {
    final file = File('input.txt');
    final lines = file.readAsLinesSync();

    if (lines.isEmpty) {
      print('0');
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
      print('0');
      return 0;
    }

    final cache = <String, Map<int, int>>{};

    int timelines(int pos, List<String> remainingLines) {
      if (remainingLines.isEmpty) {
        return 1;
      }

      final key = remainingLines.join('');

      // Check cache
      if (cache.containsKey(key)) {
        final subCache = cache[key]!;
        if (subCache.containsKey(pos)) {
          return subCache[pos]!;
        }
      }

      final currentLine = remainingLines[0];
      int result;

      // Check if position is within bounds and has '^'
      if (pos >= 0 && pos < currentLine.length && currentLine[pos] == '^') {
        // Split to left and right
        final left = timelines(pos - 1, remainingLines.sublist(1));
        final right = timelines(pos + 1, remainingLines.sublist(1));
        result = left + right;
      } else {
        // Continue straight down
        result = timelines(pos, remainingLines.sublist(1));
      }

      // Store in cache
      if (!cache.containsKey(key)) {
        cache[key] = <int, int>{};
      }
      cache[key]![pos] = result;

      return result;
    }

    final result = timelines(start, lines.sublist(1));
    print('$result');
    return result;
  } catch (e) {
    print('Error reading input file: $e');
    return 0;
  }
}

void main() {
  singleTachyonParticleTimelines();
}
