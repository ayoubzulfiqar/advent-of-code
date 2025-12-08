import 'dart:io';
// import 'dart:math';

const MAX = 188000000;

class Coord {
  final int x;
  final int y;
  final int z;

  Coord(this.x, this.y, this.z);
}

class Edge {
  final int a;
  final int b;
  final int dist;

  Edge(this.a, this.b, this.dist);
}

class DSU {
  final List<int> _parent;
  final List<int> _size;

  DSU(int n)
    : _parent = List<int>.generate(n, (i) => i),
      _size = List<int>.filled(n, 1);

  int find(int x) {
    if (_parent[x] != x) {
      _parent[x] = find(_parent[x]);
    }
    return _parent[x];
  }

  void union(int x, int y) {
    int rootX = find(x);
    int rootY = find(y);
    if (rootX == rootY) return;

    if (_size[rootX] < _size[rootY]) {
      int temp = rootX;
      rootX = rootY;
      rootY = temp;
    }

    _parent[rootY] = rootX;
    _size[rootX] += _size[rootY];
  }
}

int dist(Coord a, Coord b) {
  int dx = a.x - b.x;
  int dy = a.y - b.y;
  int dz = a.z - b.z;
  return dx * dx + dy * dy + dz * dz;
}

int minVal(int a, int b) => a < b ? a : b;

void multiplicationThreeLargestCircuits() {
  final file = File('input.txt');
  final lines = file.readAsLinesSync();

  final coords = <Coord>[];
  for (final line in lines) {
    if (line.isEmpty) continue;
    final parts = line.split(',');
    if (parts.length != 3) continue;

    final x = int.parse(parts[0]);
    final y = int.parse(parts[1]);
    final z = int.parse(parts[2]);
    coords.add(Coord(x, y, z));
  }

  final edges = <Edge>[];
  final n = coords.length;

  for (int i = 0; i < n - 1; i++) {
    for (int j = i + 1; j < n; j++) {
      final d = dist(coords[i], coords[j]);
      if (d < MAX) {
        edges.add(Edge(i, j, d));
      }
    }
  }

  edges.sort((a, b) => a.dist.compareTo(b.dist));

  final dsu = DSU(n);
  final limit = minVal(1000, edges.length);
  for (int i = 0; i < limit; i++) {
    dsu.union(edges[i].a, edges[i].b);
  }

  final circuitMap = <int, int>{};
  for (int i = 0; i < n; i++) {
    final root = dsu.find(i);
    circuitMap[root] = (circuitMap[root] ?? 0) + 1;
  }

  final circuits = circuitMap.values.toList();
  circuits.sort((a, b) => b.compareTo(a));

  final topCount = minVal(3, circuits.length);
  int product = 1;
  for (int i = 0; i < topCount; i++) {
    product *= circuits[i];
  }

  print(product);
}

void main() {
  multiplicationThreeLargestCircuits();
}
