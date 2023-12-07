import { StringArrayInputParts, Day } from '@advent/utils'

type Game = Map<number, number[]>

export const day4_2023 = ({ part1, part2 }: StringArrayInputParts): Day => ({
  part1: (): number =>
    [...createGame(part1).values()].reduce((prev, curr) => prev + curr.reduce((p) => (p == 0 ? 1 : p * 2), 0), 0),
  part2: (): number => {
    const map: Game = createGame(part2)
    return [...map.keys()].reduce((prev, curr) => {
      return prev + countCopies(map, curr, map.get(curr)!.length)
    }, map.size)
  },
})

const createGame = (input: string[], game: Game = new Map()): Game => {
  input.forEach((line) => {
    const colonSplit = line.split(':')
    const splitted = colonSplit[1].split('|')
    const winning = splitted[0]
      .trim()
      .split(/\s+/)
      .map((x) => parseInt(x))
    const numbers = splitted[1]
      .trim()
      .split(/\s+/)
      .map((x) => parseInt(x))
    game.set(
      parseInt(colonSplit[0].split(/\s+/)[1]),
      numbers.filter((n) => winning.includes(n))
    )
  })
  return game
}

const countCopies = (map: Game, cardNumber: number, total = 0): number => {
  const copies = map.get(cardNumber)?.length

  if (!copies) {
    return total
  }

  for (let i = cardNumber + 1; i <= cardNumber + copies; i++) {
    total += countCopies(map, i, map.get(i)!.length)
  }

  return total
}
