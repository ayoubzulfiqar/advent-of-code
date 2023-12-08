import 'dart:collection';
import 'dart:io';
import 'dart:convert';

void main() {
  List<String> readLines() {
    List<String> result = [];

    File file = File('Dart/Day7/input.txt');
    try {
      List<String> lines = file.readAsLinesSync(encoding: utf8);
      result.addAll(lines);
    } catch (e) {
      print(e);
    }

    return result;
  }

  int jokerWining() {
    List<String> lines = readLines();

    JokerGame g = JokerGame();
    for (String ln in lines) {
      JokerHand hand = NewJokerHand(ln);
      g.add(hand);
    }

    g.sorting();

    int sum = 0;
    for (int i = 0; i < g.length; i++) {
      sum += (i + 1) * g[i].bid;
    }

    return sum;
  }

  print(jokerWining());
}

enum JokerHandType {
  JokerHighCard,
  JokerOnePair,
  JokerTwoPair,
  JokerThreeOfAKind,
  JokerFullHouse,
  JokerFourOfAKind,
  JokerFiveOfKind,
}

class JokerHand {
  JokerHandType type;
  List<String> cards;
  int bid;

  JokerHand(this.type, this.cards, this.bid);
}

class JokerGame extends ListBase<JokerHand> {
  List<JokerHand> _hands = [];

  @override
  int get length => _hands.length;

  @override
  set length(int newLength) {
    _hands.length = newLength;
  }

  @override
  JokerHand operator [](int index) => _hands[index];

  @override
  void operator []=(int index, JokerHand value) {
    _hands[index] = value;
  }

  void add(JokerHand hand) {
    _hands.add(hand);
  }

  void sorting() {
    _hands.sort((a, b) {
      if (a.type != b.type) {
        return a.type.index - b.type.index;
      }

      Map<String, int> cardValues = {
        '2': 0,
        '3': 1,
        '4': 2,
        '5': 3,
        '6': 4,
        '7': 5,
        '8': 6,
        '9': 7,
        'T': 8,
        'J': -1,
        'Q': 10,
        'K': 11,
        'A': 12,
      };

      for (int k = 0; k < 5; k++) {
        if (cardValues[a.cards[k]] != cardValues[b.cards[k]]) {
          return (cardValues[a.cards[k]] ?? 0) - (cardValues[b.cards[k]] ?? 0);
        }
      }

      return 0;
    });
  }
}

JokerHandType getBestHand(List<String> cards) {
  Map<String, int> counts = {};

  for (String c in cards) {
    counts[c] = (counts[c] ?? 0) + 1;
  }

  if (counts['J'] != 0) {
    int maxValue = 0;
    String maxCount = '';
    for (String c in counts.keys) {
      if (c != 'J' && (counts[c] ?? 0) > maxValue) {
        maxCount = c;
        maxValue = (counts[c] ?? 0);
      }
    }

    counts[maxCount] = (counts[maxCount] ?? 0) + (counts['J'] ?? 0);
    counts['J'] = 0;
  }

  List<int> cardCounts = [];

  for (int v in counts.values) {
    if (v != 0) {
      cardCounts.add(v);
    }
  }

  cardCounts.sort();

  JokerHandType hand;
  switch (cardCounts) {
    case [5]:
      hand = JokerHandType.JokerFiveOfKind;
      break;
    case [1, 4]:
      hand = JokerHandType.JokerFourOfAKind;
      break;
    case [2, 3]:
      hand = JokerHandType.JokerFullHouse;
      break;
    case [1, 1, 3]:
      hand = JokerHandType.JokerThreeOfAKind;
      break;
    case [1, 2, 2]:
      hand = JokerHandType.JokerTwoPair;
      break;
    case [1, 1, 1, 2]:
      hand = JokerHandType.JokerOnePair;
      break;
    case [1, 1, 1, 1, 1]:
      hand = JokerHandType.JokerHighCard;
      break;
    default:
      throw FormatException('Unexpected card counts: $cardCounts');
  }

  return hand;
}

bool jokerEqual(List<int> a, List<int> b) {
  if (a.length != b.length) {
    return false;
  }

  for (int i = 0; i < a.length; i++) {
    if (a[i] != b[i]) {
      return false;
    }
  }

  return true;
}

JokerHand NewJokerHand(String ln) {
  List<String> parts = ln.split(' ');

  List<String> cards = parts[0].split('').toList();
  int bid = int.parse(parts[1]);

  JokerHandType handType = getBestHand(cards);

  return JokerHand(handType, cards, bid);
}
