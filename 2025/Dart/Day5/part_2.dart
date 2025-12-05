import 'dart:convert';
import 'dart:io';

class Range {
  int low;
  int high;

  Range(this.low, this.high);
}

Future<int> FreshIDRangeIngredients() async {
  try {
    final file = File('input.txt');
    final scanner = file
        .openRead()
        .transform(utf8.decoder)
        .transform(LineSplitter());

    final ranges = <Range>[];
    var parsingRanges = true;

    await for (final line in scanner) {
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
          if (low > high) {
            stderr.writeln('Error: invalid range (low > high): $line');
            exit(1);
          }
          ranges.add(Range(low, high));
        } catch (e) {
          stderr.writeln('Error: invalid numbers in range: $line');
          exit(1);
        }
      } else {
        
        try {
          int.parse(trimmedLine);
        } catch (e) {
          stderr.writeln('Error: invalid ID: $line');
          exit(1);
        }
      }
    }

    if (ranges.isEmpty) {
      print('result = 0');
      return 0;
    }

    ranges.sort((a, b) => a.low.compareTo(b.low));

    final merged = <Range>[];
    var current = Range(ranges[0].low, ranges[0].high);

    for (var i = 1; i < ranges.length; i++) {
      final next = ranges[i];
      if (next.low <= current.high + 1) {
        if (next.high > current.high) {
          current.high = next.high;
        }
      } else {
        merged.add(Range(current.low, current.high));
        current = Range(next.low, next.high);
      }
    }
    merged.add(current);

    var total = 0;
    for (final r in merged) {
      total += r.high - r.low + 1;
    }

    final _ = '$total';

    print('result = $total');
    return total;
  } on FileSystemException catch (e) {
    stderr.writeln('Error opening file: $e');
    exit(1);
  }
}

void main() async {
  await FreshIDRangeIngredients();
}
