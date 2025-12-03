import 'dart:io';

int outputJoltage() {
  try {
    final file = File('input.txt');
    final lines = file.readAsLinesSync();

    int total = 0;

    for (final line in lines) {
      // Find length of numeric part (until non-digit)
      int length = 0;
      for (
        int i = 0;
        i < line.length && line.codeUnitAt(i) >= 48 && line.codeUnitAt(i) <= 57;
        i++
      ) {
        length++;
      }

      // Initialize lists
      final T = List<int>.filled(length, 0);
      final R = List<int>.filled(length, 0);

      // Fill T with digits
      for (int i = 0; i < length; i++) {
        T[i] = line.codeUnitAt(i) - 48; // '0' is 48 in ASCII
      }

      // Main logic
      const summands = 2;
      for (int it = 0; it < summands; it++) {
        int m = 0;
        for (int i = 0; i < length; i++) {
          final newVal = 10 * m + T[i];
          if (R[i] > m) {
            m = R[i];
          }
          R[i] = newVal;
        }
      }

      // Final pass
      int m = 0;
      for (int i = 0; i < length; i++) {
        if (R[i] > m) {
          m = R[i];
        }
        R[i] = m;
      }

      total += R[length - 1];
    }

    return total;
  } on FileSystemException catch (e) {
    print('Error opening file: $e');
    exit(1);
  } on Exception catch (e) {
    print('Error reading file: $e');
    exit(1);
  }
}

void main() {
  final result = outputJoltage();
  print(result);
}
