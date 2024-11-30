import 'dart:io';

int count(String cfg, List<int> nums) {
  if (cfg.isEmpty) {
    if (nums.isEmpty) {
      return 1;
    }
    return 0;
  }

  if (nums.isEmpty) {
    if (cfg.contains("#")) {
      return 0;
    }
    return 1;
  }

  var result = 0;

  if (cfg[0] == '.' || cfg[0] == '?') {
    result += count(cfg.substring(1), nums);
  }

  if (cfg[0] == '#' || cfg[0] == '?') {
    if (nums[0] <= cfg.length &&
        !cfg.substring(0, nums[0]).contains(".") &&
        (nums[0] == cfg.length || cfg[nums[0]] != '#')) {
      if (nums[0] == cfg.length) {
        result += count("", nums.sublist(1));
      } else {
        result += count(cfg.substring(nums[0] + 1), nums.sublist(1));
      }
    }
  }

  return result;
}

void sumOfBrokenSprings() {
  var file = File("Dart/Day12/input.txt");
  var total = 0;

  var lines = file.readAsLinesSync();
  for (var line in lines) {
    var parts = line.split(' ');
    var cfg = parts[0];
    var numsStr = parts[1].split(',');
    var nums = numsStr.map((numStr) => int.parse(numStr)).toList();
    total += count(cfg, nums);
  }

  print(total);
}

void main() {
  sumOfBrokenSprings();
}
