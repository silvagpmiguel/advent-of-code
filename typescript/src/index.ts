import { readDayInputAsNumberArray, readDayInputAsStringArray } from './parse/input'
import { exit } from 'process'
import { Day } from './day/Day'
import Day1 from './day/Day1'
import Day2 from './day/Day2';

type KnownDays = 'day1' | 'day2';
interface YearDays {
  year2021: Record<KnownDays, Day>
}

const PUZZLES: Readonly<YearDays | any> = {
  year2021: {
    day1: Day1(readDayInputAsNumberArray(1), 3),
    day2: Day2(readDayInputAsStringArray(2))
  },
}

if (process.argv.length != 4) {
  console.error('Error, insert advent of code year and day!')
  console.error('npm start <year> <day>')
  exit()
}

const YEAR: Readonly<string> = `year${process.argv[2]}`
const DAY: Readonly<string> = `day${process.argv[3]}`

if (!(YEAR in PUZZLES && DAY in PUZZLES[YEAR])) {
  console.error(`Error, unfortunately there is no implementation for ${YEAR} - ${DAY} :(`)
  exit()
}

console.log(`Advent of Code ${YEAR} - ${DAY}`)
console.log(`part1: ${PUZZLES[YEAR][DAY].part1()}`)
console.log(`part2: ${PUZZLES[YEAR][DAY].part2()}`)
