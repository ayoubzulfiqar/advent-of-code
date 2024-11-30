import 'dart:io';
import 'dart:convert';

int process(String contents) {
  final scanner = LineSplitter.split(contents).iterator;
  scanner.moveNext(); // Read turns
  final turns = scanner.current;
  scanner.moveNext(); // Skip a line
  final nodeRe =
      RegExp(r'([A-Z][A-Z][A-Z]) = \(([A-Z][A-Z][A-Z]), ([A-Z][A-Z][A-Z])\)');
  final nodesMap = <String, Map<String, String>>{};
  while (scanner.moveNext()) {
    final line = scanner.current;
    final nodeCaps = nodeRe.firstMatch(line)!.groups([0, 1, 2, 3]);
    final source = nodeCaps[1]!;
    final lDest = nodeCaps[2]!;
    final rDest = nodeCaps[3]!;
    nodesMap[source] = {'Left': lDest, 'Right': rDest};
  }
  String currentNode = 'AAA';
  int steps = 0;
  var turnIdx = 0;
  while (currentNode != 'ZZZ') {
    var turn = turns[turnIdx];
    switch (turn) {
      case 'L':
        currentNode = nodesMap[currentNode]!['Left']!;
        break;
      case 'R':
        currentNode = nodesMap[currentNode]!['Right']!;
        break;
      default:
        throw Exception('unexpected char in turns');
    }
    steps++;
    turnIdx = (turnIdx + 1) % turns.length;
  }
  return steps;
}

void stepsToReachZ() {
  try {
    final file = File('Dart/Day8/input.txt');
    final contents = file.readAsStringSync();

    final result = process(contents);
    print('result = $result');
  } catch (e) {
    print('Error opening file: $e');
    exit(1);
  }
}
