import 'dart:io';

void engineGearRatioSum() async {
  try {
    final file = File("Dart/Day3/input.txt");
    final lines = await file.readAsLines();

    var total = 0;

    for (int r = 0; r < lines.length; r++) {
      final row = lines[r];
      for (int c = 0; c < row.length; c++) {
        final ch = row[c];

        if (ch != '*') {
          continue;
        }

        final cs = <String, bool>{};

        for (final cr in [r - 1, r, r + 1]) {
          for (int cc in [c - 1, c, c + 1]) {
            if (cr < 0 ||
                cr >= lines.length ||
                cc < 0 ||
                cc >= lines[cr].length ||
                lines[cr][cc].codeUnitAt(0) < '0'.codeUnitAt(0) ||
                lines[cr][cc].codeUnitAt(0) > '9'.codeUnitAt(0)) {
              continue;
            }

            while (cc > 0 &&
                lines[cr][cc - 1].codeUnitAt(0) >= '0'.codeUnitAt(0) &&
                lines[cr][cc - 1].codeUnitAt(0) <= '9'.codeUnitAt(0)) {
              cc--;
            }

            cs['($cr,$cc)'] = true;
          }
        }

        if (cs.length != 2) {
          continue;
        }

        final ns = <int>[];

        for (final key in cs.keys) {
          final match = RegExp(r'\((\d+),(\d+)\)').firstMatch(key);

          if (match != null) {
            final cr = int.parse(match.group(1)!);
            int cc = int.parse(match.group(2)!);
            var s = '';

            while (cc < lines[cr].length &&
                lines[cr][cc].codeUnitAt(0) >= '0'.codeUnitAt(0) &&
                lines[cr][cc].codeUnitAt(0) <= '9'.codeUnitAt(0)) {
              s += lines[cr][cc];
              cc++;
            }

            final n = int.tryParse(s);

            if (n != null) {
              ns.add(n);
            }
          }
        }

        total += ns[0] * ns[1];
      }
    }

    print(total);
  } catch (e) {
    print('Error: $e');
  }
}
