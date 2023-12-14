import 'dart:io';

List<List<String>> _map = [];

int TotalLoadOnNorthSupport() {
  final data = File("Dart/Day14/input.txt").readAsLinesSync();
  _map = List<List<String>>.from(data.map((line) => line.split('')));
  tilt();
  final result = countTotalLoad();
  print(result);
  return result;
}

void tilt() {
  for (var i = 1; i < _map.length; i++) {
    for (var x = 0; x < _map[i].length; x++) {
      if (_map[i][x] == 'O') {
        final col = List<String>.generate(_map.length, (y) => _map[y][x]);
        var prevY = i;
        for (var y = i - 1; y >= 0; y--) {
          if (col[y] == '.') {
            _map[y][x] = 'O';
            _map[prevY][x] = '.';
            prevY = y;
          } else if (col[y] == '#') {
            break;
          }
        }
      }
    }
  }
}

List<List<String>> turn() {
  _map = List<List<String>>.generate(
    _map[0].length,
    (i) => List<String>.from(
      _map.map((row) => row[i]).toList().reversed,
    ),
  );
  return _map;
}

int countTotalLoad() {
  final height = _map.length;
  return List.generate(
    height,
    (i) => (height - i) * _map[i].where((c) => c == 'O').length,
  ).fold(0, (acc, value) => acc + value);
}
