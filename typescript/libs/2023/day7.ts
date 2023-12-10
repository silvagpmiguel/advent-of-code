import { StringArrayInputParts, Day, countOccurrences } from '@advent/utils'

type Card = 'A' | 'K' | 'Q' | 'J' | 'T' | '9' | '8' | '7' | '6' | '5' | '4' | '3' | '2'
type RankStats = { rank: number; handStr: string; hand: Card[]; bid: number }
type Ranks = RankStats[]

export const day7_2023 = ({ part1, part2 }: StringArrayInputParts): Day => ({
  part1: (): number => {
    const cardScore: Map<Card, number> = new Map([
      ['A', 13],
      ['K', 12],
      ['Q', 11],
      ['J', 10],
      ['T', 9],
      ['9', 8],
      ['8', 7],
      ['7', 6],
      ['6', 5],
      ['5', 4],
      ['4', 3],
      ['3', 2],
      ['2', 1],
    ])
    return createRanks(part1)
      .sort((a, b) => handSort(cardScore, a, b))
      .reduce((prev, stats, index) => prev + (index + 1) * stats.bid, 0)
  },
  part2: (): number => {
    const cardScore: Map<Card, number> = new Map([
      ['A', 12],
      ['K', 11],
      ['Q', 10],
      ['T', 9],
      ['9', 8],
      ['8', 7],
      ['7', 6],
      ['6', 5],
      ['5', 4],
      ['4', 3],
      ['3', 2],
      ['2', 1],
      ['J', 0],
    ])
    return createRanks(part2, true)
      .sort((a, b) => handSort(cardScore, a, b))
      .reduce((prev, stats, index) => prev + (index + 1) * stats.bid, 0)
  },
})

const createRanks = (input: string[], joker = false): Ranks =>
  input.reduce((prev: Ranks, line: string) => {
    let [handStr, bidStr] = line.split(' ')
    const hand = handStr.split('') as Card[]
    const occurrences = joker ? getJokerCardOcurrences(handStr, hand) : getCardOcurrences(handStr, hand)
    return prev.concat({
      rank: computeRank(occurrences),
      hand,
      handStr,
      bid: parseInt(bidStr),
    } as RankStats)
  }, [])
const getJokerCardOcurrences = (handStr: string, hand: Card[]): Map<string, number> => {
  const occurrences = new Map()
  hand.forEach((card) => occurrences.set(card, countOccurrences(handStr, card)))
  if (occurrences.has('J')) {
    const jokerCount = occurrences.get('J')
    const highestOccurrence = [...occurrences.keys()]
      .filter((k) => k != 'J')
      .sort((a, b) => occurrences.get(b) - occurrences.get(a))[0]
    occurrences.delete('J')
    occurrences.set(highestOccurrence, occurrences.get(highestOccurrence) + jokerCount)
  }
  return occurrences
}
const getCardOcurrences = (handStr: string, hand: Card[]): Map<string, number> => {
  const map = new Map()
  hand.forEach((card) => map.set(card, countOccurrences(handStr, card)))
  return map
}
const computeRank = (occurrences: Map<string, number>): number => {
  if (occurrences.size == 1) {
    return 7
  } else if (occurrences.size == 2) {
    return [...occurrences.values()].some((v) => v == 4) ? 6 : 5
  } else if (occurrences.size == 3) {
    return [...occurrences.values()].some((v) => v == 3) ? 4 : 3
  } else if (occurrences.size == 4) {
    return 2
  } else if (occurrences.size == 5) {
    return 1
  } else {
    throw new Error('computeRank: wrong input?')
  }
}
const handSort = (cardScore: Map<Card, number>, stats1: RankStats, stats2: RankStats): number => {
  if (stats1.rank != stats2.rank) {
    return stats1.rank - stats2.rank
  }
  for (let i = 0; i < stats1.hand.length; i++) {
    const hand1Score = cardScore.get(stats1.hand[i])!
    const hand2Score = cardScore.get(stats2.hand[i])!
    if (hand1Score > hand2Score) {
      return 1
    } else if (hand1Score < hand2Score) {
      return -1
    }
  }
  throw new Error('handSort: wrong input?')
}

/**
 * NOT NEEDED, BUT WOULD BE INTERESTING TO SORT FIRST BY OCCURENCE AND THEN FOR SCORE TO REALLY CHECK THE TOTAL WINNINGS :).
 */
const cardSort = (cardScore: Map<Card, number>, cardOccurrences: Map<string, number>, a: Card, b: Card): number => {
  const occurenceA = cardOccurrences.get(a)!
  const occurenceB = cardOccurrences.get(b)!
  const scoreA = cardScore.get(a)!
  const scoreB = cardScore.get(b)!
  if (occurenceA > occurenceB) {
    return -1
  } else if (occurenceA < occurenceB) {
    return 1
  } else {
    return scoreB - scoreA
  }
}
