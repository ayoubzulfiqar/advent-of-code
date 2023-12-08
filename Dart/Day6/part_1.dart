import 'dart:io';

Future<int> desertIslandNumbers() async {
  final file = File('Dart/Day6/input.txt');
  try {
    final contents = await file.readAsString();
    final game = parseRecords(contents);
    final result = game.result();
    print(result);
    return result;
  } catch (e) {
    print(e);
    return 0;
  }
}

class Game {
  final List<Race> records;

  Game({required this.records});

  int result() {
    var wins = 1;
    for (final record in records) {
      wins *= record.findWinningTimes();
    }
    return wins;
  }
}

Game parseRecords(String input) {
  final lines = input.trim().split('\n');
  final times = lines[0].split(' ').skip(1).map(int.parse).toList();
  final distances = lines[1].split(' ').skip(1).map(int.parse).toList();

  final records = <Race>[];
  for (var i = 0; i < times.length; i++) {
    records.add(
      Race(
        times[i],
        distances[i],
      ),
    );
  }

  return Game(records: records);
}

class Race {
  late int time;
  late int distance;
  Race(int time, int distance) {
    this.time = time;
    this.distance = distance;
  }

  int findWinningTimes() {
    var wins = 0;
    for (var holdTime = 1; holdTime < time; holdTime++) {
      final distance = travelDistance(holdTime);
      if (distance > this.distance) {
        wins += 1;
      } else if (wins > 0) {
        break;
      }
    }
    return wins;
  }

  int travelDistance(int holdTime) => holdTime * (time - holdTime);
}
