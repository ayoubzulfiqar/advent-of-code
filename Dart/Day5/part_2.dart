import 'dart:io';

class SubMapping {
  final int source;
  late int size;
  late int offset;
  SubMapping({required this.source, required this.size, required this.offset});
}

class Mapping {
  List<SubMapping> mappings = [];

  int convert(int input) {
    for (var c in mappings) {
      if (input >= c.source && input <= c.source + c.size) {
        return input + c.offset;
      }
    }
    return input;
  }
}

List<List<int>> splitRangesAt(List<List<int>> ranges, int n) {
  for (var i = 0; i < ranges.length; i++) {
    var ss = ranges[i];
    if (n > ss[0] && n <= ss[1]) {
      ranges[i][1] = n - 1;
      ranges.add([n, ss[1]]);
      return ranges;
    }
  }
  return ranges;
}

int findMin(int a, int b) {
  return (a < b) ? a : b;
}

int lowestInitialSeedNumber() {
  try {
    var file = File('Dart/Day5/input.txt');
    var lines = file.readAsLinesSync();
    List<List<int>> seedPairs = [];
    List<Mapping> mappings = [];

    var currentMapping = Mapping();
    for (var line in lines) {
      if (line.isEmpty) {
        continue;
      }

      if (line.contains('seeds: ')) {
        var seedsString = line.split('seeds: ');
        var seedList = seedsString[1].split(' ');

        for (var i = 0; i < seedList.length; i += 2) {
          seedPairs.add([
            int.parse(seedList[i]),
            int.parse(seedList[i]) + int.parse(seedList[i + 1])
          ]);
        }
        continue;
      }

      if (line.contains('-')) {
        if (currentMapping.mappings.isNotEmpty) {
          mappings.add(currentMapping);
        }
        currentMapping = Mapping();
        continue;
      }

      var values = line.split(' ');

      currentMapping.mappings.add(SubMapping(
        source: int.parse(values[1]),
        size: int.parse(values[2]),
        offset: int.parse(values[0]) - int.parse(values[1]),
      ));
    }
    if (currentMapping.mappings.isNotEmpty) {
      mappings.add(currentMapping);
    }

    var lowest = -1;

    for (var pair in seedPairs) {
      var ranges = [pair];

      for (var mapping in mappings) {
        for (var subMapping in mapping.mappings) {
          ranges = splitRangesAt(ranges, subMapping.source);
        }

        for (var i = 0; i < ranges.length; i++) {
          ranges[i][0] = mapping.convert(ranges[i][0]);
          ranges[i][1] = mapping.convert(ranges[i][1]);
        }
      }

      for (var i = 0; i < ranges.length; i++) {
        if (lowest == -1) {
          lowest = ranges[i][0];
        }

        lowest = findMin(lowest, ranges[i][0]);
      }
    }
    print(lowest);
    return lowest;
  } catch (e) {
    print('Error opening file: $e');
    return 0;
  }
}

void main() {
  lowestInitialSeedNumber();
}
