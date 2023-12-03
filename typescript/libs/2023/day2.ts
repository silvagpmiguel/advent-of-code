import { Day, StringArrayInputParts } from '@advent/utils'

export type Color = 'red' | 'green' | 'blue'
export type Score = Record<Color, number>
export type ScoreSet = Partial<Score>
export type Game = Map<number, ScoreSet[]>

export const day2_2023 = ({ part1, part2 }: StringArrayInputParts, score: Score): Day => ({
  part1: (): number => getGameIdsWithLessOrEqualScore(createGame(part1), score).reduce((prev, curr) => prev + curr, 0),
  part2: (): number => getMinimumScoreForEachColor(createGame(part2)).reduce((prev, curr) => prev + curr, 0),
})

const createGame = (input: string[], game = new Map()): Game => {
  input.forEach((line) => {
    const splitted = line.split(':')
    const id = parseInt(splitted[0].split(' ')[1])
    const input = splitted[1].split(';')
    game.set(id, createScoreSet(input))
  })
  return game
}

const createScoreSet = (input: string[], scoreSet: ScoreSet[] = []): ScoreSet[] => {
  input.forEach((setInput) => {
    setInput.split(',').forEach((score) => {
      const splitted = score.trim().split(/\s+/)
      const value = parseInt(splitted[0].trim())
      const key = splitted[1].trim() as Color
      scoreSet.push({ [key]: value })
    })
  })
  return scoreSet
}

const getMinimumScoreForEachColor = (game: Game): number[] => {
  return [...game.keys()].map((key) => {
    let minScore: Score = { red: 0, green: 0, blue: 0 }
    game.get(key)?.forEach((scoreSet) => {
      const red = scoreSet.red ?? 0
      const green = scoreSet.green ?? 0
      const blue = scoreSet.blue ?? 0
      minScore.red = red > minScore.red ? red : minScore.red
      minScore.green = green > minScore.green ? green : minScore.green
      minScore.blue = blue > minScore.blue ? blue : minScore.blue
    })
    return minScore.blue * minScore.green * minScore.red
  })
}

const getGameIdsWithLessOrEqualScore = (game: Game, score: Score): number[] =>
  [...game.keys()].filter((key) => lessOrEqual(score, game.get(key)))

const lessOrEqual = (score: Score, scoreSet: ScoreSet[] = []): boolean => {
  return scoreSet.every(
    ({ blue, red, green }) => (blue ?? 0) <= score.blue && (red ?? 0) <= score.red && (green ?? 0) <= score.green
  )
}
