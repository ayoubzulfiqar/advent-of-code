import 'dart:collection';
import 'dart:io';

class SubMapping {
  int source, size, offset;

  SubMapping({required this.source, required this.size, required this.offset});
}

class Mapping extends ListMixin<SubMapping> {
  final List<SubMapping> _list = [];

  @override
  int length = 0;

  @override
  SubMapping operator [](int index) => _list[index];

  @override
  void operator []=(int index, SubMapping value) {
    _list[index] = value;
  }

  @override
  void add(SubMapping value) {
    _list.add(value);
    length = _list.length;
  }
}

int lowestLocationSeedNumber() {
  try {
    final file = File('Dart/Day5/input.txt');
    final lines = file.readAsLinesSync();
    final seeds = <int>[];
    final mappings = <Mapping>[];
    var currentMapping = Mapping();

    for (final line in lines) {
      if (line.isEmpty) {
        continue;
      }

      if (line.contains("seeds: ")) {
        final seedsString = line.split("seeds: ");
        final seedList = seedsString[1].split(" ");
        for (final seedItem in seedList) {
          final seed = int.tryParse(seedItem);
          if (seed != null) {
            seeds.add(seed);
          }
        }
        continue;
      }

      if (line.contains("-")) {
        if (currentMapping.isNotEmpty) {
          mappings.add(currentMapping);
        }
        currentMapping = Mapping();
        continue;
      }

      final values = line.split(' ');

      final source = int.tryParse(values[1]);
      final size = int.tryParse(values[2]);

      if (source == null || size == null) {
        print("Error converting source or size to integer");
        return 0;
      }

      final offset = sti(values[0]) - sti(values[1]);

      currentMapping.add(SubMapping(
        source: source,
        size: size,
        offset: offset,
      ));
    }

    if (currentMapping.isNotEmpty) {
      mappings.add(currentMapping);
    }

    var lowest = -1;

    for (final seed in seeds) {
      var val = seed;

      for (final mapping in mappings) {
        for (final subMapping in mapping) {
          if (val >= subMapping.source &&
              val <= subMapping.source + subMapping.size) {
            val += subMapping.offset;
            break;
          }
        }
      }

      if (lowest == -1 || val < lowest) {
        lowest = val;
      }
    }

    print(lowest);
    return lowest;
  } catch (e) {
    print('Error: $e');
    return 0;
  }
}

int sti(String s) {
  final i = int.tryParse(s);
  if (i == null) {
    print("Error converting string to integer");
  }
  return i ?? 0;
}
