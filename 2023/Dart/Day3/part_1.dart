import 'dart:io';

Future<int> engineSchematicSum() async {
  try {
    final file = File("Dart/Day3/input.txt");
    final lines = await file.readAsLines();

    final Set<String> cs = {};

    for (int r = 0; r < lines.length; r++) {
      final row = lines[r];
      for (int c = 0; c < row.length; c++) {
        final int ch = row[c].codeUnitAt(0);

        if (ch >= '0'.codeUnitAt(0) && ch <= '9'.codeUnitAt(0) ||
            ch == '.'.codeUnitAt(0)) {
          continue;
        }

        for (int dr = -1; dr <= 1; dr++) {
          for (int dc = -1; dc <= 1; dc++) {
            final nr = r + dr;
            int nc = c + dc;

            if (nr < 0 ||
                nr >= lines.length ||
                nc < 0 ||
                nc >= lines[nr].length ||
                lines[nr][nc].codeUnitAt(0) < '0'.codeUnitAt(0) ||
                lines[nr][nc].codeUnitAt(0) > '9'.codeUnitAt(0)) {
              continue;
            }

            while (nc > 0 &&
                lines[nr][nc - 1].codeUnitAt(0) >= '0'.codeUnitAt(0) &&
                lines[nr][nc - 1].codeUnitAt(0) <= '9'.codeUnitAt(0)) {
              nc--;
            }

            cs.add('($nr,$nc)');
          }
        }
      }
    }

    final List<int> ns = [];

    for (final key in cs) {
      final match = RegExp(r'\((\d+),(\d+)\)').firstMatch(key);
      if (match != null) {
        final r = int.parse(match.group(1)!);
        int c = int.parse(match.group(2)!);
        var s = '';

        while (c < lines[r].length &&
            lines[r][c].codeUnitAt(0) >= '0'.codeUnitAt(0) &&
            lines[r][c].codeUnitAt(0) <= '9'.codeUnitAt(0)) {
          s += lines[r][c];
          c++;
        }

        final n = int.tryParse(s);

        if (n != null) {
          ns.add(n);
        }
      }
    }

    final sum = ns.fold(0, (previous, number) => previous + number);
    print(sum);
    return sum;
  } catch (e) {
    print('Error: $e');
  }
  return 0;
}
