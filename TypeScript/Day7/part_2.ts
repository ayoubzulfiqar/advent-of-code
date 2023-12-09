import * as fs from 'fs';

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
  type: JokerHandType;
  cards: string[];
  bid: number;

  constructor(type: JokerHandType, cards: string[], bid: number) {
    this.type = type;
    this.cards = cards;
    this.bid = bid;
  }
}

class JokerGame extends Array<JokerHand> {
  sorting(): void {
    this.sort((a, b) => {
      if (a.type !== b.type) {
        return a.type - b.type;
      }

      const cardValues: Record<string, number> = {
        '2': 0, '3': 1, '4': 2, '5': 3, '6': 4,
        '7': 5, '8': 6, '9': 7, 'T': 8, 'J': -1,
        'Q': 10, 'K': 11, 'A': 12,
      };

      for (let k = 0; k < 5; k++) {
        if (cardValues[a.cards[k]] !== cardValues[b.cards[k]]) {
          return (cardValues[a.cards[k]] ?? 0) - (cardValues[b.cards[k]] ?? 0);
        }
      }

      return 0;
    });
  }
}

function getBestHand(cards: string[]): JokerHandType {
  const counts: Record<string, number> = {};

  for (const c of cards) {
    counts[c] = (counts[c] ?? 0) + 1;
  }

  if (counts['J'] !== 0) {
    let maxValue = 0;
    let maxCount = '';
    for (const c of Object.keys(counts)) {
      if (c !== 'J' && (counts[c] ?? 0) > maxValue) {
        maxCount = c;
        maxValue = counts[c] ?? 0;
      }
    }

    counts[maxCount] = (counts[maxCount] ?? 0) + (counts['J'] ?? 0);
    counts['J'] = 0;
  }

  const cardCounts: number[] = [];

  for (const v of Object.values(counts)) {
    if (v !== 0) {
      cardCounts.push(v);
    }
  }

  cardCounts.sort();

  let hand: JokerHandType;

  switch (JSON.stringify(cardCounts)) {
    case '[5]':
      hand = JokerHandType.JokerFiveOfKind;
      break;
    case '[1,4]':
      hand = JokerHandType.JokerFourOfAKind;
      break;
    case '[2,3]':
      hand = JokerHandType.JokerFullHouse;
      break;
    case '[1,1,3]':
      hand = JokerHandType.JokerThreeOfAKind;
      break;
    case '[1,2,2]':
      hand = JokerHandType.JokerTwoPair;
      break;
    case '[1,1,1,2]':
      hand = JokerHandType.JokerOnePair;
      break;
    case '[1,1,1,1,1]':
      hand = JokerHandType.JokerHighCard;
      break;
    default:
      throw new Error(`Unexpected card counts: ${JSON.stringify(cardCounts)}`);
  }

  return hand;
}

function jokerEqual(a: number[], b: number[]): boolean {
  if (a.length !== b.length) {
    return false;
  }

  for (let i = 0; i < a.length; i++) {
    if (a[i] !== b[i]) {
      return false;
    }
  }

  return true;
}

function NewJokerHand(ln: string): JokerHand {
  const parts = ln.split(' ');
  const cards = parts[0].split('');
  const bid = parseInt(parts[1], 10);

  const handType = getBestHand(cards);

  return new JokerHand(handType, cards, bid);
}

function jokerWining(): number {
  const lines: string[] = readLines();
  const g = new JokerGame();

  for (const ln of lines) {
    const hand = NewJokerHand(ln);
    g.push(hand);
  }

  g.sorting();

  let sum = 0;
  for (let i = 0; i < g.length; i++) {
    sum += (i + 1) * g[i].bid;
  }

  return sum;
}

function readLines(): string[] {
  const result: string[] = [];

  const file = fs.readFileSync('./input.txt', 'utf-8');
  try {
    const lines = file.split('\n');
    result.push(...lines);
  } catch (e) {
    console.log(e);
  }

  return result;
}

console.log(jokerWining());
