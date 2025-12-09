import 'dart:io';

class Point {
  final int x;
  final int y;

  Point(this.x, this.y);
}

List<Point> getCoords(String input) {
  final coords = <Point>[];
  final lines = input.split('\n');

  for (final line in lines) {
    final trimmed = line.trim();
    if (trimmed.isEmpty) {
      continue;
    }

    final parts = trimmed.split(',');
    if (parts.length != 2) {
      continue;
    }

    final x = int.tryParse(parts[0].trim());
    final y = int.tryParse(parts[1].trim());

    if (x != null && y != null) {
      coords.add(Point(x, y));
    }
  }

  return coords;
}

int largestAreaForRectangle(List<Point> coords) {
  int maxArea = 0;

  for (int i = 0; i < coords.length; i++) {
    final p1 = coords[i];

    for (int j = i + 1; j < coords.length; j++) {
      final p2 = coords[j];

      final width = (p1.x > p2.x ? p1.x - p2.x : p2.x - p1.x) + 1;
      final height = (p1.y > p2.y ? p1.y - p2.y : p2.y - p1.y) + 1;

      final area = width * height;

      if (area > maxArea) {
        maxArea = area;
      }
    }
  }

  return maxArea;
}

void main() {
  final file = File('input.txt');
  final input = file.readAsStringSync();

  final coords = getCoords(input);
  final part1 = largestAreaForRectangle(coords);
  print('Part 1 Answer: $part1');
}
