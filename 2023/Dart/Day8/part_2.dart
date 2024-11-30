import 'dart:io';
import 'dart:convert';
import 'dart:typed_data';

Uint64List process(String contents) {
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
  final startNodes = <String>[];
  nodesMap.forEach((k, _) {
    if (k.endsWith('A')) {
      startNodes.add(k);
    }
  });
  final allSteps = <int>[];
  for (var startNode in startNodes) {
    String currentNode = startNode;
    int steps = 0;
    int turnIdx = 0;
    while (!currentNode.endsWith('Z')) {
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
    allSteps.add(steps);
  }
  BigInt result = BigInt.from(1);
  for (int x in allSteps) {
    result = result *
        BigInt.from(x) ~/
        BigInt.from(result.gcd(BigInt.from(x)).toInt());
  }
  return Uint64List.fromList([result.toInt()]);
}

void onlyNodesEndWithZ() {
  try {
    final file = File('Dart/Day8/input.txt');
    final contents = file.readAsStringSync();

    final result = process(contents);
    print('result = ${result[0]}');
  } catch (e) {
    print('Error opening file: $e');
    exit(1);
  }
}
