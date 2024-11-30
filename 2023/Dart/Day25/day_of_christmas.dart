import 'dart:collection';
import 'dart:io';
import 'dart:math';

typedef NetworkEdge = ({String node, double weight});

int multipleSizeOfTwoGroups() {
  final graph = <String, List<NetworkEdge>>{};
  final file = File("Dart/Day25/input.txt").readAsLinesSync();

  for (final line in file.where((line) => line.isNotEmpty)) {
    final parts = line.split(':');

    final (from, to) = (parts[0], parts[1].trim().split(' '));
    graph
        .putIfAbsent(from, () => [])
        .addAll(to.map((node) => (node: node, weight: 1)));

    for (final node in to) {
      graph.putIfAbsent(node, () => []).add((node: from, weight: 1));
    }
  }

  var ans = 0;
  for (final source in graph.keys) {
    for (final sink in graph.keys) {
      if (source != sink) {
        final (:maxFlow, :minCutPartitionSize) = edmondKarp(
          graph,
          source,
          sink,
        );
        if (maxFlow == 3) {
          ans = minCutPartitionSize * (graph.keys.length - minCutPartitionSize);
          break;
        }
      }
    }
  }
  print(ans);
  return ans;
}

({bool isPath, List<String> path}) bfs(
  Map<String, List<NetworkEdge>> residualGraph,
  String source,
  String sink,
) {
  final visited = <String>{};
  final queue = Queue<String>();
  final paths = <String, List<String>>{};

  queue.add(source);
  visited.add(source);
  paths[source] = [source];

  while (queue.isNotEmpty) {
    final node = queue.removeFirst();

    if (node == sink) {
      return (isPath: true, path: paths[node]!);
    }

    for (final edge in residualGraph[node]!) {
      if (!visited.contains(edge.node) && edge.weight > 0) {
        queue.add(edge.node);
        visited.add(edge.node);
        paths[edge.node] = [...paths[node]!, edge.node];
      }
    }
  }

  return (isPath: false, path: []);
}

void dfs(
  Map<String, List<NetworkEdge>> residualGraph,
  String node,
  Set<String> visited,
) {
  visited.add(node);

  for (final edge in residualGraph[node]!) {
    if (!visited.contains(edge.node) && edge.weight > 0) {
      dfs(residualGraph, edge.node, visited);
    }
  }
}

({int maxFlow, int minCutPartitionSize}) edmondKarp(
  Map<String, List<NetworkEdge>> graph,
  String source,
  String sink,
) {
  final residualGraph = <String, List<NetworkEdge>>{};
  for (final node in graph.keys) {
    residualGraph[node] = [];
    for (final edge in graph[node]!) {
      residualGraph[node]!.add(edge);
    }
  }

  var maxFlow = 0.0;
  while (true) {
    final (:isPath, :path) = bfs(residualGraph, source, sink);
    if (!isPath) {
      break;
    }

    var minCapacity = double.infinity;
    for (var i = 0; i < path.length - 1; i++) {
      final from = path[i];
      final to = path[i + 1];

      for (final edge in residualGraph[from]!) {
        if (edge.node == to) {
          minCapacity = min(minCapacity, edge.weight);
          break;
        }
      }
    }

    for (var i = 0; i < path.length - 1; i++) {
      final from = path[i];
      final to = path[i + 1];

      for (final edge in residualGraph[from]!) {
        if (edge.node == to) {
          residualGraph[from]!.remove(edge);
          residualGraph[from]!.add(
            (node: to, weight: edge.weight - minCapacity),
          );
          break;
        }
      }

      for (final edge in residualGraph[to]!) {
        if (edge.node == from) {
          residualGraph[to]!.remove(edge);
          residualGraph[to]!.add(
            (node: from, weight: edge.weight + minCapacity),
          );
          break;
        }
      }
    }

    maxFlow += minCapacity;
  }

  final visited = <String>{};
  dfs(residualGraph, source, visited);

  return (
    maxFlow: maxFlow.toInt(),
    minCutPartitionSize: visited.length,
  );
}
