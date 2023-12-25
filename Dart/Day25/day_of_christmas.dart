import 'dart:io';

class Node {
  late String name;
  late Map<Edge, Node> edges;
  bool traveled = false;

  Node(this.name) : edges = {};
}

class Edge {
  bool traveled = false;
}

class QueueItem {
  late Edge? edge;
  late Node node;
  late QueueItem? previous;

  QueueItem({required this.node, this.previous, required this.edge});
}

class Graph {
  late Map<String, Node> nodes;
  late List<Edge> edges;

  Graph()
      : nodes = {},
        edges = [];

  void resetNodes() {
    for (var node in nodes.values) {
      node.traveled = false;
    }
  }

  void resetEdges() {
    for (var edge in edges) {
      edge.traveled = false;
    }
  }

  bool removeShortestPath(Node source, Node dest) {
    var queue = <QueueItem>[QueueItem(node: source, edge: null)];
    var found = false;

    while (queue.isNotEmpty) {
      var current = queue.removeAt(0);

      if (current.node == dest) {
        for (var itr = current; itr.edge != null; itr = itr.previous!) {
          itr.edge?.traveled = true;
        }
        found = true;
        break;
      }

      for (var entry in current.node.edges.entries) {
        var e = entry.key;
        var n = entry.value;
        if (e.traveled || n.traveled) {
          continue;
        }
        n.traveled = true;
        queue.add(QueueItem(edge: e, node: n, previous: current));
      }
    }

    resetNodes();
    return found;
  }

  bool cutPaths(Node source, Node dest, int pathNum) {
    var complete = true;
    for (var i = 0; i < pathNum; i++) {
      if (!removeShortestPath(source, dest)) {
        complete = false;
        break;
      }
    }
    resetEdges();
    return complete;
  }

  List<Node> split(int cuts) {
    var g1 = <Node>[];
    var g2 = <Node>[];

    Node? source;
    for (var n in nodes.values) {
      source = n;
      break;
    }
    g1.add(source!);

    for (var dest in nodes.values) {
      if (source == dest) {
        continue;
      }

      if (cutPaths(source, dest, cuts + 1)) {
        g1.add(dest);
      } else {
        g2.add(dest);
      }
    }
    return [...g1, ...g2];
  }
}

Graph multipleOfTwoGroups(List<String> input) {
  var nodes = <String, Node>{};
  for (var line in input) {
    var parts = line.split(": ");
    var name = parts[0];
    nodes[name] = Node(name);
  }

  var edges = <Edge>[];
  for (var line in input) {
    var parts = line.split(": ");
    var sourceName = parts[0];
    var destNames = parts[1].split(" ");
    var source = nodes[sourceName]!;
    for (var destName in destNames) {
      var dest = nodes[destName] ?? Node(destName);
      var newEdge = Edge();
      edges.add(newEdge);
      source.edges[newEdge] = dest;
      dest.edges[newEdge] = source;
    }
  }

  return Graph()
    ..nodes = nodes
    ..edges = edges;
}

void main() {
  var input = <String>[];
  var file = File("Dart/Day25/input.txt");
  var lines = file.readAsLinesSync();
  input.addAll(lines);

  var groups = multipleOfTwoGroups(input);
  var result = groups.split(3);
  print(result.length * result.length);
}
