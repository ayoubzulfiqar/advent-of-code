import 'dart:io';

Future<int> sumCalibration() async {
  final file = File('input.txt');
  try {
    final lines = await file.readAsLines();

    int sum = 0;
    for (final String line in lines) {
      final firstDigit = findFirstDigit(line);
      final lastDigit = findLastDigit(line);

      if (firstDigit != null && lastDigit != null) {
        final calibrationValue = firstDigit * 10 + lastDigit;
        sum += calibrationValue;
      }
    }

    return sum;
  } catch (e) {
    print('Error: $e');
    return 0;
  }
}

int? findFirstDigit(String s) {
  for (int i = 0; i < s.length; i++) {
    final char = s[i];
    if (isDigit(char)) {
      final digit = toInt(char);
      if (digit != null) {
        return digit;
      }
    }
  }
  return null;
}

int? findLastDigit(String s) {
  for (int i = s.length - 1; i >= 0; i--) {
    final String char = s[i];
    if (isDigit(char)) {
      final digit = toInt(char);
      if (digit != null) {
        return digit;
      }
    }
  }
  return null;
}

bool isDigit(String char) {
  return RegExp(r'\d').hasMatch(char);
}

int? toInt(String char) {
  if (RegExp(r'\d').hasMatch(char)) {
    return int.parse(char);
  }
  return null;
}
