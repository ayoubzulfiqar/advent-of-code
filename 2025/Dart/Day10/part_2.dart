import 'dart:async';
import 'dart:io';
import 'dart:isolate';
import 'dart:math' as math;

class Puzzle {
  final String target;
  final List<int> joltage;
  final List<List<int>> buttons;

  Puzzle(this.target, this.joltage, this.buttons);
}

List<Puzzle> parseInput(String content) {
  final puzzles = <Puzzle>[];
  final lines = LineSplitter.split(
    content,
  ).map((e) => e.trim()).where((e) => e.isNotEmpty);

  for (final line in lines) {
    final parts = line
        .split(RegExp(r'\s+'))
        .where((p) => p.isNotEmpty)
        .toList();
    if (parts.length < 2) continue;

    // Parse target: [###..]
    final targetRaw = parts[0];
    if (targetRaw.length < 2 ||
        !targetRaw.startsWith('[') ||
        !targetRaw.endsWith(']'))
      continue;
    final target = targetRaw.substring(1, targetRaw.length - 1);

    // Parse joltage: {1,2,3}
    final joltageRaw = parts.last;
    if (joltageRaw.length < 2 ||
        !joltageRaw.startsWith('{') ||
        !joltageRaw.endsWith('}'))
      continue;
    final joltageStr = joltageRaw.substring(1, joltageRaw.length - 1);
    final joltage = <int>[];
    if (joltageStr.trim().isNotEmpty) {
      for (final s in joltageStr.split(',')) {
        final trimmed = s.trim();
        if (trimmed.isEmpty) continue;
        try {
          joltage.add(int.parse(trimmed));
        } catch (_) {
          // skip invalid
        }
      }
    }

    // Parse buttons: (0,2) (1,3)
    final buttons = <List<int>>[];
    for (int i = 1; i < parts.length - 1; i++) {
      final btn = parts[i];
      if (btn.length < 2 || !btn.startsWith('(') || !btn.endsWith(')'))
        continue;
      final inner = btn.substring(1, btn.length - 1).trim();
      final button = <int>[];
      if (inner.isNotEmpty) {
        for (final s in inner.split(',')) {
          final t = s.trim();
          if (t.isEmpty) continue;
          try {
            button.add(int.parse(t));
          } catch (_) {}
        }
      }
      buttons.add(button);
    }

    puzzles.add(Puzzle(target, joltage, buttons));
  }

  return puzzles;
}

class LineSplitter {
  static List<String> split(String text) {
    return text.split(RegExp(r'\r?\n'));
  }
}

// --- Gaussian Elimination (over integers) ---

Tuple2<List<int>, List<List<int>>> gaussianElimination(
  List<List<int>> originalMatrix,
) {
  if (originalMatrix.isEmpty) return Tuple2(<int>[], <List<int>>[]);

  final m = originalMatrix.length;
  final n = originalMatrix[0].length - 1; // last column is constant

  // Deep copy
  final mat = List<List<int>>.generate(
    m,
    (i) => List<int>.from(originalMatrix[i]),
  );

  final pivotCols = <int>[];
  var currentRow = 0;

  for (var col = 0; col < n && currentRow < m; col++) {
    // Find pivot
    int? pivotRow;
    for (var row = currentRow; row < m; row++) {
      if (mat[row][col] != 0) {
        pivotRow = row;
        break;
      }
    }

    if (pivotRow == null) continue;

    // Swap
    final temp = mat[currentRow];
    mat[currentRow] = mat[pivotRow];
    mat[pivotRow] = temp;
    pivotCols.add(col);

    // Eliminate below
    for (var row = currentRow + 1; row < m; row++) {
      if (mat[row][col] != 0) {
        final factor = mat[row][col];
        final pivotVal = mat[currentRow][col];
        for (var j = col; j <= n; j++) {
          // Cross-multiply to avoid fractions
          mat[row][j] = mat[row][j] * pivotVal - mat[currentRow][j] * factor;
        }
      }
    }

    currentRow++;
  }

  return Tuple2(pivotCols, mat);
}

List<int> solveSystem(List<List<int>> buttons, List<int> joltages) {
  final n = buttons.length; // number of buttons (variables)
  final m = joltages.length; // number of equations (lights)

  // Build augmented matrix: m rows, n+1 cols
  final matrix = List.generate(m, (i) {
    final row = List.filled(n + 1, 0);
    for (int j = 0; j < n; j++) {
      // Does button j affect light i?
      if (buttons[j].contains(i)) {
        row[j] = 1;
      }
    }
    row[n] = joltages[i]; // constant term
    return row;
  });

  final result = gaussianElimination(matrix);
  final pivotCols = result.item1;
  final reducedMatrix = result.item2;

  if (reducedMatrix.isEmpty) {
    return List.filled(n, 0);
  }

  final pivotSet = <int>{...pivotCols};
  final freeVars = List<int>.generate(
    n,
    (i) => i,
  ).where((i) => !pivotSet.contains(i)).toList();

  List<int> bestSolution = List.filled(n, 0);
  int bestSum = 1000000000; // or int.MaxValue equivalent

  void trySolution(List<int> freeValues) {
    final solution = List<int>.filled(n, 0);

    // Assign free variables
    for (int i = 0; i < freeVars.length; i++) {
      if (i < freeValues.length) {
        solution[freeVars[i]] = freeValues[i];
      }
    }

    // Back-substitute pivot variables (from bottom up)
    for (int idx = pivotCols.length - 1; idx >= 0; idx--) {
      final row = idx;
      final col = pivotCols[idx];
      var total = reducedMatrix[row][n]; // constant

      for (int j = col + 1; j < n; j++) {
        total -= reducedMatrix[row][j] * solution[j];
      }

      final coeff = reducedMatrix[row][col];
      if (coeff == 0) return; // inconsistent

      if (total % coeff != 0) return;
      final val = total ~/ coeff;
      if (val < 0) return;

      solution[col] = val;
    }

    for (int i = 0; i < m; i++) {
      int total = 0;
      for (int j = 0; j < n; j++) {
        if (solution[j] > 0 && buttons[j].contains(i)) {
          total += solution[j];
        }
      }
      if (total != joltages[i]) return;
    }

    final sum = solution.reduce((a, b) => a + b);
    if (sum < bestSum) {
      bestSum = sum;
      bestSolution = List.from(solution);
    }
  }

  // Try values for free variables
  if (freeVars.isEmpty) {
    trySolution([]);
  } else if (freeVars.length == 1) {
    final maxJ = joltages.isEmpty ? 0 : joltages.reduce(math.max);
    final maxVal = math.min(3 * maxJ, 1000);
    for (int v = 0; v <= maxVal; v++) {
      if (v > bestSum) break;
      trySolution([v]);
    }
  } else if (freeVars.length == 2) {
    final maxJ = joltages.isEmpty ? 0 : joltages.reduce(math.max);
    final maxVal = math.max(200, maxJ);
    for (int v1 = 0; v1 <= maxVal; v1++) {
      for (int v2 = 0; v2 <= maxVal; v2++) {
        if (v1 + v2 > bestSum) continue;
        trySolution([v1, v2]);
      }
    }
  } else if (freeVars.length == 3) {
    for (int v1 = 0; v1 < 250; v1++) {
      for (int v2 = 0; v2 < 250; v2++) {
        for (int v3 = 0; v3 < 250; v3++) {
          if (v1 + v2 + v3 > bestSum) continue;
          trySolution([v1, v2, v3]);
        }
      }
    }
  } else if (freeVars.length == 4) {
    for (int v1 = 0; v1 < 30; v1++) {
      for (int v2 = 0; v2 < 30; v2++) {
        for (int v3 = 0; v3 < 30; v3++) {
          for (int v4 = 0; v4 < 30; v4++) {
            if (v1 + v2 + v3 + v4 > bestSum) continue;
            trySolution([v1, v2, v3, v4]);
          }
        }
      }
    }
  } else {
    // Fallback: all zeros
    trySolution(List.filled(freeVars.length, 0));
  }

  return bestSolution;
}

// --- Parallel Isolates ---

class WorkItem {
  final Puzzle puzzle;
  WorkItem(this.puzzle);
}

class WorkResult {
  final int sum;
  WorkResult(this.sum);
}

int solvePuzzleLocally(Puzzle puzzle) {
  final solution = solveSystem(puzzle.buttons, puzzle.joltage);
  return solution.reduce((a, b) => a + b);
}

void _isolateEntry(dynamic args) {
  final List<dynamic> list = args;
  final List<Puzzle> puzzles = (list[0] as List)
      .map((e) => _puzzleFromJson(e as Map<String, dynamic>))
      .toList();
  final SendPort sendPort = list[1] as SendPort;

  int total = 0;
  for (final p in puzzles) {
    total += solvePuzzleLocally(p);
  }
  sendPort.send(total);
}

Map<String, dynamic> _puzzleToJson(Puzzle p) {
  return {'target': p.target, 'joltage': p.joltage, 'buttons': p.buttons};
}

Puzzle _puzzleFromJson(Map<String, dynamic> json) {
  return Puzzle(
    json['target'] as String,
    List<int>.from(json['joltage'] as List),
    (json['buttons'] as List).map((e) => List<int>.from(e as List)).toList(),
  );
}

Future<int> solvePuzzlesInParallel(List<Puzzle> puzzles) async {
  if (puzzles.isEmpty) return 0;

  final numIsolates = Platform.numberOfProcessors;
  final batchSize = (puzzles.length / numIsolates).ceil();
  final futures = <Future<int>>[];

  for (int i = 0; i < puzzles.length; i += batchSize) {
    final batch = puzzles.sublist(i, math.min(i + batchSize, puzzles.length));
    final jsonBatch = batch.map(_puzzleToJson).toList();
    final completer = Completer<int>();
    futures.add(completer.future);

    final rp = ReceivePort();
    await Isolate.spawn(_isolateEntry, [jsonBatch, rp.sendPort]);
    rp.listen((message) {
      rp.close();
      if (message is int) {
        completer.complete(message);
      } else {
        completer.completeError('Invalid message');
      }
    });
  }

  final results = await Future.wait(futures);
  return results.reduce((a, b) => a + b);
}

// --- Main ---

void main() async {
  final file = File('input.txt');
  if (!await file.exists()) {
    print('Error: input.txt not found.');
    exit(1);
  }

  final content = await file.readAsString();
  final puzzles = parseInput(content);

  if (puzzles.isEmpty) {
    print('Part-2: 0');
    return;
  }

  final total = await solvePuzzlesInParallel(puzzles);
  print('Part-2: $total');
}

class Tuple2<T1, T2> {
  final T1 item1;
  final T2 item2;
  const Tuple2(this.item1, this.item2);
}
