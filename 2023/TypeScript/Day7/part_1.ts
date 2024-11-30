import * as fs from 'fs';

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
  type: HandType;
  cards: string[];
  bid: number;

  constructor(type: HandType, cards: string[], bid: number) {
    this.type = type;
    this.cards = cards;
    this.bid = bid;
  }
}

class Game extends Array<Hand> {
  sorting(): void {
    this.sort((a, b) => {
      if (a.type !== b.type) {
        return a.type - b.type;
      }

      const cardValues: Record<string, number> = {
        '2': 0, '3': 1, '4': 2, '5': 3, '6': 4,
        '7': 5, '8': 6, '9': 7, 'T': 8, 'J': 9,
        'Q': 10, 'K': 11, 'A': 12,
      };

      for (let k = 0; k < 5; k++) {
        if (cardValues[a.cards[k]] !== cardValues[b.cards[k]]) {
          return cardValues[a.cards[k]]! - cardValues[b.cards[k]]!;
        }
      }

      return 0;
    });
  }
}

function extractInfo(a: string[]): HandType {
  let mmax = 0;
  const c: Record<string, number> = {};

  for (const k of a) {
    c[k] = (c[k] ?? 0) + 1;
    mmax = Math.max(mmax, c[k]!);
  }

  return mmax as HandType;
}

function calculateHandType(cards: string[]): HandType {
  const cardCounts: number[] = [];
  const counts: Record<string, number> = {};

  for (const c of cards) {
    counts[c] = (counts[c] ?? 0) + 1;
  }

  for (const v of Object.values(counts)) {
    cardCounts.push(v);
  }

  cardCounts.sort();

  let hand: HandType;

  switch (JSON.stringify(cardCounts)) {
    case '[5]':
      hand = HandType.FiveOfAKind;
      break;
    case '[1,4]':
      hand = HandType.FourOfAKind;
      break;
    case '[2,3]':
      hand = HandType.FullHouse;
      break;
    case '[1,1,3]':
      hand = HandType.ThreeOfAKind;
      break;
    case '[1,2,2]':
      hand = HandType.TwoPair;
      break;
    case '[1,1,1,2]':
      hand = HandType.OnePair;
      break;
    case '[1,1,1,1,1]':
      hand = HandType.HighCard;
      break;
    default:
      throw new Error(`Unexpected card counts: ${JSON.stringify(cardCounts)}`);
  }

  return hand;
}

function NewHand(ln: string): Hand {
  const parts = ln.split(' ');
  const cards = parts[0].split('');
  const bid = parseInt(parts[1], 10);

  const handType = calculateHandType(cards);

  return new Hand(handType, cards, bid);
}

function cardGame(): void {
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

  const lines = readLines();
  const g = new Game();

  for (const ln of lines) {
    const hand = NewHand(ln);
    g.push(hand);
  }

  g.sorting();

  let sum = 0;
  for (let i = 0; i < g.length; i++) {
    sum += (i + 1) * g[i].bid;
  }

  console.log(sum);
}

cardGame();
