import { Day, NumberArrayInputParts } from '@advent/utils'

export function day1_2021(inputs: NumberArrayInputParts, slidingMeasurement: number): Day {
  const p1Input: number[] = inputs.part1
  const p2Input: number[] = inputs.part2
  return {
    part1: (): number => p1Input.filter((val, ind) => val > p1Input[ind - 1]).length,
    part2: (): number => {
      const sum: number[] = []
      let latestSum: number = -1,
        acc: number = 0
      for (let num of p2Input) {
        if (sum.length == slidingMeasurement) {
          const actualSum: number = sum.reduce((a, b) => a + b, 0)
          if (actualSum > latestSum) {
            acc++
          }
          latestSum = actualSum
          sum.shift()
        }
        sum.push(num)
      }
      return acc
    },
  }
}
