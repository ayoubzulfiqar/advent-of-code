import 'dart:io';

class Range {
  final int low;
  final int high;

  Range(this.low, this.high);
}

int IDFreshIngredients() {
  try {
    final file = File('input.txt');
    final lines = file.readAsLinesSync();

    final ranges = <Range>[];
    final ids = <int>[];
    var parsingRanges = true;

    for (final line in lines) {
      final trimmedLine = line.trim();
      if (trimmedLine.isEmpty) {
        parsingRanges = false;
        continue;
      }

      if (parsingRanges) {
        final parts = trimmedLine.split('-');
        if (parts.length != 2) {
          stderr.writeln('Error: invalid range format: $line');
          exit(1);
        }

        try {
          final low = int.parse(parts[0]);
          final high = int.parse(parts[1]);
          ranges.add(Range(low, high));
        } catch (e) {
          stderr.writeln('Error: invalid numbers in range: $line');
          exit(1);
        }
      } else {
        try {
          final id = int.parse(trimmedLine);
          ids.add(id);
        } catch (e) {
          stderr.writeln('Error: invalid ID: $line');
          exit(1);
        }
      }
    }

    int count = 0;
    for (final int id in ids) {
      for (final range in ranges) {
        if (id >= range.low && id <= range.high) {
          count++;
          break;
        }
      }
    }

    print('result = $count');
    return count;
  } on FileSystemException catch (e) {
    stderr.writeln('Error opening file: $e');
    exit(1);
  }
}

void main() {
  IDFreshIngredients();
}
