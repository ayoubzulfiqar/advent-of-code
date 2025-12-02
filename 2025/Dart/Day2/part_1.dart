import 'dart:io';
import 'dart:math';

class IDRange {
  final int startID;
  final int endID;

  IDRange(this.startID, this.endID);
}

List<IDRange> parseRanges(String input) {
  final ranges = <IDRange>[];
  final linePairs = input.split(',');

  for (final pair in linePairs) {
    final ids = pair.split('-');
    if (ids.length != 2) {
      continue;
    }

    try {
      final startID = int.parse(ids[0].trim());
      final endID = int.parse(ids[1].trim());
      ranges.add(IDRange(startID, endID));
    } catch (e) {
      continue;
    }
  }

  return ranges;
}

int countDigits(int num) {
  if (num == 0) {
    return 1;
  }
  return (log(num) / ln10).floor() + 1;
}

bool isDivisibleBy(int dividend, int divisor) {
  return dividend % divisor == 0;
}

List<int> findInvalidIDs(int startID, int endID) {
  final invalidIDs = <int>[];

  final startDigits = countDigits(startID);
  final endDigits = countDigits(endID);

  int minHalfDigits;
  if (isDivisibleBy(startDigits, 2)) {
    minHalfDigits = startDigits ~/ 2;
    if (minHalfDigits < 1) {
      minHalfDigits = 1;
    }
  } else {
    minHalfDigits = startDigits ~/ 2 + 1;
  }

  final maxHalfDigits = endDigits ~/ 2;

  for (
    var halfLength = minHalfDigits;
    halfLength <= maxHalfDigits;
    halfLength++
  ) {
    var firstHalfLowerBound = pow(10, halfLength - 1).toInt();
    final startDivisor = pow(10, halfLength).toInt();

    if (firstHalfLowerBound < startID ~/ startDivisor) {
      firstHalfLowerBound = startID ~/ startDivisor;
    }

    var firstHalfUpperBound = pow(10, halfLength).toInt() - 1;
    if (firstHalfUpperBound > endID ~/ startDivisor) {
      firstHalfUpperBound = endID ~/ startDivisor;
    }

    if (firstHalfLowerBound > firstHalfUpperBound) {
      continue;
    }

    for (
      var firstHalf = firstHalfLowerBound;
      firstHalf <= firstHalfUpperBound;
      firstHalf++
    ) {
      final repeatingID = firstHalf * pow(10, halfLength).toInt() + firstHalf;

      if (repeatingID >= startID && repeatingID <= endID) {
        invalidIDs.add(repeatingID);
      }
    }
  }

  return invalidIDs;
}

int calculateInvalidIDSum(List<IDRange> idRanges) {
  var totalSum = 0;

  for (final idRange in idRanges) {
    final invalidIDs = findInvalidIDs(idRange.startID, idRange.endID);
    for (final id in invalidIDs) {
      totalSum += id;
    }
  }

  return totalSum;
}

void solvePart1() {
  try {
    final file = File('input.txt');
    final data = file.readAsStringSync();
    final idRanges = parseRanges(data);
    final solution = calculateInvalidIDSum(idRanges);
    print('Part 1 Result: $solution');
  } catch (e) {
    print('Error: $e');
  }
}

void main() {
  solvePart1();
}
