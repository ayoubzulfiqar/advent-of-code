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

  List<String> lines = readLines();
  Game g = Game();

  for (String ln in lines) {
    Hand hand = NewHand(ln);
    g.add(hand);
  }

  g.sorting();

  int sum = 0;
  for (int i = 0; i < g.length; i++) {
    sum += (i + 1) * g[i].bid;
  }

  print(sum);
}

enum HandType {
  HighCard,
  OnePair,
  TwoPair,
  ThreeOfAKind,
  FullHouse,
  FourOfAKind,
  FiveOfAKind,
}

class Hand {
  HandType type;
  List<String> cards;
  int bid;

  Hand(this.type, this.cards, this.bid);
}

class Game extends ListBase<Hand> {
  List<Hand> _hands = [];

  @override
  int get length => _hands.length;

  @override
  set length(int newLength) {
    _hands.length = newLength;
  }

  @override
  Hand operator [](int index) => _hands[index];

  @override
  void operator []=(int index, Hand value) {
    _hands[index] = value;
  }

  void add(Hand hand) {
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
        'J': 9,
        'Q': 10,
        'K': 11,
        'A': 12,
      };

      for (int k = 0; k < 5; k++) {
        if (cardValues[a.cards[k]] != cardValues[b.cards[k]]) {
          return cardValues[a.cards[k]]! - cardValues[b.cards[k]]!;
        }
      }

      return 0;
    });
  }
}

HandType extractInfo(List<String> a) {
  int mmax = 0;
  Map<String, int> c = {};
  for (String k in a) {
    c[k] = (c[k] ?? 0) + 1;
    mmax = (mmax > c[k]! ? mmax : c[k])!;
  }
  return HandType.values[mmax];
}

HandType calculateHandType(List<String> cards) {
  List<int> cardCounts = [];
  Map<String, int> counts = {};

  for (String c in cards) {
    counts[c] = (counts[c] ?? 0) + 1;
  }

  for (int v in counts.values) {
    cardCounts.add(v);
  }

  cardCounts.sort();

  HandType hand;
  switch (cardCounts) {
    case [5]:
      hand = HandType.FiveOfAKind;
      break;
    case [1, 4]:
      hand = HandType.FourOfAKind;
      break;
    case [2, 3]:
      hand = HandType.FullHouse;
      break;
    case [1, 1, 3]:
      hand = HandType.ThreeOfAKind;
      break;
    case [1, 2, 2]:
      hand = HandType.TwoPair;
      break;
    case [1, 1, 1, 2]:
      hand = HandType.OnePair;
      break;
    case [1, 1, 1, 1, 1]:
      hand = HandType.HighCard;
      break;
    default:
      throw FormatException('Unexpected card counts: $cardCounts');
  }

  return hand;
}

Hand NewHand(String ln) {
  List<String> parts = ln.split(' ');

  List<String> cards = parts[0].split('').toList();
  int bid = int.parse(parts[1]);

  HandType handType = calculateHandType(cards);

  return Hand(handType, cards, bid);
}
