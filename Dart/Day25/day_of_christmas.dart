import 'dart:collection';
import 'dart:io';

void main() {
  var input = File('Dart/Day25/input.txt').readAsLinesSync();

  var dotFile = StringBuffer();
  dotFile.writeln("graph {");

  var wires = <String, Set<String>>{};
  for (var line in input) {
    var parts = line.split(": ");
    var wire = parts[0].trim();
    var connections = parts[1].trim().split(" ").toList();

    wires[wire] ??= {}; // Ensure wires[wire] is initialized as an empty set.

    for (var connection in connections) {
      // Generate content for a GraphViz graph.
      dotFile.writeln("    $wire -- $connection");

      wires[connection] ??= {};

      wires[wire]?.add(connection);
      wires[connection]?.add(wire);
    }
  }

  dotFile.writeln("}");
  File('graph.dot').writeAsStringSync(dotFile.toString());

  // Edges to be removed
  const a1 = "cbl";
  const b1 = "vmq";

  const a2 = "bvz";
  const b2 = "nvf";

  const a3 = "klk";
  const b3 = "xgz";

  wires[a1]?.remove(b1);
  wires[b1]?.remove(a1);
  wires[a2]?.remove(b2);
  wires[b2]?.remove(a2);
  wires[a3]?.remove(b3);
  wires[b3]?.remove(a3);

  // BFS to count the number of reachable nodes from a random single node of the removed pairs
  var visited = Set<String>();
  var queue = Queue<String>();
  queue.add(a1);

  while (queue.isNotEmpty) {
    var current = queue.removeFirst();

    if (!visited.add(current)) {
      continue;
    }

    for (var connection in wires[current] ?? {}) {
      queue.add(connection);
    }
  }

  // Calculate the size of the other subset
  var part1 = visited.length * (wires.length - visited.length);
  print("Part 1: $part1");
}
