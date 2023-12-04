import { StringArrayInputParts, Day } from '@advent/utils'

export const day3_2023 = ({ part1, part2 }: StringArrayInputParts): Day => ({
  part1: (): number => {
    const row_len = part1.length
    const col_len = part1[0].length
    let partNumber = []
    let adjacentSymbols = false
    let count = 0
    for (let i = 0; i < row_len; i++) {
      for (let j = 0; j < col_len; j++) {
        const currentChar = part1[i][j]
        if (isNumber(currentChar)) {
          partNumber.push(parseInt(currentChar))
          adjacentSymbols = adjacentSymbols || hasAdjacentSymbols(part1, i, j, row_len, col_len)
        } else if (adjacentSymbols) {
          count += parseInt(partNumber.join(''))
          partNumber = []
          adjacentSymbols = false
        } else {
          partNumber = []
          adjacentSymbols = false
        }
      }
    }
    return count
  },
  part2: (): number => {
    const row_len = part2.length
    const col_len = part2[0].length
    let partNumber = []
    let gears = new Map()
    let adjacentSymbolPos: string | undefined
    for (let i = 0; i < row_len; i++) {
      for (let j = 0; j < col_len; j++) {
        const currentChar = part2[i][j]
        if (isNumber(currentChar)) {
          partNumber.push(parseInt(currentChar))
          adjacentSymbolPos = adjacentSymbolPos || getAdjacentSymbolPos(part2, i, j, row_len, col_len)
          if (adjacentSymbolPos && !gears.has(adjacentSymbolPos)) {
            gears.set(adjacentSymbolPos, [])
          }
        } else if (adjacentSymbolPos) {
          gears.get(adjacentSymbolPos).push(parseInt(partNumber.join('')))
          partNumber = []
          adjacentSymbolPos = undefined
        } else {
          partNumber = []
          adjacentSymbolPos = undefined
        }
      }
    }
    return [...gears.values()]
      .filter((parts) => parts.length > 1)
      .reduce((prev: number, curr: number[]) => prev + curr.reduce((p, c) => p * c), 0)
  },
})

const hasAdjacentSymbols = (
  input: string[],
  i: number,
  j: number,
  row_len: number,
  col_len: number,
  is = isSymbol
): boolean => {
  return (
    (i + 1 < row_len && is(input[i + 1][j])) ||
    (i - 1 >= 0 && is(input[i - 1][j])) ||
    (j + 1 < col_len && is(input[i][j + 1])) ||
    (j - 1 >= 0 && is(input[i][j - 1])) ||
    (i + 1 < row_len && j + 1 < col_len && is(input[i + 1][j + 1])) ||
    (i + 1 < row_len && j - 1 >= 0 && is(input[i + 1][j - 1])) ||
    (i - 1 >= 0 && j + 1 < col_len && is(input[i - 1][j + 1])) ||
    (i - 1 >= 0 && j - 1 >= 0 && is(input[i - 1][j - 1]))
  )
}
const getAdjacentSymbolPos = (
  input: string[],
  i: number,
  j: number,
  row_len: number,
  col_len: number,
  is = isAsterisk
): string | undefined => {
  if (i + 1 < row_len && is(input[i + 1][j])) return `${i + 1}, ${j}`
  else if (i - 1 >= 0 && is(input[i - 1][j])) return `${i - 1}, ${j}`
  else if (j + 1 < col_len && is(input[i][j + 1])) return `${i}, ${j + 1}`
  else if (j - 1 >= 0 && is(input[i][j - 1])) return `${i}, ${j - 1}`
  else if (i + 1 < row_len && j + 1 < col_len && is(input[i + 1][j + 1])) return `${i + 1}, ${j + 1}`
  else if (i + 1 < row_len && j - 1 >= 0 && is(input[i + 1][j - 1])) return `${i + 1}, ${j - 1}`
  else if (i - 1 >= 0 && j + 1 < col_len && is(input[i - 1][j + 1])) return `${i - 1}, ${j + 1}`
  else if (i - 1 >= 0 && j - 1 >= 0 && is(input[i - 1][j - 1])) return `${i - 1}, ${j - 1}`
}
const isNumber = (str: string) => ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'].includes(str)
const isSymbol = (str: string) => /[^.\w]/g.test(str)
const isAsterisk = (str: string) => str == '*'
