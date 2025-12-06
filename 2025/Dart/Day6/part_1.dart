import 'dart:io';

class Data {
  List<String> operators = [];
  List<int> positions = [];
  List<String> lines = [];
}

int individualGrandTotals() {
  try {
    final file = File('input.txt');
    final content = file.readAsStringSync();

    var lines = content.split('\n');

    if (lines.isNotEmpty && lines.last.isEmpty) {
      lines = lines.sublist(0, lines.length - 1);
    }

    if (lines.isEmpty) {
      return 0;
    }

    final operationLine = lines.last;
    final operators = <String>[];
    final operatorColumns = <int>[];

    for (var column = 0; column < operationLine.length; column++) {
      if (operationLine[column] != ' ') {
        operators.add(operationLine[column]);
        operatorColumns.add(column);
      }
    }
    operatorColumns.add(operationLine.length);

    final dataRowCount = lines.length - 1;
    if (dataRowCount <= 0) {
      return 0;
    }

    var grandTotal = 0;

    for (
      var operatorIndex = 0;
      operatorIndex < operators.length;
      operatorIndex++
    ) {
      final currentOperator = operators[operatorIndex];
      var columnResult = 0;

      if (currentOperator == '*') {
        columnResult = 1;
      } else {
        columnResult = 0;
      }

      var columnStart = operatorColumns[operatorIndex];
      var columnEnd = operatorColumns[operatorIndex + 1];

      if (operatorIndex < operators.length - 1) {
        columnEnd -= 1;
      }

      for (var row = 0; row < dataRowCount; row++) {
        var parsedNumber = 0;

        for (var col = columnStart; col < columnEnd; col++) {
          final digitChar = lines[row][col];
          if (digitChar != ' ') {
            parsedNumber = parsedNumber * 10 + int.parse(digitChar);
          }
        }

        if (currentOperator == '*') {
          columnResult *= parsedNumber;
        } else {
          columnResult += parsedNumber;
        }
      }

      grandTotal += columnResult;
    }

    return grandTotal;
  } catch (e) {
    print('Error reading input.txt: $e');
    exit(1);
  }
}

void main() {
  final result = individualGrandTotals();
  print('Part 1: $result');
}
