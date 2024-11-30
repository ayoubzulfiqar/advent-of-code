import 'dart:io';

int brickSafelyChosen() {
  File file = File('Dart/Day22/input.txt');
  List<List<int>> bricks = [];

  file.readAsLinesSync().forEach((line) {
    List<String> values = line.replaceAll('~', ',').split(',');
    List<int> brick = [];
    values.forEach((v) {
      int num = int.parse(v);
      brick.add(num);
    });
    bricks.add(brick);
  });

  bricks.sort((a, b) => a[2].compareTo(b[2]));

  bool overlaps(List<int> a, List<int> b) {
    return max(a[0], b[0]) <= min(a[3], b[3]) &&
        max(a[1], b[1]) <= min(a[4], b[4]);
  }

  for (int index = 0; index < bricks.length; index++) {
    int maxZ = 1;
    for (int i = 0; i < index; i++) {
      if (overlaps(bricks[index], bricks[i])) {
        maxZ = max(maxZ, bricks[i][5] + 1);
      }
    }
    bricks[index][5] -= bricks[index][2] - maxZ;
    bricks[index][2] = maxZ;
  }

  bricks.sort((a, b) => a[2].compareTo(b[2]));

  Map<int, Set<int>> kSupportsV = {};
  Map<int, Set<int>> vSupportsK = {};

  for (int i = 0; i < bricks.length; i++) {
    kSupportsV[i] = {};
    vSupportsK[i] = {};
  }

  for (int j = 0; j < bricks.length; j++) {
    for (int i = 0; i < j; i++) {
      if (overlaps(bricks[i], bricks[j]) && bricks[j][2] == bricks[i][5] + 1) {
        kSupportsV[i]!.add(j);
        vSupportsK[j]!.add(i);
      }
    }
  }

  int total = 0;

  for (int i = 0; i < bricks.length; i++) {
    bool satisfies = true;
    for (int j in kSupportsV[i]!) {
      if (vSupportsK[j]!.length < 2) {
        satisfies = false;
        break;
      }
    }
    if (satisfies) {
      total++;
    }
  }

  print(total);
  return total;
}

int max(int a, int b) => a > b ? a : b;

int min(int a, int b) => a < b ? a : b;
