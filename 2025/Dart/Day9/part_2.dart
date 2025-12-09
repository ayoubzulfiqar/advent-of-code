import 'dart:io';
import 'dart:math';

class Point {
  final int x;
  final int y;

  Point(this.x, this.y);

  @override
  bool operator ==(Object other) =>
      other is Point && x == other.x && y == other.y;

  @override
  int get hashCode => Object.hash(x, y);
}

List<Point> getCoordinates(String input) {
  final coords = <Point>[];
  final lines = input.split('\n');

  for (final line in lines) {
    final trimmed = line.trim();
    if (trimmed.isEmpty) continue;

    final commaIdx = trimmed.indexOf(',');
    if (commaIdx == -1) continue;

    final xStr = trimmed.substring(0, commaIdx).trim();
    final yStr = trimmed.substring(commaIdx + 1).trim();

    final x = int.tryParse(xStr);
    final y = int.tryParse(yStr);

    if (x != null && y != null) {
      coords.add(Point(x, y));
    }
  }

  return coords;
}

int largestAreaWithGreenAndRedTiles(List<Point> coords) {
  if (coords.isEmpty) return 0;

  int maxArea = 0;
  final n = coords.length;

  final polyEdges = <({Point u, Point v})>[];
  int minX = coords[0].x;
  int maxX = coords[0].x;
  int minY = coords[0].y;
  int maxY = coords[0].y;

  for (int i = 0; i < n; i++) {
    final u = coords[i];
    final v = coords[(i + 1) % n];
    polyEdges.add((u: u, v: v));

    minX = min(minX, u.x);
    maxX = max(maxX, u.x);
    minY = min(minY, u.y);
    maxY = max(maxY, u.y);
  }

  final pipCache = <String, bool>{};

  for (int i = 0; i < n; i++) {
    final p1 = coords[i];
    final x1 = p1.x;
    final y1 = p1.y;

    for (int j = i + 1; j < n; j++) {
      final p2 = coords[j];

      final rx1 = min(x1, p2.x);
      final rx2 = max(x1, p2.x);
      final ry1 = min(y1, p2.y);
      final ry2 = max(y1, p2.y);

      final width = rx2 - rx1 + 1;
      final height = ry2 - ry1 + 1;
      final area = width * height;

      if (area <= maxArea) continue;

      if (rx1 < minX || rx2 > maxX || ry1 < minY || ry2 > maxY) continue;

      if (_isValidRectOptimized(
        rx1,
        rx2,
        ry1,
        ry2,
        coords,
        polyEdges,
        pipCache,
      )) {
        maxArea = area;
      }
    }
  }

  return maxArea;
}

bool _isValidRectOptimized(
  int x1,
  int x2,
  int y1,
  int y2,
  List<Point> poly,
  List<({Point u, Point v})> polyEdges,
  Map<String, bool> pipCache,
) {
  final mx = x1 + x2;
  final my = y1 + y2;
  final cacheKey = '$mx,$my';

  final cached = pipCache[cacheKey];
  if (cached != null) {
    if (!cached) return false;
  } else {
    final inPoly = _isPointInPolyOptimized(mx, my, poly, polyEdges);
    pipCache[cacheKey] = inPoly;
    if (!inPoly) return false;
  }

  for (final edge in polyEdges) {
    final u = edge.u;
    final v = edge.v;

    if (u.x == v.x) {
      final ex = u.x;
      if (ex <= x1 || ex >= x2) continue;

      final eyMin = min(u.y, v.y);
      final eyMax = max(u.y, v.y);

      final overlapYMin = max(y1, eyMin);
      final overlapYMax = min(y2, eyMax);
      if (overlapYMin < overlapYMax) return false;
    } else {
      final ey = u.y;
      if (ey <= y1 || ey >= y2) continue;

      final exMin = min(u.x, v.x);
      final exMax = max(u.x, v.x);

      final overlapXMin = max(x1, exMin);
      final overlapXMax = min(x2, exMax);
      if (overlapXMin < overlapXMax) return false;
    }
  }

  return true;
}

bool _isPointInPolyOptimized(
  int x,
  int y,
  List<Point> poly,
  List<({Point u, Point v})> polyEdges,
) {
  final n = poly.length;

  for (final edge in polyEdges) {
    final u = edge.u;
    final v = edge.v;

    final ux2 = u.x * 2;
    final uy2 = u.y * 2;
    final vx2 = v.x * 2;
    final vy2 = v.y * 2;

    if (u.x == v.x && ux2 == x) {
      final minY = min(uy2, vy2);
      final maxY = max(uy2, vy2);
      if (minY <= y && y <= maxY) return true;
    } else if (u.y == v.y && uy2 == y) {
      final minX = min(ux2, vx2);
      final maxX = max(ux2, vx2);
      if (minX <= x && x <= maxX) return true;
    }
  }

  int intersections = 0;
  int j = n - 1;

  for (int i = 0; i < n; i++) {
    final u = poly[i];
    final v = poly[j];

    if (u.x == v.x) {
      final uy2 = u.y * 2;
      final vy2 = v.y * 2;
      final ex = u.x * 2;

      final minY = min(uy2, vy2);
      final maxY = max(uy2, vy2);

      if (minY <= y && y < maxY && ex > x) {
        intersections++;
      }
    }
    j = i;
  }

  return intersections.isOdd;
}

void main() {
  final file = File('input.txt');
  final input = file.readAsStringSync();

  final coords = getCoordinates(input);
  final part2 = largestAreaWithGreenAndRedTiles(coords);
  print('Part 2 Answer: $part2');
}
