import 'dart:io';

class IterativePathCounter {
  final Map<String, List<String>> _adjacency;
  final Map<String, int> _memo;

  IterativePathCounter(this._adjacency) : _memo = {};

  int countPaths(String startNode) {
    if (startNode == "out") return 1;

    final memo = _memo;
    memo["out"] = 1;

    final stack = <String>[startNode];
    final processing = <String>{};
    final visited = <String>{};

    while (stack.isNotEmpty) {
      final currentNode = stack.last;

      if (memo.containsKey(currentNode)) {
        stack.removeLast();
        continue;
      }

      if (processing.contains(currentNode)) {
        var total = 0;
        final children = _adjacency[currentNode] ?? [];
        for (final child in children) {
          total += memo[child]!;
        }
        memo[currentNode] = total;
        processing.remove(currentNode);
        visited.add(currentNode);
        stack.removeLast();
        continue;
      }

      if (visited.contains(currentNode)) {
        stack.removeLast();
        continue;
      }

      processing.add(currentNode);

      final children = _adjacency[currentNode] ?? [];
      var allChildrenComputed = true;

      for (final child in children) {
        if (!memo.containsKey(child) && child != "out") {
          allChildrenComputed = false;
          if (!processing.contains(child) && !visited.contains(child)) {
            stack.add(child);
          }
        }
      }

      if (allChildrenComputed) {
        var total = 0;
        for (final child in children) {
          total += memo[child]!;
        }
        memo[currentNode] = total;
        processing.remove(currentNode);
        visited.add(currentNode);
        stack.removeLast();
      }
    }

    return memo[startNode] ?? 0;
  }
}

int youPathOutIterative() {
  try {
    final file = File('input.txt');
    final lines = file.readAsLinesSync();

    final adjacencyList = <String, List<String>>{};

    for (var line in lines) {
      final parts = line.trim().split(RegExp(r'\s+'));
      if (parts.length < 2) continue;

      String key = parts[0];
      if (key.endsWith(':')) {
        key = key.substring(0, key.length - 1);
      }

      adjacencyList[key] = parts.sublist(1);
    }

    final pathCounter = IterativePathCounter(adjacencyList);
    final result = pathCounter.countPaths("you");

    print("Result: $result");

    return result;
  } catch (e) {
    print("Error: $e");
  }
  return -1;
}

void main() {
  youPathOutIterative();
}
