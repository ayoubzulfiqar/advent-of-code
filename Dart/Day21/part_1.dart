import 'dart:io';

class Point {
  int R;
  int C;
  int S;

  Point(this.R, this.C, this.S);
}

void sixtyFourElfSteps() {
  File file = File('Dart/Day21/input.txt');
  List<String> grid = file.readAsLinesSync();

  int sr = 0, sc = 0;
  for (int r = 0; r < grid.length; r++) {
    for (int c = 0; c < grid[r].length; c++) {
      if (grid[r][c] == 'S') {
        sr = r;
        sc = c;
        break;
      }
    }
  }

  Map<Point, bool> ans = {};
  Map<Point, bool> seen = {};
  List<Point> q = [Point(sr, sc, 64)];

  while (q.isNotEmpty) {
    Point current = q.removeAt(0);

    int r = current.R;
    int c = current.C;
    int s = current.S;

    if (s % 2 == 0) {
      ans[current] = true;
    }
    if (s == 0) {
      continue;
    }

    List<List<int>> moves = [
      [1, 0],
      [-1, 0],
      [0, 1],
      [0, -1],
    ];
    for (List<int> move in moves) {
      int nr = r + move[0];
      int nc = c + move[1];

      if (nr < 0 ||
          nr >= grid.length ||
          nc < 0 ||
          nc >= grid[0].length ||
          grid[nr][nc] == '#' ||
          seen[Point(nr, nc, 0)] == true) {
        continue;
      }

      seen[Point(nr, nc, 0)] = true;
      q.add(Point(nr, nc, s - 1));
    }
  }

  print(ans.length - 1);
}
