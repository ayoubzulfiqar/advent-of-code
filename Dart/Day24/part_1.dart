import 'dart:convert';
import 'dart:io';

class Hailstone {
  final double px, py, pz, vx, vy, vz;

  Hailstone(this.px, this.py, this.pz, this.vx, this.vy, this.vz);
}

class Intersection {
  final double x, y;

  Intersection(this.x, this.y);
}

Intersection? getLinesIntersection(double p1x, double v1x, double p1y,
    double v1y, double p2x, double v2x, double p2y, double v2y) {
  final x1 = p1x, x2 = p1x + 100000000000000 * v1x;
  final y1 = p1y, y2 = p1y + 100000000000000 * v1y;
  final x3 = p2x, x4 = p2x + 100000000000000 * v2x;
  final y3 = p2y, y4 = p2y + 100000000000000 * v2y;

  final denominator = (x1 - x2) * (y3 - y4) - (y1 - y2) * (x3 - x4);
  if (denominator == 0) {
    return null;
  }

  final x =
      ((x1 * y2 - y1 * x2) * (x3 - x4) - (x1 - x2) * (x3 * y4 - y3 * x4)) /
          denominator;
  final y =
      ((x1 * y2 - y1 * x2) * (y3 - y4) - (y1 - y2) * (x3 * y4 - y3 * x4)) /
          denominator;

  return Intersection(x, y);
}

double getTime(double s, double v, double p) {
  return (p - s) / v;
}

final List<Hailstone> hailstones = [];
const double min = 200000000000000;
const double max = 400000000000000;
int count = 0;

Future<void> interactionWithinTestArea() async {
  // Assuming input.txt is in the current working directory
  final file = await File('Dart/Day24/input.txt').openRead();
  final lines =
      await file.transform(utf8.decoder).transform(LineSplitter()).toList();

  for (final line in lines) {
    final parts = line.split(' @ ');
    final positions = parts[0];
    final velocity = parts[1];

    final pos = positions.split(', ').map((n) => double.parse(n)).toList();
    final px = pos[0], py = pos[1], pz = pos[2];

    final vel = velocity.split(', ').map((n) => double.parse(n)).toList();
    final vx = vel[0], vy = vel[1], vz = vel[2];

    hailstones.add(Hailstone(px, py, pz, vx, vy, vz));
  }

  for (var i = 0; i < hailstones.length; i++) {
    for (var j = i + 1; j < hailstones.length; j++) {
      final h1 = hailstones[i];
      final h2 = hailstones[j];

      final intersection = getLinesIntersection(
          h1.px, h1.vx, h1.py, h1.vy, h2.px, h2.vx, h2.py, h2.vy);

      if (intersection == null) {
        continue;
      }

      if (intersection.x < min ||
          intersection.x > max ||
          intersection.y < min ||
          intersection.y > max) {
        continue;
      }

      final timeH1 = getTime(h1.px, h1.vx, intersection.x);
      final timeH2 = getTime(h2.px, h2.vx, intersection.x);

      if (timeH1 < 0 || timeH2 < 0) {
        continue;
      }

      count++;
    }
  }

  print(count);
}
