import 'dart:io';

int totalScratchcards() {
  try {
    final file = File('Dart/Day4/input.txt');
    final lines = file.readAsLinesSync();
    final s = <String>[];

    for (final line in lines) {
      s.add(line);
    }

    var total = 0;
    final multiString = <int, int>{};

    for (var i = 0; i < s.length; i++) {
      multiString[i] = 1;
    }

    for (var i = 0; i < s.length; i++) {
      final game = s[i].split(': ');

      final games = game[1].split('|');

      final left = games[0].split(' ');
      final right = games[1].split(' ');

      final leftNumbers = parseNumbers(left);
      final rightNumbers = parseNumbers(right);

      final matches = intersection(leftNumbers, rightNumbers);

      final multiplier = multiString[i];

      total += multiplier!;

      for (var j = 0; j < matches.length; j++) {
        multiString[i + j + 1] = (multiString[i + j + 1] ?? 0) + multiplier;
      }
    }

    print(total);
    return total;
  } catch (e) {
    print('Error: $e');
    return 0;
  }
}

List<int> parseNumbers(List<String> nums) {
  final result = <int>[];
  for (final n in nums) {
    if (n.isEmpty) {
      continue;
    }
    final val = int.tryParse(n.trim());
    if (val != null) {
      result.add(val);
    }
  }
  return result;
}

List<int> intersection(List<int> a, List<int> b) {
  final result = <int>[];
  final seen = <int, bool>{};
  for (final v in a) {
    seen[v] = true;
  }
  for (final v in b) {
    if (seen[v] == true) {
      result.add(v);
    }
  }
  return result;
}
