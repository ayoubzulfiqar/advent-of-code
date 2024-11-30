import 'dart:io';

void main() {
  final file = File("Dart/Day18/input.txt");
  List<String> L = file.readAsStringSync().trim().split('\n');

  final DIRS = 'RDLU';
  final Map<String, List<int>> D = {
    'U': [-1, 0],
    'D': [1, 0],
    'L': [0, -1],
    'R': [0, 1]
  };

  void part1(List<String> inputList) {
    List<int> pos = [0, 0];
    int count = 0;

    for (String l in inputList) {
      List<String> parts = l.split(' ');
      String dir_ = parts[0];
      int ln = int.parse(parts[1]);

      List<int> delta = D[dir_]!;
      count += ln;

      int dr = delta[0];
      int dc = delta[1];
      int rr = pos[0] + dr * ln;
      int cc = pos[1] + dc * ln;
      pos = [rr, cc];
    }

    print(count);
  }

  void part2(List<String> inputList) {
    List<int> pos = [0, 0];
    List<List<int>> V = [];
    int count = 0;

    for (String l in inputList) {
      List<String> parts = l.split(' ');
      String dir_ = parts[0];
      int ln = int.parse(parts[1]);
      String col = parts[2];

      col = col.substring(2, col.length - 1);
      String lnHex = col.substring(0, col.length - 1);
      String d = col[col.length - 1];
      ln = int.parse(lnHex, radix: 16);
      dir_ = DIRS[int.parse(d)];

      List<int> delta = D[dir_]!;
      count += ln;

      int dr = delta[0];
      int dc = delta[1];
      int rr = pos[0] + dr * ln;
      int cc = pos[1] + dc * ln;
      pos = [rr, cc];
      V.add(List.from(pos));
    }

    int S = 0;
    for (int v = 0; v < V.length; v++) {
      int bef = (v - 1 + V.length) % V.length;
      int af = (v + 1) % V.length;
      S += V[v][1] * (V[bef][0] - V[af][0]);
    }

    S ~/= 2;
    S = S.abs();
    int tot = (S + 1) - (count ~/ 2);
    print(tot + count);
  }

  part1(L);
  part2(L);
}
