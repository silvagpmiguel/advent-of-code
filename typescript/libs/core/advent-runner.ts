import { Day, readDayInputAsNumberArray, readDayInputAsStringArray } from '@advent/utils'
import { day1_2021, day2_2021, day3_2021, day4_2021, day5_2021, day6_2021, day7_2021 } from '@advent/2021'
import { day1_2023, day2_2023, day3_2023, day4_2023, day5_2023, day6_2023, day7_2023, day8_2023 } from '@advent/2023'

export class AdventRunner {
  private readonly puzzles: Record<string, Record<string, Day>> = {
    '2021': {
      '1': day1_2021(readDayInputAsNumberArray(1, 2021), 3),
      '2': day2_2021(readDayInputAsStringArray(2, 2021)),
      '3': day3_2021(readDayInputAsStringArray(3, 2021)),
      '4': day4_2021(readDayInputAsStringArray(4, 2021), 5),
      '5': day5_2021(readDayInputAsStringArray(5, 2021)),
      '6': day6_2021(readDayInputAsNumberArray(6, 2021, ','), 6, 8, { part1: 80, part2: 256 }),
      '7': day7_2021(readDayInputAsNumberArray(7, 2021, ',')),
    },
    '2023': {
      '1': day1_2023(readDayInputAsStringArray(1, 2023)),
      '2': day2_2023(readDayInputAsStringArray(2, 2023), { red: 12, green: 13, blue: 14 }),
      '3': day3_2023(readDayInputAsStringArray(3, 2023)),
      '4': day4_2023(readDayInputAsStringArray(4, 2023)),
      '5': day5_2023(readDayInputAsStringArray(5, 2023)),
      '6': day6_2023(readDayInputAsStringArray(6, 2023)),
      '7': day7_2023(readDayInputAsStringArray(7, 2023)),
      '8': day8_2023(readDayInputAsStringArray(8, 2023), { start: 'AAA', end: 'ZZZ' }, { start: 'A', end: 'Z' }),
    },
  }

  constructor(private year: string, private day: string) {}

  run() {
    if (!(this.year in this.puzzles && this.day && this.puzzles[this.year])) {
      console.error(`Error, no implementation for advent of code year: ${this.year}, day: ${this.day}!`)
      return
    }

    const puzzle = this.puzzles[this.year][this.day]
    console.log(`Advent of Code ${this.year} - ${this.day}`)
    console.log(`part1: ${puzzle.part1()}`)
    console.log(`part2: ${puzzle.part2()}`)
  }
}
