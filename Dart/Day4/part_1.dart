import 'dart:io';

int scratchcardsWorth() {
  try {
    final file = File('Dart/Day4/input.txt');
    final lines = file.readAsLinesSync();
    var t = 0;

    for (final line in lines) {
      final parts = line.split(":");
      if (parts.length < 2) {
        continue;
      }

      var x = parts[1].trim();
      final arrays = x.split(" | ");
      if (arrays.length != 2) {
        continue;
      }

      final a = parseIntArray(arrays[0]);
      final b = parseIntArray(arrays[1]);

      var j = 0;
      for (final q in b) {
        for (final value in a) {
          if (q == value) {
            j++;
            break;
          }
        }
      }

      if (j > 0) {
        t += 1 << (j - 1);
      }
    }
    print(t);
    return t;
  } catch (e) {
    print('Error: $e');
    return 0;
  }
}

List<int> parseIntArray(String s) {
  final result = <int>[];
  final parts = s.split(' ');
  for (final part in parts) {
    final num = int.tryParse(part);
    if (num != null) {
      result.add(num);
    }
  }
  return result;
}
