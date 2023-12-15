import 'dart:io';

int hashed(String s) {
  var v = 0;

  for (var ch in s.runes) {
    v += ch;
    v *= 17;
    v %= 256;
  }

  return v;
}

void powerOfResultLens() {
  List<List<String>> boxes = List.generate(256, (_) => [], growable: false);
  Map<String, int> focalLengths = {};

  List<String> instructions =
      File('Dart/Day15/input.txt').readAsStringSync().split(',');

  for (var instruction in instructions) {
    if (instruction.contains('-')) {
      var label = instruction.substring(0, instruction.length - 1);
      var index = hashed(label);

      boxes[index].removeWhere((l) => l == label);
    } else {
      var parts = instruction.split('=');
      var label = parts[0];
      var length = parts[1];

      var lengthValue = int.parse(length);

      var index = hashed(label);
      if (!boxes[index].contains(label)) {
        boxes[index].add(label);
      }

      focalLengths[label] = lengthValue;
    }
  }

  var total = 0;

  for (var i = 0; i < boxes.length; i++) {
    for (var j = 0; j < boxes[i].length; j++) {
      var label = boxes[i][j];
      total += (i + 1) * (j + 1) * focalLengths[label]!;
    }
  }

  print(total);
}
