import 'dart:io';

class Grid<T> {
  Map<int, Map<int, T>> _data = {};

  void set(int x, int y, T value) {
    if (!_data.containsKey(x)) {
      _data[x] = {};
    }
    _data[x]![y] = value;
  }

  T? get(int x, int y) {
    return _data[x]?[y];
  }

  Tuple<int, int> xRange() {
    final xValues = _data.keys;
    if (xValues.isEmpty) {
      return Tuple(0, 0);
    }
    final min =
        xValues.reduce((value, element) => value < element ? value : element);
    final max =
        xValues.reduce((value, element) => value > element ? value : element);
    return Tuple(min, max);
  }

  Tuple<int, int> yRange() {
    int minY = double.maxFinite.toInt();
    int maxY = -1;

    for (final xMap in _data.values) {
      final yValues = xMap.keys;
      if (yValues.isEmpty) {
        continue;
      }
      final min =
          yValues.reduce((value, element) => value < element ? value : element);
      final max =
          yValues.reduce((value, element) => value > element ? value : element);
      minY = minY < min ? minY : min;
      maxY = maxY > max ? maxY : max;
    }

    return Tuple(minY, maxY);
  }
}

class Tuple<A, B> {
  final A first;
  final B second;

  Tuple(this.first, this.second);
}

int ReflectionSummarizingAllNotes() {
  final input = File("./input.txt").readAsLinesSync();

  List<Grid<bool>> grids = [];

  int counter = 0;
  int offsetI = 0;
  for (int i = 0; i < input.length; i++) {
    final line = input[i];

    if (line.isEmpty) {
      offsetI = i + 1;
      counter++;
      continue;
    }

    if (grids.length <= counter) {
      grids.add(Grid<bool>());
    }

    for (int j = 0; j < line.length; j++) {
      grids[counter].set(j, i - offsetI, line[j] == '#');
    }
  }

  int total = 0;

  for (final g in grids) {
    final Tuple<int, int> xRange = g.xRange();
    final Tuple<int, int> yRange = g.yRange();

    for (int tryReflectX = xRange.first + 1;
        tryReflectX <= xRange.second;
        tryReflectX++) {
      bool allMatches = true;
      for (int i1 = xRange.first; i1 <= xRange.second; i1++) {
        int i2 = tryReflectX + (tryReflectX - i1) - 1;

        for (int j = yRange.first; j <= yRange.second; j++) {
          final a = g.get(i1, j);
          final b = g.get(i2, j);
          if (a == null || b == null) {
            continue;
          }
          if (a != b) {
            allMatches = false;
          }
        }
      }

      if (allMatches) {
        total += tryReflectX;
        break;
      }
    }

    for (int tryReflectY = yRange.first + 1;
        tryReflectY <= yRange.second;
        tryReflectY++) {
      bool allMatches = true;
      for (int j1 = yRange.first; j1 <= yRange.second; j1++) {
        int j2 = tryReflectY + (tryReflectY - j1) - 1;

        for (int i = xRange.first; i <= xRange.second; i++) {
          final a = g.get(i, j1);
          final b = g.get(i, j2);
          if (a == null || b == null) {
            continue;
          }
          if (a != b) {
            allMatches = false;
          }
        }
      }

      if (allMatches) {
        total += tryReflectY * 100;
        break;
      }
    }
  }
  print(total);
  return total;
}
