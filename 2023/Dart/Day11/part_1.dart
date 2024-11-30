import 'dart:io';

void main() {
  shortestPathGalaxyLength();
}

List<List<String>> addLines(List<List<String>> mat) {
  int size = mat.length;
  List<List<String>> nm = [];

  for (int i = 0; i < size; i++) {
    bool r = true;
    for (int j = 0; j < size; j++) {
      r = r && mat[i][j] != '#';
    }
    nm.add(List.from(mat[i]));
    if (r) {
      nm.add(List.filled(size, '.'));
    }
  }

  return nm;
}

List<List<String>> addColumn(List<List<String>> mat) {
  int size = mat.length;
  List<List<String>> nm = [];

  for (int i = 0; i < size; i++) {
    nm.add([]);
  }

  for (int i = 0; i < mat[0].length; i++) {
    bool c = true;
    for (int j = 0; j < size; j++) {
      c = c && mat[j][i] != '#';
      nm[j].add(mat[j][i]);
    }
    if (c) {
      for (int j = 0; j < size; j++) {
        nm[j].add('.');
      }
    }
  }

  return nm;
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

void shortestPathGalaxyLength() {
  var file = File("Dart/Day11/input.txt");
  var lines = file.readAsLinesSync();

  var mat = lines.map((line) => line.split('')).toList();
  mat = addLines(mat);
  mat = addColumn(mat);
  var points = findPoints(mat);

  var dists = <int>[];

  for (int i = 0; i < points.length; i++) {
    var p1 = points[i];
    for (int j = i + 1; j < points.length; j++) {
      var p2 = points[j];
      var x = (p1.x - p2.x).abs() + (p1.y - p2.y).abs();
      dists.add(x);
    }
  }

  var total = dists.fold(0, (sum, dist) => sum + dist);
  print(total);
}

class Point {
  int x, y;

  Point(this.x, this.y);
}
