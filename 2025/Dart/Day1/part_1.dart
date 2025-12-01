import 'dart:io';

abstract class Direction {
  bool get isDirection => true;
}

class L extends Direction {
  final int n;
  L(this.n);

  @override
  String toString() => 'L($n)';
}

class R extends Direction {
  final int n;
  R(this.n);

  @override
  String toString() => 'R($n)';
}

Direction parseTurn(String s) {
  if (s.isEmpty) {
    throw FormatException('empty string');
  }

  final direction = s[0];
  final numberPart = s.substring(1);

  try {
    final n = int.parse(numberPart);

    switch (direction) {
      case 'L':
        return L(n);
      case 'R':
        return R(n);
      default:
        throw FormatException('malformed input: expected L<n> or R<n>');
    }
  } on FormatException {
    throw FormatException('invalid number: $numberPart');
  }
}

({int dialPrime, int zerosPrime}) step(int dial, int zeros, Direction turn) {
  int dialPrime;

  if (turn is L) {
    dialPrime = (dial - turn.n) % 100;
    if (dialPrime < 0) {
      dialPrime += 100;
    }
  } else if (turn is R) {
    dialPrime = (dial + turn.n) % 100;
  } else {
    throw ArgumentError('Unknown direction type');
  }

  int zerosPrime = zeros;
  if (dialPrime == 0) {
    zerosPrime++;
  }

  return (dialPrime: dialPrime, zerosPrime: zerosPrime);
}

int process(String content) {
  final lines = content.trim().split('\n');
  final turns = <Direction>[];

  for (var line in lines) {
    line = line.trim();
    if (line.isEmpty) {
      continue;
    }

    final turn = parseTurn(line);
    turns.add(turn);
  }

  int dial = 50;
  int zeros = 0;

  for (final turn in turns) {
    final result = step(dial, zeros, turn);
    dial = result.dialPrime;
    zeros = result.zerosPrime;
  }

  return zeros;
}

void ActualPasswordOfTheDoor() {
  try {
    final file = File('input.txt');
    final content = file.readAsStringSync();
    final result = process(content);
    print('result = $result');
  } on FileSystemException catch (e) {
    stderr.writeln('error reading file: $e');
    exit(1);
  } on FormatException catch (e) {
    stderr.writeln('error processing input: $e');
    exit(1);
  }
}

// void main() {
//   ActualPasswordOfTheDoor();
// }
