import 'dart:io';

int regionsCanFitPresents() {
  final file = File('input.txt');
  if (!file.existsSync()) {
    throw Exception('File input.txt not found');
  }

  final lines = file.readAsLinesSync();
  final buffer = StringBuffer();
  var readingPatterns = true;
  final patterns = <int>[];
  final regionLines = <String>[];

  for (final line in lines) {
    if (line.isEmpty) {
      if (readingPatterns && buffer.isNotEmpty) {
        patterns.add(buffer.toString().split('#').length - 1);
        buffer.clear();
      }
      continue;
    }

    if (readingPatterns) {
      if (line.contains(': ')) {
        readingPatterns = false;
        if (buffer.isNotEmpty) {
          patterns.add(buffer.toString().split('#').length - 1);
          buffer.clear();
        }
        regionLines.add(line);
      } else {
        if (buffer.isNotEmpty) {
          buffer.write('\n');
        }
        buffer.write(line);
      }
    } else {
      regionLines.add(line);
    }
  }

  if (readingPatterns && buffer.isNotEmpty) {
    patterns.add(buffer.toString().split('#').length - 1);
  }

  var regions = 0;

  for (final line in regionLines) {
    final colonIndex = line.indexOf(': ');
    if (colonIndex == -1) {
      continue;
    }

    final areaStr = line.substring(0, colonIndex);
    final numsStr = line.substring(colonIndex + 2);

    final xIndex = areaStr.indexOf('x');
    if (xIndex == -1) {
      continue;
    }

    final width = int.tryParse(areaStr.substring(0, xIndex));
    final height = int.tryParse(areaStr.substring(xIndex + 1));
    if (width == null || height == null) {
      continue;
    }

    final area = width * height;

    final fields = numsStr.split(RegExp(r'\s+')).where((s) => s.isNotEmpty);
    var size = 0;

    var i = 0;
    for (final field in fields) {
      if (i >= patterns.length) {
        break;
      }
      final num = int.tryParse(field);
      if (num != null) {
        size += patterns[i] * num;
      }
      i++;
    }

    if (size > area) {
      continue;
    }

    if (size * 1.2 < area) {
      regions++;
    }
  }

  print('Regions: $regions');
  return regions;
}

void main() {
  regionsCanFitPresents();
}
