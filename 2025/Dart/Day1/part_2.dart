import 'dart:io';

abstract class TurnDirection {
  bool get isTurnDirection;
  int getAmount();
  TurnDirection withAmount(int amount);
}

class LeftTurn implements TurnDirection {
  final int amount;
  LeftTurn(this.amount);

  @override
  bool get isTurnDirection => true;

  @override
  int getAmount() => amount;

  @override
  TurnDirection withAmount(int amount) => LeftTurn(amount);

  @override
  String toString() => 'LeftTurn($amount)';
}

class RightTurn implements TurnDirection {
  final int amount;
  RightTurn(this.amount);

  @override
  bool get isTurnDirection => true;

  @override
  int getAmount() => amount;

  @override
  TurnDirection withAmount(int amount) => RightTurn(amount);

  @override
  String toString() => 'RightTurn($amount)';
}

TurnDirection parseDirection(String s) {
  if (s.isEmpty) {
    throw FormatException('empty string');
  }

  final direction = s[0];
  final numberPart = s.substring(1);

  try {
    final amount = int.parse(numberPart);

    switch (direction) {
      case 'L':
        return LeftTurn(amount);
      case 'R':
        return RightTurn(amount);
      default:
        throw FormatException('malformed input: expected L<n> or R<n>');
    }
  } on FormatException {
    throw FormatException('invalid number: $numberPart');
  }
}

({int pointer, int zeroCount}) applyTurn(
  int pointer,
  int zeroCount,
  TurnDirection direction,
) {
  return processRecursively(pointer, zeroCount, direction);
}

({int pointer, int zeroCount}) processRecursively(
  int p,
  int z,
  TurnDirection d,
) {
  // Base case: if turn amount is 0, return current state
  if (d.getAmount() == 0) {
    return (pointer: p, zeroCount: z);
  }

  // Handle the recursive step
  int nextPointer;
  int nextZeros;
  TurnDirection nextDirection;

  if (d is LeftTurn) {
    if (p == 0) {
      nextPointer = (p - 1) % 100;
      if (nextPointer < 0) {
        nextPointer += 100;
      }
      nextZeros = z + 1;
    } else {
      nextPointer = (p - 1) % 100;
      if (nextPointer < 0) {
        nextPointer += 100;
      }
      nextZeros = z;
    }
    nextDirection = d.withAmount(d.amount - 1);
  } else if (d is RightTurn) {
    if (p == 0) {
      nextPointer = (p + 1) % 100;
      nextZeros = z + 1;
    } else {
      nextPointer = (p + 1) % 100;
      nextZeros = z;
    }
    nextDirection = d.withAmount(d.amount - 1);
  } else {
    throw ArgumentError('Unknown direction type');
  }

  return processRecursively(nextPointer, nextZeros, nextDirection);
}

int computeResult(String input) {
  final lines = input.trim().split('\n');
  final directions = <TurnDirection>[];

  for (var line in lines) {
    line = line.trim();
    if (line.isEmpty) {
      continue;
    }

    final direction = parseDirection(line);
    directions.add(direction);
  }

  int pointer = 50;
  int zeroCount = 0;

  for (final direction in directions) {
    final result = applyTurn(pointer, zeroCount, direction);
    pointer = result.pointer;
    zeroCount = result.zeroCount;
  }

  return zeroCount;
}

void Method0x434C49434BToOpenTheDoor() {
  try {
    final file = File('input.txt');
    final data = file.readAsStringSync();
    final output = computeResult(data);
    print('result = $output');
  } on FileSystemException catch (e) {
    stderr.writeln('error reading file: $e');
    exit(1);
  } on FormatException catch (e) {
    stderr.writeln('error processing input: $e');
    exit(1);
  }
}

// void main() {
//   Method0x434C49434BToOpenTheDoor();
// }
