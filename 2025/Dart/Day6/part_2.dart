import 'dart:io';

int individualGrandAnswersTotals() {
  try {
    final file = File('input.txt');
    final contentString = file.readAsStringSync();

    var allLines = contentString.split('\n');

    if (allLines.isNotEmpty && allLines.last.isEmpty) {
      allLines = allLines.sublist(0, allLines.length - 1);
    }

    if (allLines.isEmpty) {
      return 0;
    }

    final bottomOperationLine = allLines.last;
    final operators = <String>[];
    final operatorColumnStarts = <int>[];

    for (
      var columnIndex = 0;
      columnIndex < bottomOperationLine.length;
      columnIndex++
    ) {
      if (bottomOperationLine[columnIndex] != ' ') {
        operators.add(bottomOperationLine[columnIndex]);
        operatorColumnStarts.add(columnIndex);
      }
    }
    operatorColumnStarts.add(bottomOperationLine.length);

    final dataRowsCount = allLines.length - 1;
    if (dataRowsCount <= 0) {
      return 0;
    }

    var totalAnswerSum = 0;

    for (var operatorIdx = 0; operatorIdx < operators.length; operatorIdx++) {
      final currentOperator = operators[operatorIdx];
      var verticalColumnResult = 0;

      if (currentOperator == '*') {
        verticalColumnResult = 1;
      }

      var groupStartColumn = operatorColumnStarts[operatorIdx];
      var groupEndColumn = operatorColumnStarts[operatorIdx + 1];

      if (operatorIdx < operators.length - 1) {
        groupEndColumn -= 1;
      }

      for (
        var currentColumn = groupStartColumn;
        currentColumn < groupEndColumn;
        currentColumn++
      ) {
        var verticalNumber = 0;

        for (var rowIndex = 0; rowIndex < dataRowsCount; rowIndex++) {
          final character = allLines[rowIndex][currentColumn];
          if (character != ' ') {
            verticalNumber = verticalNumber * 10 + int.parse(character);
          }
        }

        if (currentOperator == '*') {
          verticalColumnResult *= verticalNumber;
        } else {
          verticalColumnResult += verticalNumber;
        }
      }

      totalAnswerSum += verticalColumnResult;
    }

    return totalAnswerSum;
  } catch (e) {
    print('Error reading input.txt: $e');
    exit(1);
  }
}

void main() {
  final result = individualGrandAnswersTotals();
  print('Part 2: $result');
}
