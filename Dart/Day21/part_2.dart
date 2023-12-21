import 'dart:io';

const STEPS = 26501365;

List<List<int>> drs = [
  [1, 0],
  [0, 1],
  [-1, 0],
  [0, -1],
];

Map<String, int> bfs(List<String> m, int sx, int sy, int steps) {
  Map<String, int> dst = {'$sx,$sy': 0};
  List<List<int>> tVst = [
    [sx, sy, steps]
  ];

  while (tVst.isNotEmpty) {
    List<int> v = tVst.removeAt(0);

    for (List<int> d in drs) {
      int wx = v[0] + d[0];
      int wy = v[1] + d[1];

      int tCWx = wx, tCWy = wy;

      if (wy >= m.length) {
        tCWy = wy % m.length;
      }
      if (wy < 0) {
        tCWy = (wy % m.length + m.length) % m.length;
      }
      if (wx >= m[tCWy].length) {
        tCWx = wx % m[tCWy].length;
      }
      if (wx < 0) {
        tCWx = (wx % m[tCWy].length + m[tCWy].length) % m[tCWy].length;
      }

      if (m[tCWy][tCWx] != '#') {
        String sw = '$wx,$wy';
        if (!dst.containsKey(sw) && v[2] - 1 >= 0) {
          tVst.add([wx, wy, v[2] - 1]);
          dst[sw] = dst['${v[0]},${v[1]}']! + 1;
        }
      }
    }
  }
  return dst;
}

int firstTermArray(int n, List<int> p) {
  return p[0] +
      n * (p[1] - p[0]) +
      n * (n - 1) ~/ 2 * ((p[2] - p[1]) - (p[1] - p[0]));
}

void elfReachInGardenPlots() {
  int sx = -1, sy = 0;
  List<String> m = [];
  List<int> params = [];

  File file = File('Dart/Day21/input.txt');
  List<String> lines = file.readAsLinesSync();

  for (String line in lines) {
    m.add(line);
    for (int i = 0; i < m.last.length; i++) {
      if (m.last[i] == 'S') {
        sx = i;
      }
    }
    if (sx == -1) {
      sy++;
    }
  }

  for (int i = 0; i < m.length * 3; i++) {
    if (i % m.length == (m.length ~/ 2).floor()) {
      int r = 0;
      bfs(m, sx, sy, i).forEach((_, d) {
        if ((d + i % 2) % 2 == 0) {
          r++;
        }
      });
      params.add(r);
    }
  }

  print(firstTermArray((STEPS ~/ m.length).floor(), params));
}
