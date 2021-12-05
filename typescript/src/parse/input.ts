import { readFileSync } from 'fs'

export interface StringArrayInputParts {
  part1: string[]
  part2: string[]
}

export interface NumberArrayInputParts {
  part1: number[]
  part2: number[]
}

const INPUT_FOLDER = '../input/2021'

export const readPartAsStringArray = (num: number, part: number, sep?: string): string[] =>
  readFileSync(`${INPUT_FOLDER}/${num}.${part}`, 'utf-8').split(sep ?? '\n')

export const readPartAsNumberArray = (num: number, part: number, sep?: string): number[] =>
  readFileSync(`${INPUT_FOLDER}/${num}.${part}`, 'utf-8')
    .split(sep ?? '\n')
    .map(Number)

export const readDayInputAsNumberArray = (num: number, sep?: string): NumberArrayInputParts => ({
  part1: readPartAsNumberArray(num, 1, sep),
  part2: readPartAsNumberArray(num, 2, sep),
})

export const readDayInputAsStringArray = (num: number, sep?: string): StringArrayInputParts => ({
  part1: readPartAsStringArray(num, 1, sep),
  part2: readPartAsStringArray(num, 2, sep),
})
