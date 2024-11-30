import 'dart:io';

void main() {
  List<List<int>> histories = [];

  try {
    File file = File('Dart/Day9/input.txt');
    List<String> data = file.readAsLinesSync();

    for (String line in data) {
      List<int> values = parseIntegers(line);
      histories.add(values);
    }
  } catch (e) {
    print('error reading input.txt: $e');
    exit(1);
  }

  List<int> extrapolatedValues = [];

  for (List<int> history in histories) {
    List<List<int>> subLists = [history];
    bool allZeroes = false;

    while (!allZeroes) {
      List<int> sublist = [];

      for (int i = 0; i < subLists.last.length - 1; i++) {
        int difference = subLists.last[i + 1] - subLists.last[i];
        sublist.add(difference);
      }

      allZeroes = allZeroesInList(sublist);
      subLists.add(sublist);
    }

    subLists.last.insert(0, 0);

    for (int i = subLists.length - 2; i >= 0; i--) {
      int extrapolatedValue = subLists[i].first - subLists[i + 1].first;
      subLists[i].insert(0, extrapolatedValue);
    }

    extrapolatedValues.add(subLists.first.first);
  }

  int ans = sum(extrapolatedValues);
  print(ans);
}

List<int> parseIntegers(String input) {
  List<int> nums = [];
  List<String> fields = input.split(' ');

  for (String field in fields) {
    int num = int.tryParse(field) ?? 0;
    nums.add(num);
  }

  return nums;
}

bool allZeroesInList(List<int> nums) {
  for (int num in nums) {
    if (num != 0) {
      return false;
    }
  }
  return true;
}

int sum(List<int> nums) {
  int result = 0;

  for (int num in nums) {
    result += num;
  }

  return result;
}
