import 'dart:async';
import 'dart:io';
import 'dart:isolate';

class WorkItem {
  final String line;
  WorkItem(this.line);
}

class WorkResult {
  final int iterations;
  WorkResult(this.iterations);
}

// Convert list of indices to bitmask
int toBitmask(List<int> indices) {
  int mask = 0;
  for (final v in indices) {
    mask |= 1 << v;
  }
  return mask;
}

// Parse a single line (to be run in isolate)
WorkResult processLine(String line) {
  line = line.trim();
  if (line.isEmpty) return WorkResult(0);

  final parts = line.split(RegExp(r'\s+')).where((p) => p.isNotEmpty).toList();
  if (parts.length < 2) return WorkResult(0);

  final lightsStr = parts[0];
  // Extract wiring: skip first and last (assuming last is '()' or similar)
  final wiringParts = parts.sublist(1, parts.length - 1);

  // Parse lights: e.g., "[###..]" → indices of '#'
  final lightsClean = lightsStr.replaceAll(RegExp(r'[\[\]]'), '');
  final startIndices = <int>[];
  for (int i = 0; i < lightsClean.length; i++) {
    if (lightsClean[i] == '#') {
      startIndices.add(i);
    }
  }

  // Parse wiring buttons
  final buttonMasks = <int>[];
  for (final wire in wiringParts) {
    final cleanWire = wire.replaceAll(RegExp(r'[()]'), '').trim();
    if (cleanWire.isEmpty) continue;

    final nums = cleanWire.split(',');
    final buttonIndices = <int>[];
    for (final numStr in nums) {
      final trimmed = numStr.trim();
      if (trimmed.isEmpty) continue;
      try {
        final num = int.parse(trimmed);
        buttonIndices.add(num);
      } catch (e) {
        // Skip invalid numbers
        continue;
      }
    }
    if (buttonIndices.isNotEmpty) {
      buttonMasks.add(toBitmask(buttonIndices));
    }
  }

  final startMask = toBitmask(startIndices);
  final endMask = 0;

  if (startMask == endMask) return WorkResult(0);

  // BFS using Set<int> for visited states
  final current = <int>{startMask};
  int iterations = 0;
  const maxIterations = 100000;

  while (iterations <= maxIterations) {
    iterations++;
    final next = <int>{};

    for (final mask in current) {
      for (final button in buttonMasks) {
        final nextState = mask ^ button;
        if (nextState == endMask) {
          return WorkResult(iterations);
        }
        next.add(nextState);
      }
    }

    current
      ..clear()
      ..addAll(next);
  }

  // Safety break
  return WorkResult(maxIterations);
}

void main() async {
  final file = File('input.txt');
  if (!await file.exists()) {
    print('Error: input.txt not found.');
    return;
  }

  final content = await file.readAsString();
  final lines = LineSplitter.split(
    content,
  ).where((line) => line.trim().isNotEmpty).toList();

  if (lines.isEmpty) {
    print('Total: 0');
    return;
  }

  // Isolates parallel processing
  final total = await _processLinesInParallel(lines);
  print('Total: $total');
}

// Split lines robustly (handles \r\n, \n, etc.)
class LineSplitter {
  static List<String> split(String text) {
    return text.split(RegExp(r'\r?\n'));
  }
}

Future<int> _processLinesInParallel(List<String> lines) async {
  final int numIsolates = Platform.numberOfProcessors;
  final batchSize = (lines.length / numIsolates).ceil();

  final futures = <Future<int>>[];

  for (int i = 0; i < lines.length; i += batchSize) {
    final batch = lines.sublist(
      i,
      i + batchSize > lines.length ? lines.length : i + batchSize,
    );
    final completer = Completer<int>();
    futures.add(completer.future);

    final receivePort = ReceivePort();
    await Isolate.spawn(_isolateEntry, [batch, receivePort.sendPort]);

    receivePort.listen((message) {
      receivePort.close();
      if (message is int) {
        completer.complete(message);
      } else {
        completer.completeError('Unexpected isolate message');
      }
    });
  }

  final results = await Future.wait(futures);
  return results.reduce((a, b) => a + b);
}

// Isolate entry
void _isolateEntry(List<dynamic> args) {
  final List<String> lines = List<String>.from(args[0]);
  final SendPort sendPort = args[1] as SendPort;

  int total = 0;
  for (final line in lines) {
    total += processLine(line).iterations;
  }

  sendPort.send(total);
}
