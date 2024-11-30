import 'dart:io';

void desertIslandNumbers() {
  String input = File('Dart/Day6/input.txt').readAsStringSync();
  List<String> lines = input.split('\n');

  List<int> times = lines[0]
      .split(RegExp(r'\s+'))
      .skip(1)
      .where((s) => s.isNotEmpty)
      .map((s) => int.parse(s))
      .toList();

  List<int> records = lines[1]
      .split(RegExp(r'\s+'))
      .skip(1)
      .where((s) => s.isNotEmpty)
      .map((s) => int.parse(s))
      .toList();

  List<Tuple2<int, int>> races = List.generate(
      times.length, (index) => Tuple2(times[index], records[index]));

  int totalWays = 1;

  for (Tuple2<int, int> race in races) {
    int waysToWin = 0;

    for (int holdTime = 0; holdTime < race.item1; holdTime++) {
      int distance = (race.item1 - holdTime) * holdTime;

      if (distance > race.item2) {
        waysToWin += 1;
      }
    }

    totalWays *= waysToWin;
  }

  print('$totalWays');
}

class Tuple2<T1, T2> {
  final T1 item1;
  final T2 item2;

  Tuple2(this.item1, this.item2);
}
