import 'dart:io';

int hash(String s) {
  int v = 0;

  for (int i = 0; i < s.length; i++) {
    v += s.codeUnitAt(i);
    v *= 17;
    v %= 256;
  }

  return v;
}

void hashSumOfResults() {
  var input = File('Dart/Day15/input.txt').readAsStringSync();
  var inputList = input.split(',');

  var result = 0;
  for (var str in inputList) {
    result += hash(str);
  }

  print(result);
}
