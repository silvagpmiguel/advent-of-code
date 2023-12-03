import { readFileSync } from 'fs'

const INPUT_FOLDER = '../input'

export interface StringArrayInputParts {
  part1: string[]
  part2: string[]
}

export interface NumberArrayInputParts {
  part1: number[]
  part2: number[]
}

export const readDayInputAsNumberArray = (num: number, year: number, sep = '\n'): NumberArrayInputParts => ({
  part1: readPartAsNumberArray(num, 1, `${INPUT_FOLDER}/${year}`, sep),
  part2: readPartAsNumberArray(num, 2, `${INPUT_FOLDER}/${year}`, sep),
})

export const readDayInputAsStringArray = (num: number, year: number, sep = '\n'): StringArrayInputParts => ({
  part1: readPartAsStringArray(num, 1, `${INPUT_FOLDER}/${year}`, sep),
  part2: readPartAsStringArray(num, 2, `${INPUT_FOLDER}/${year}`, sep),
})

const readPartAsStringArray = (num: number, part: number, folder: string, sep: string): string[] =>
  readFileSync(`${folder}/${num}.${part}`, 'utf-8').trim().split(sep)

const readPartAsNumberArray = (num: number, part: number, folder: string, sep: string): number[] =>
  readPartAsStringArray(num, part, folder, sep).map(Number)
