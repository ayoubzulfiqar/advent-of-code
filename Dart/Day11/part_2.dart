import 'dart:io';

void main() {
  shortestPathLengthInPairGalaxies();
}

Map<String, dynamic> empty(List<List<String>> mat) {
  int size = mat.length;
  Map<int, dynamic> er = {};
  Map<int, dynamic> ec = {};

  for (int i = 0; i < size; i++) {
    bool r = true;
    bool c = true;
    for (int j = 0; j < size; j++) {
      r = r && mat[i][j] != '#';
      c = c && mat[j][i] != '#';
    }
    if (r) {
      er[i] = {};
    }
    if (c) {
      ec[i] = {};
    }
  }
  return {'er': er, 'ec': ec};
}

List<Point> findPoints(List<List<String>> mat) {
  List<Point> p = [];

  for (int i = 0; i < mat.length; i++) {
    for (int j = 0; j < mat[i].length; j++) {
      if (mat[i][j] == '#') {
        p.add(Point(i, j));
      }
    }
  }

  return p;
}

void shortestPathLengthInPairGalaxies() {
  var file = File("Dart/Day11/input.txt");
  var lines = file.readAsLinesSync();

  var mat = lines.map((line) => line.split('')).toList();
  var result = empty(mat);
  Map<int, dynamic> el = result['er'];
  Map<int, dynamic> ec = result['ec'];
  var points = findPoints(mat);
  var distS = <int>[];
  var weight = 1000000;

  for (int i = 0; i < points.length; i++) {
    var p1 = points[i];
    for (int j = i + 1; j < points.length; j++) {
      var p2 = points[j];
      var x = 0;

      for (var k = min(p1.x, p2.x); k < max(p1.x, p2.x); k++) {
        if (el.containsKey(k)) {
          x += weight;
        } else {
          x++;
        }
      }

      for (var k = min(p1.y, p2.y); k < max(p1.y, p2.y); k++) {
        if (ec.containsKey(k)) {
          x += weight;
        } else {
          x++;
        }
      }

      distS.add(x);
    }
  }

  var total = distS.fold(0, (sum, dist) => sum + dist);
  print(total);
}

class Point {
  int x, y;

  Point(this.x, this.y);
}

int min(int a, int b) => a < b ? a : b;

int max(int a, int b) => a > b ? a : b;
