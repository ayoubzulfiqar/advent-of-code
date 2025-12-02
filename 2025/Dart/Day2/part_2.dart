import 'dart:io';
import 'dart:math';

class MultiRepeatRange {
  final int lowerBound;
  final int upperBound;

  MultiRepeatRange(this.lowerBound, this.upperBound);
}

List<MultiRepeatRange> parseMultiRepeatRanges(String inputData) {
  final rangeCollection = <MultiRepeatRange>[];
  final pairStrings = inputData.split(',');

  for (final rangeStr in pairStrings) {
    final bounds = rangeStr.split('-');
    if (bounds.length != 2) {
      continue;
    }

    try {
      final lowerValue = int.parse(bounds[0].trim());
      final upperValue = int.parse(bounds[1].trim());
      rangeCollection.add(MultiRepeatRange(lowerValue, upperValue));
    } catch (e) {
      continue;
    }
  }

  return rangeCollection;
}

int digitCount(int value) {
  if (value == 0) {
    return 1;
  }
  return (log(value) / ln10).floor() + 1;
}

bool isEvenlyDivisible(int dividend, int divisor) {
  return dividend % divisor == 0;
}

Set<int> discoverMultiRepeatingIDs(int minID, int maxID) {
  final foundIDs = <int>{};

  final minDigits = digitCount(minID);
  final maxDigits = digitCount(maxID);

  for (var repeatTimes = 2; repeatTimes <= maxDigits; repeatTimes++) {
    int segmentMinLength;
    if (isEvenlyDivisible(minDigits, repeatTimes)) {
      segmentMinLength = minDigits ~/ repeatTimes;
      if (segmentMinLength < 1) {
        segmentMinLength = 1;
      }
    } else {
      segmentMinLength = minDigits ~/ repeatTimes + 1;
    }

    final segmentMaxLength = maxDigits ~/ repeatTimes;

    for (
      var segmentLength = segmentMinLength;
      segmentLength <= segmentMaxLength;
      segmentLength++
    ) {
      var minSegmentValue = pow(10, segmentLength - 1).toInt();
      final divisor = pow(10, segmentLength * (repeatTimes - 1)).toInt();

      if (minSegmentValue < minID ~/ divisor) {
        minSegmentValue = minID ~/ divisor;
      }

      var maxSegmentValue = pow(10, segmentLength).toInt() - 1;
      if (maxSegmentValue > maxID ~/ divisor) {
        maxSegmentValue = maxID ~/ divisor;
      }

      if (minSegmentValue > maxSegmentValue) {
        continue;
      }

      for (
        var segmentNumber = minSegmentValue;
        segmentNumber <= maxSegmentValue;
        segmentNumber++
      ) {
        var constructedID = segmentNumber;
        for (var repetition = 1; repetition < repeatTimes; repetition++) {
          constructedID =
              constructedID * pow(10, segmentLength).toInt() + segmentNumber;
        }

        if (constructedID >= minID && constructedID <= maxID) {
          foundIDs.add(constructedID);
        }
      }
    }
  }

  return foundIDs;
}

int computeMultiRepeatIDSum(List<MultiRepeatRange> ranges) {
  var cumulativeSum = 0;

  for (final currentRange in ranges) {
    final multiRepeatIDs = discoverMultiRepeatingIDs(
      currentRange.lowerBound,
      currentRange.upperBound,
    );

    for (final id in multiRepeatIDs) {
      cumulativeSum += id;
    }
  }

  return cumulativeSum;
}

void solvePart2() {
  try {
    final file = File('input.txt');
    final fileData = file.readAsStringSync();
    final rangeCollection = parseMultiRepeatRanges(fileData);
    final totalResult = computeMultiRepeatIDSum(rangeCollection);
    print('Part 2 Result: $totalResult');
  } catch (e) {
    print('Error: $e');
  }
}

void main() {
  solvePart2();
}
