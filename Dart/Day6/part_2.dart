import 'dart:io';
import 'dart:math';

void desertIslandLongerRace() {
  var lines = File('Dart/Day6/input.txt').readAsLinesSync();
  List<int> time = [
    int.parse(lines[0].split(":")[1].trim().replaceAll(" ", ""))
  ];
  List<int> distance = [
    int.parse(lines[1].split(":")[1].trim().replaceAll(" ", ""))
  ];

  int result = 1;

  for (int i = 0; i < time.length; i++) {
    int b = time[i];
    int c = distance[i];

    double delta = sqrt(b * b - 4 * c);

    int minR = (b - delta / 2 + 1).floor();
    int maxR = (b + delta / 2 - 1).ceil();
    int diff = maxR - minR + 1;

    result = result * diff;
  }

  print(result);
}
